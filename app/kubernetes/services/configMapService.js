import angular from 'angular';
import _ from 'lodash-es';
import PortainerError from 'Portainer/error';
import KubernetesConfigMapConverter from 'Kubernetes/converters/configMap';
import { KubernetesCommonParams } from 'Kubernetes/models/common/params';

class KubernetesConfigMapService {
  /* @ngInject */
  constructor($async, KubernetesConfigMaps) {
    this.$async = $async;
    this.KubernetesConfigMaps = KubernetesConfigMaps;

    this.getAsync = this.getAsync.bind(this);
    this.getAllAsync = this.getAllAsync.bind(this);
    this.createAsync = this.createAsync.bind(this);
    this.updateAsync = this.updateAsync.bind(this);
    this.deleteAsync = this.deleteAsync.bind(this);
  }

  /**
   * GET
   */
  async getAsync(namespace, name) {
    try {
      const params = new KubernetesCommonParams();
      params.id = name;
      const [rawPromise, yamlPromise] = await Promise.allSettled([
        this.KubernetesConfigMaps(namespace).get(params).$promise,
        this.KubernetesConfigMaps(namespace).getYaml(params).$promise,
      ]);
      const configMap = KubernetesConfigMapConverter.apiToConfigMap(rawPromise.value, yamlPromise.value);
      return configMap;
    } catch (err) {
      if (err.status === 404) {
        return KubernetesConfigMapConverter.defaultConfigMap(namespace, name);
      }
      throw new PortainerError('Unable to retrieve config map', err);
    }
  }

  async getAllAsync(namespace) {
    try {
      const data = await this.KubernetesConfigMaps(namespace).get().$promise;
      return _.map(data.items, (item) => KubernetesConfigMapConverter.apiToConfigMap(item));
    } catch (err) {
      throw new PortainerError('Unable to retrieve config maps', err);
    }
  }

  get(namespace, name) {
    if (name) {
      return this.$async(this.getAsync, namespace, name);
    }
    return this.$async(this.getAllAsync, namespace);
  }

  /**
   * CREATE
   */
  async createAsync(config) {
    try {
      const payload = KubernetesConfigMapConverter.createPayload(config);
      const params = {};
      const namespace = payload.metadata.namespace;
      const data = await this.KubernetesConfigMaps(namespace).create(params, payload).$promise;
      return KubernetesConfigMapConverter.apiToConfigMap(data);
    } catch (err) {
      throw new PortainerError('Unable to create config map', err);
    }
  }

  create(config) {
    return this.$async(this.createAsync, config);
  }

  /**
   * UPDATE
   */
  async updateAsync(config) {
    try {
      if (!config.Id) {
        return await this.create(config);
      }
      const payload = KubernetesConfigMapConverter.updatePayload(config);
      const params = new KubernetesCommonParams();
      params.id = payload.metadata.name;
      const namespace = payload.metadata.namespace;
      const data = await this.KubernetesConfigMaps(namespace).update(params, payload).$promise;
      return KubernetesConfigMapConverter.apiToConfigMap(data);
    } catch (err) {
      throw new PortainerError('Unable to update config map', err);
    }
  }
  update(config) {
    return this.$async(this.updateAsync, config);
  }

  /**
   * DELETE
   */
  async deleteAsync(config) {
    try {
      const params = new KubernetesCommonParams();
      params.id = config.Name;
      const namespace = config.Namespace;
      await this.KubernetesConfigMaps(namespace).delete(params).$promise;
    } catch (err) {
      throw new PortainerError('Unable to delete config map', err);
    }
  }

  delete(config) {
    return this.$async(this.deleteAsync, config);
  }
}

export default KubernetesConfigMapService;
angular.module('portainer.kubernetes').service('KubernetesConfigMapService', KubernetesConfigMapService);
