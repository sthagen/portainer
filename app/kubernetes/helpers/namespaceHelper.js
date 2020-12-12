import _ from 'lodash-es';
import angular from 'angular';

class KubernetesNamespaceHelper {
  /* @ngInject */
  constructor(KUBERNETES_SYSTEM_NAMESPACES) {
    this.KUBERNETES_SYSTEM_NAMESPACES = KUBERNETES_SYSTEM_NAMESPACES;
  }

  isSystemNamespace(namespace) {
    return _.includes(this.KUBERNETES_SYSTEM_NAMESPACES, namespace);
  }
}

export default KubernetesNamespaceHelper;
angular.module('portainer.app').service('KubernetesNamespaceHelper', KubernetesNamespaceHelper);
