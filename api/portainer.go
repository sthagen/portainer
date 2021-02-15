package portainer

import (
	"io"
	"time"
)

type (
	// AccessPolicy represent a policy that can be associated to a user or team
	AccessPolicy struct {
		RoleID RoleID `json:"RoleId"`
	}

	// AgentPlatform represents a platform type for an Agent
	AgentPlatform int

	// AuthenticationMethod represents the authentication method used to authenticate a user
	AuthenticationMethod int

	// Authorization represents an authorization associated to an operation
	Authorization string

	// Authorizations represents a set of authorizations associated to a role
	Authorizations map[Authorization]bool

	// AzureCredentials represents the credentials used to connect to an Azure
	// environment.
	AzureCredentials struct {
		ApplicationID     string `json:"ApplicationID"`
		TenantID          string `json:"TenantID"`
		AuthenticationKey string `json:"AuthenticationKey"`
	}

	// CLIFlags represents the available flags on the CLI
	CLIFlags struct {
		Addr                      *string
		TunnelAddr                *string
		TunnelPort                *string
		AdminPassword             *string
		AdminPasswordFile         *string
		Assets                    *string
		Data                      *string
		EnableEdgeComputeFeatures *bool
		EndpointURL               *string
		Labels                    *[]Pair
		Logo                      *string
		NoAnalytics               *bool
		Templates                 *string
		TLS                       *bool
		TLSSkipVerify             *bool
		TLSCacert                 *string
		TLSCert                   *string
		TLSKey                    *string
		SSL                       *bool
		SSLCert                   *string
		SSLKey                    *string
		SnapshotInterval          *string
	}

	// CustomTemplate represents a custom template
	CustomTemplate struct {
		ID              CustomTemplateID       `json:"Id"`
		Title           string                 `json:"Title"`
		Description     string                 `json:"Description"`
		ProjectPath     string                 `json:"ProjectPath"`
		EntryPoint      string                 `json:"EntryPoint"`
		CreatedByUserID UserID                 `json:"CreatedByUserId"`
		Note            string                 `json:"Note"`
		Platform        CustomTemplatePlatform `json:"Platform"`
		Logo            string                 `json:"Logo"`
		Type            StackType              `json:"Type"`
		ResourceControl *ResourceControl       `json:"ResourceControl"`
	}

	// CustomTemplateID represents a custom template identifier
	CustomTemplateID int

	// CustomTemplatePlatform represents a custom template platform
	CustomTemplatePlatform int

	// DockerHub represents all the required information to connect and use the
	// Docker Hub
	DockerHub struct {
		Authentication bool   `json:"Authentication"`
		Username       string `json:"Username"`
		Password       string `json:"Password,omitempty"`
	}

	// DockerSnapshot represents a snapshot of a specific Docker endpoint at a specific time
	DockerSnapshot struct {
		Time                    int64             `json:"Time"`
		DockerVersion           string            `json:"DockerVersion"`
		Swarm                   bool              `json:"Swarm"`
		TotalCPU                int               `json:"TotalCPU"`
		TotalMemory             int64             `json:"TotalMemory"`
		RunningContainerCount   int               `json:"RunningContainerCount"`
		StoppedContainerCount   int               `json:"StoppedContainerCount"`
		HealthyContainerCount   int               `json:"HealthyContainerCount"`
		UnhealthyContainerCount int               `json:"UnhealthyContainerCount"`
		VolumeCount             int               `json:"VolumeCount"`
		ImageCount              int               `json:"ImageCount"`
		ServiceCount            int               `json:"ServiceCount"`
		StackCount              int               `json:"StackCount"`
		SnapshotRaw             DockerSnapshotRaw `json:"DockerSnapshotRaw"`
	}

	// DockerSnapshotRaw represents all the information related to a snapshot as returned by the Docker API
	DockerSnapshotRaw struct {
		Containers interface{} `json:"Containers"`
		Volumes    interface{} `json:"Volumes"`
		Networks   interface{} `json:"Networks"`
		Images     interface{} `json:"Images"`
		Info       interface{} `json:"Info"`
		Version    interface{} `json:"Version"`
	}

	// EdgeGroup represents an Edge group
	EdgeGroup struct {
		ID           EdgeGroupID  `json:"Id"`
		Name         string       `json:"Name"`
		Dynamic      bool         `json:"Dynamic"`
		TagIDs       []TagID      `json:"TagIds"`
		Endpoints    []EndpointID `json:"Endpoints"`
		PartialMatch bool         `json:"PartialMatch"`
	}

	// EdgeGroupID represents an Edge group identifier
	EdgeGroupID int

	// EdgeJob represents a job that can run on Edge environments.
	EdgeJob struct {
		ID             EdgeJobID                          `json:"Id"`
		Created        int64                              `json:"Created"`
		CronExpression string                             `json:"CronExpression"`
		Endpoints      map[EndpointID]EdgeJobEndpointMeta `json:"Endpoints"`
		Name           string                             `json:"Name"`
		ScriptPath     string                             `json:"ScriptPath"`
		Recurring      bool                               `json:"Recurring"`
		Version        int                                `json:"Version"`
	}

	// EdgeJobEndpointMeta represents a meta data object for an Edge job and Endpoint relation
	EdgeJobEndpointMeta struct {
		LogsStatus  EdgeJobLogsStatus
		CollectLogs bool
	}

	// EdgeJobID represents an Edge job identifier
	EdgeJobID int

	// EdgeJobLogsStatus represent status of logs collection job
	EdgeJobLogsStatus int

	// EdgeSchedule represents a scheduled job that can run on Edge environments.
	// Deprecated in favor of EdgeJob
	EdgeSchedule struct {
		ID             ScheduleID   `json:"Id"`
		CronExpression string       `json:"CronExpression"`
		Script         string       `json:"Script"`
		Version        int          `json:"Version"`
		Endpoints      []EndpointID `json:"Endpoints"`
	}

	//EdgeStack represents an edge stack
	EdgeStack struct {
		ID           EdgeStackID                    `json:"Id"`
		Name         string                         `json:"Name"`
		Status       map[EndpointID]EdgeStackStatus `json:"Status"`
		CreationDate int64                          `json:"CreationDate"`
		EdgeGroups   []EdgeGroupID                  `json:"EdgeGroups"`
		ProjectPath  string                         `json:"ProjectPath"`
		EntryPoint   string                         `json:"EntryPoint"`
		Version      int                            `json:"Version"`
		Prune        bool                           `json:"Prune"`
	}

	//EdgeStackID represents an edge stack id
	EdgeStackID int

	//EdgeStackStatus represents an edge stack status
	EdgeStackStatus struct {
		Type       EdgeStackStatusType `json:"Type"`
		Error      string              `json:"Error"`
		EndpointID EndpointID          `json:"EndpointID"`
	}

	//EdgeStackStatusType represents an edge stack status type
	EdgeStackStatusType int

	// Endpoint represents a Docker endpoint with all the info required
	// to connect to it
	Endpoint struct {
		ID                      EndpointID          `json:"Id"`
		Name                    string              `json:"Name"`
		Type                    EndpointType        `json:"Type"`
		URL                     string              `json:"URL"`
		GroupID                 EndpointGroupID     `json:"GroupId"`
		PublicURL               string              `json:"PublicURL"`
		TLSConfig               TLSConfiguration    `json:"TLSConfig"`
		Extensions              []EndpointExtension `json:"Extensions"`
		AzureCredentials        AzureCredentials    `json:"AzureCredentials,omitempty"`
		TagIDs                  []TagID             `json:"TagIds"`
		Status                  EndpointStatus      `json:"Status"`
		Snapshots               []DockerSnapshot    `json:"Snapshots"`
		UserAccessPolicies      UserAccessPolicies  `json:"UserAccessPolicies"`
		TeamAccessPolicies      TeamAccessPolicies  `json:"TeamAccessPolicies"`
		EdgeID                  string              `json:"EdgeID,omitempty"`
		EdgeKey                 string              `json:"EdgeKey"`
		EdgeCheckinInterval     int                 `json:"EdgeCheckinInterval"`
		Kubernetes              KubernetesData      `json:"Kubernetes"`
		ComposeSyntaxMaxVersion string              `json:"ComposeSyntaxMaxVersion"`
		SecuritySettings        EndpointSecuritySettings

		// Deprecated fields
		// Deprecated in DBVersion == 4
		TLS           bool   `json:"TLS,omitempty"`
		TLSCACertPath string `json:"TLSCACert,omitempty"`
		TLSCertPath   string `json:"TLSCert,omitempty"`
		TLSKeyPath    string `json:"TLSKey,omitempty"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []TeamID `json:"AuthorizedTeams"`

		// Deprecated in DBVersion == 22
		Tags []string `json:"Tags"`
	}

	// EndpointAuthorizations represents the authorizations associated to a set of endpoints
	EndpointAuthorizations map[EndpointID]Authorizations

	// EndpointExtension represents a deprecated form of Portainer extension
	// TODO: legacy extension management
	EndpointExtension struct {
		Type EndpointExtensionType `json:"Type"`
		URL  string                `json:"URL"`
	}

	// EndpointExtensionType represents the type of an endpoint extension. Only
	// one extension of each type can be associated to an endpoint
	EndpointExtensionType int

	// EndpointGroup represents a group of endpoints
	EndpointGroup struct {
		ID                 EndpointGroupID    `json:"Id"`
		Name               string             `json:"Name"`
		Description        string             `json:"Description"`
		UserAccessPolicies UserAccessPolicies `json:"UserAccessPolicies"`
		TeamAccessPolicies TeamAccessPolicies `json:"TeamAccessPolicies"`
		TagIDs             []TagID            `json:"TagIds"`

		// Deprecated fields
		Labels []Pair `json:"Labels"`

		// Deprecated in DBVersion == 18
		AuthorizedUsers []UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []TeamID `json:"AuthorizedTeams"`

		// Deprecated in DBVersion == 22
		Tags []string `json:"Tags"`
	}

	// EndpointGroupID represents an endpoint group identifier
	EndpointGroupID int

	// EndpointID represents an endpoint identifier
	EndpointID int

	// EndpointStatus represents the status of an endpoint
	EndpointStatus int

	// EndpointSyncJob represents a scheduled job that synchronize endpoints based on an external file
	// Deprecated
	EndpointSyncJob struct{}

	// EndpointSecuritySettings represents settings for an endpoint
	EndpointSecuritySettings struct {
		AllowBindMountsForRegularUsers            bool `json:"allowBindMountsForRegularUsers"`
		AllowPrivilegedModeForRegularUsers        bool `json:"allowPrivilegedModeForRegularUsers"`
		AllowVolumeBrowserForRegularUsers         bool `json:"allowVolumeBrowserForRegularUsers"`
		AllowHostNamespaceForRegularUsers         bool `json:"allowHostNamespaceForRegularUsers"`
		AllowDeviceMappingForRegularUsers         bool `json:"allowDeviceMappingForRegularUsers"`
		AllowStackManagementForRegularUsers       bool `json:"allowStackManagementForRegularUsers"`
		AllowContainerCapabilitiesForRegularUsers bool `json:"allowContainerCapabilitiesForRegularUsers"`
		EnableHostManagementFeatures              bool `json:"enableHostManagementFeatures"`
	}

	// EndpointType represents the type of an endpoint
	EndpointType int

	// EndpointRelation represnts a endpoint relation object
	EndpointRelation struct {
		EndpointID EndpointID
		EdgeStacks map[EdgeStackID]bool
	}

	// Extension represents a deprecated Portainer extension
	Extension struct {
		ID               ExtensionID        `json:"Id"`
		Enabled          bool               `json:"Enabled"`
		Name             string             `json:"Name,omitempty"`
		ShortDescription string             `json:"ShortDescription,omitempty"`
		Description      string             `json:"Description,omitempty"`
		DescriptionURL   string             `json:"DescriptionURL,omitempty"`
		Price            string             `json:"Price,omitempty"`
		PriceDescription string             `json:"PriceDescription,omitempty"`
		Deal             bool               `json:"Deal,omitempty"`
		Available        bool               `json:"Available,omitempty"`
		License          LicenseInformation `json:"License,omitempty"`
		Version          string             `json:"Version"`
		UpdateAvailable  bool               `json:"UpdateAvailable"`
		ShopURL          string             `json:"ShopURL,omitempty"`
		Images           []string           `json:"Images,omitempty"`
		Logo             string             `json:"Logo,omitempty"`
	}

	// ExtensionID represents a extension identifier
	ExtensionID int

	// GitlabRegistryData represents data required for gitlab registry to work
	GitlabRegistryData struct {
		ProjectID   int    `json:"ProjectId"`
		InstanceURL string `json:"InstanceURL"`
		ProjectPath string `json:"ProjectPath"`
	}

	// JobType represents a job type
	JobType int

	// KubernetesData contains all the Kubernetes related endpoint information
	KubernetesData struct {
		Snapshots     []KubernetesSnapshot    `json:"Snapshots"`
		Configuration KubernetesConfiguration `json:"Configuration"`
	}

	// KubernetesSnapshot represents a snapshot of a specific Kubernetes endpoint at a specific time
	KubernetesSnapshot struct {
		Time              int64  `json:"Time"`
		KubernetesVersion string `json:"KubernetesVersion"`
		NodeCount         int    `json:"NodeCount"`
		TotalCPU          int64  `json:"TotalCPU"`
		TotalMemory       int64  `json:"TotalMemory"`
	}

	// KubernetesConfiguration represents the configuration of a Kubernetes endpoint
	KubernetesConfiguration struct {
		UseLoadBalancer  bool                           `json:"UseLoadBalancer"`
		UseServerMetrics bool                           `json:"UseServerMetrics"`
		StorageClasses   []KubernetesStorageClassConfig `json:"StorageClasses"`
		IngressClasses   []KubernetesIngressClassConfig `json:"IngressClasses"`
	}

	// KubernetesStorageClassConfig represents a Kubernetes Storage Class configuration
	KubernetesStorageClassConfig struct {
		Name                 string   `json:"Name"`
		AccessModes          []string `json:"AccessModes"`
		Provisioner          string   `json:"Provisioner"`
		AllowVolumeExpansion bool     `json:"AllowVolumeExpansion"`
	}

	// KubernetesIngressClassConfig represents a Kubernetes Ingress Class configuration
	KubernetesIngressClassConfig struct {
		Name string `json:"Name"`
		Type string `json:"Type"`
	}

	// LDAPGroupSearchSettings represents settings used to search for groups in a LDAP server
	LDAPGroupSearchSettings struct {
		GroupBaseDN    string `json:"GroupBaseDN"`
		GroupFilter    string `json:"GroupFilter"`
		GroupAttribute string `json:"GroupAttribute"`
	}

	// LDAPSearchSettings represents settings used to search for users in a LDAP server
	LDAPSearchSettings struct {
		BaseDN            string `json:"BaseDN"`
		Filter            string `json:"Filter"`
		UserNameAttribute string `json:"UserNameAttribute"`
	}

	// LDAPSettings represents the settings used to connect to a LDAP server
	LDAPSettings struct {
		AnonymousMode       bool                      `json:"AnonymousMode"`
		ReaderDN            string                    `json:"ReaderDN"`
		Password            string                    `json:"Password,omitempty"`
		URL                 string                    `json:"URL"`
		TLSConfig           TLSConfiguration          `json:"TLSConfig"`
		StartTLS            bool                      `json:"StartTLS"`
		SearchSettings      []LDAPSearchSettings      `json:"SearchSettings"`
		GroupSearchSettings []LDAPGroupSearchSettings `json:"GroupSearchSettings"`
		AutoCreateUsers     bool                      `json:"AutoCreateUsers"`
	}

	// LicenseInformation represents information about an extension license
	LicenseInformation struct {
		LicenseKey string `json:"LicenseKey,omitempty"`
		Company    string `json:"Company,omitempty"`
		Expiration string `json:"Expiration,omitempty"`
		Valid      bool   `json:"Valid,omitempty"`
	}

	// MembershipRole represents the role of a user within a team
	MembershipRole int

	// OAuthSettings represents the settings used to authorize with an authorization server
	OAuthSettings struct {
		ClientID             string `json:"ClientID"`
		ClientSecret         string `json:"ClientSecret,omitempty"`
		AccessTokenURI       string `json:"AccessTokenURI"`
		AuthorizationURI     string `json:"AuthorizationURI"`
		ResourceURI          string `json:"ResourceURI"`
		RedirectURI          string `json:"RedirectURI"`
		UserIdentifier       string `json:"UserIdentifier"`
		Scopes               string `json:"Scopes"`
		OAuthAutoCreateUsers bool   `json:"OAuthAutoCreateUsers"`
		DefaultTeamID        TeamID `json:"DefaultTeamID"`
	}

	// Pair defines a key/value string pair
	Pair struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	// Registry represents a Docker registry with all the info required
	// to connect to it
	Registry struct {
		ID                      RegistryID                       `json:"Id"`
		Type                    RegistryType                     `json:"Type"`
		Name                    string                           `json:"Name"`
		URL                     string                           `json:"URL"`
		Authentication          bool                             `json:"Authentication"`
		Username                string                           `json:"Username"`
		Password                string                           `json:"Password,omitempty"`
		ManagementConfiguration *RegistryManagementConfiguration `json:"ManagementConfiguration"`
		Gitlab                  GitlabRegistryData               `json:"Gitlab"`
		UserAccessPolicies      UserAccessPolicies               `json:"UserAccessPolicies"`
		TeamAccessPolicies      TeamAccessPolicies               `json:"TeamAccessPolicies"`

		// Deprecated fields
		// Deprecated in DBVersion == 18
		AuthorizedUsers []UserID `json:"AuthorizedUsers"`
		AuthorizedTeams []TeamID `json:"AuthorizedTeams"`
	}

	// RegistryID represents a registry identifier
	RegistryID int

	// RegistryManagementConfiguration represents a configuration that can be used to query
	// the registry API via the registry management extension.
	RegistryManagementConfiguration struct {
		Type           RegistryType     `json:"Type"`
		Authentication bool             `json:"Authentication"`
		Username       string           `json:"Username"`
		Password       string           `json:"Password"`
		TLSConfig      TLSConfiguration `json:"TLSConfig"`
	}

	// RegistryType represents a type of registry
	RegistryType int

	// ResourceAccessLevel represents the level of control associated to a resource
	ResourceAccessLevel int

	// ResourceControl represent a reference to a Docker resource with specific access controls
	ResourceControl struct {
		ID                 ResourceControlID    `json:"Id"`
		ResourceID         string               `json:"ResourceId"`
		SubResourceIDs     []string             `json:"SubResourceIds"`
		Type               ResourceControlType  `json:"Type"`
		UserAccesses       []UserResourceAccess `json:"UserAccesses"`
		TeamAccesses       []TeamResourceAccess `json:"TeamAccesses"`
		Public             bool                 `json:"Public"`
		AdministratorsOnly bool                 `json:"AdministratorsOnly"`
		System             bool                 `json:"System"`

		// Deprecated fields
		// Deprecated in DBVersion == 2
		OwnerID     UserID              `json:"OwnerId,omitempty"`
		AccessLevel ResourceAccessLevel `json:"AccessLevel,omitempty"`
	}

	// ResourceControlID represents a resource control identifier
	ResourceControlID int

	// ResourceControlType represents the type of resource associated to the resource control (volume, container, service...)
	ResourceControlType int

	// Role represents a set of authorizations that can be associated to a user or
	// to a team.
	Role struct {
		ID             RoleID         `json:"Id"`
		Name           string         `json:"Name"`
		Description    string         `json:"Description"`
		Authorizations Authorizations `json:"Authorizations"`
		Priority       int            `json:"Priority"`
	}

	// RoleID represents a role identifier
	RoleID int

	// Schedule represents a scheduled job.
	// It only contains a pointer to one of the JobRunner implementations
	// based on the JobType.
	// NOTE: The Recurring option is only used by ScriptExecutionJob at the moment
	// Deprecated in favor of EdgeJob
	Schedule struct {
		ID             ScheduleID `json:"Id"`
		Name           string
		CronExpression string
		Recurring      bool
		Created        int64
		JobType        JobType
		EdgeSchedule   *EdgeSchedule
	}

	// ScheduleID represents a schedule identifier.
	// Deprecated in favor of EdgeJob
	ScheduleID int

	// ScriptExecutionJob represents a scheduled job that can execute a script via a privileged container
	ScriptExecutionJob struct {
		Endpoints     []EndpointID
		Image         string
		ScriptPath    string
		RetryCount    int
		RetryInterval int
	}

	// Settings represents the application settings
	Settings struct {
		LogoURL                   string               `json:"LogoURL"`
		BlackListedLabels         []Pair               `json:"BlackListedLabels"`
		AuthenticationMethod      AuthenticationMethod `json:"AuthenticationMethod"`
		LDAPSettings              LDAPSettings         `json:"LDAPSettings"`
		OAuthSettings             OAuthSettings        `json:"OAuthSettings"`
		SnapshotInterval          string               `json:"SnapshotInterval"`
		TemplatesURL              string               `json:"TemplatesURL"`
		EdgeAgentCheckinInterval  int                  `json:"EdgeAgentCheckinInterval"`
		EnableEdgeComputeFeatures bool                 `json:"EnableEdgeComputeFeatures"`
		UserSessionTimeout        string               `json:"UserSessionTimeout"`
		EnableTelemetry           bool                 `json:"EnableTelemetry"`

		// Deprecated fields
		DisplayDonationHeader       bool
		DisplayExternalContributors bool

		// Deprecated fields v26
		EnableHostManagementFeatures              bool `json:"EnableHostManagementFeatures"`
		AllowVolumeBrowserForRegularUsers         bool `json:"AllowVolumeBrowserForRegularUsers"`
		AllowBindMountsForRegularUsers            bool `json:"AllowBindMountsForRegularUsers"`
		AllowPrivilegedModeForRegularUsers        bool `json:"AllowPrivilegedModeForRegularUsers"`
		AllowHostNamespaceForRegularUsers         bool `json:"AllowHostNamespaceForRegularUsers"`
		AllowStackManagementForRegularUsers       bool `json:"AllowStackManagementForRegularUsers"`
		AllowDeviceMappingForRegularUsers         bool `json:"AllowDeviceMappingForRegularUsers"`
		AllowContainerCapabilitiesForRegularUsers bool `json:"AllowContainerCapabilitiesForRegularUsers"`
	}

	// SnapshotJob represents a scheduled job that can create endpoint snapshots
	SnapshotJob struct{}

	// Stack represents a Docker stack created via docker stack deploy
	Stack struct {
		ID              StackID          `json:"Id"`
		Name            string           `json:"Name"`
		Type            StackType        `json:"Type"`
		EndpointID      EndpointID       `json:"EndpointId"`
		SwarmID         string           `json:"SwarmId"`
		EntryPoint      string           `json:"EntryPoint"`
		Env             []Pair           `json:"Env"`
		ResourceControl *ResourceControl `json:"ResourceControl"`
		Status          StackStatus      `json:"Status"`
		CreationDate    int64
		CreatedBy       string
		UpdateDate      int64
		UpdatedBy       string
		ProjectPath     string
	}

	// StackID represents a stack identifier (it must be composed of Name + "_" + SwarmID to create a unique identifier)
	StackID int

	// StackStatus represent a status for a stack
	StackStatus int

	// StackType represents the type of the stack (compose v2, stack deploy v3)
	StackType int

	// Status represents the application status
	Status struct {
		Version string `json:"Version"`
	}

	// Tag represents a tag that can be associated to a resource
	Tag struct {
		ID             TagID
		Name           string                   `json:"Name"`
		Endpoints      map[EndpointID]bool      `json:"Endpoints"`
		EndpointGroups map[EndpointGroupID]bool `json:"EndpointGroups"`
	}

	// TagID represents a tag identifier
	TagID int

	// Team represents a list of user accounts
	Team struct {
		ID   TeamID `json:"Id"`
		Name string `json:"Name"`
	}

	// TeamAccessPolicies represent the association of an access policy and a team
	TeamAccessPolicies map[TeamID]AccessPolicy

	// TeamID represents a team identifier
	TeamID int

	// TeamMembership represents a membership association between a user and a team
	TeamMembership struct {
		ID     TeamMembershipID `json:"Id"`
		UserID UserID           `json:"UserID"`
		TeamID TeamID           `json:"TeamID"`
		Role   MembershipRole   `json:"Role"`
	}

	// TeamMembershipID represents a team membership identifier
	TeamMembershipID int

	// TeamResourceAccess represents the level of control on a resource for a specific team
	TeamResourceAccess struct {
		TeamID      TeamID              `json:"TeamId"`
		AccessLevel ResourceAccessLevel `json:"AccessLevel"`
	}

	// Template represents an application template that can be used as an App Template
	// or an Edge template
	Template struct {
		// Mandatory container/stack fields
		ID                TemplateID   `json:"Id"`
		Type              TemplateType `json:"type"`
		Title             string       `json:"title"`
		Description       string       `json:"description"`
		AdministratorOnly bool         `json:"administrator_only"`

		// Mandatory container fields
		Image string `json:"image"`

		// Mandatory stack fields
		Repository TemplateRepository `json:"repository"`

		// Mandatory Edge stack fields
		StackFile string `json:"stackFile"`

		// Optional stack/container fields
		Name       string        `json:"name,omitempty"`
		Logo       string        `json:"logo,omitempty"`
		Env        []TemplateEnv `json:"env,omitempty"`
		Note       string        `json:"note,omitempty"`
		Platform   string        `json:"platform,omitempty"`
		Categories []string      `json:"categories,omitempty"`

		// Optional container fields
		Registry      string           `json:"registry,omitempty"`
		Command       string           `json:"command,omitempty"`
		Network       string           `json:"network,omitempty"`
		Volumes       []TemplateVolume `json:"volumes,omitempty"`
		Ports         []string         `json:"ports,omitempty"`
		Labels        []Pair           `json:"labels,omitempty"`
		Privileged    bool             `json:"privileged,omitempty"`
		Interactive   bool             `json:"interactive,omitempty"`
		RestartPolicy string           `json:"restart_policy,omitempty"`
		Hostname      string           `json:"hostname,omitempty"`
	}

	// TemplateEnv represents a template environment variable configuration
	TemplateEnv struct {
		Name        string              `json:"name"`
		Label       string              `json:"label,omitempty"`
		Description string              `json:"description,omitempty"`
		Default     string              `json:"default,omitempty"`
		Preset      bool                `json:"preset,omitempty"`
		Select      []TemplateEnvSelect `json:"select,omitempty"`
	}

	// TemplateEnvSelect represents text/value pair that will be displayed as a choice for the
	// template user
	TemplateEnvSelect struct {
		Text    string `json:"text"`
		Value   string `json:"value"`
		Default bool   `json:"default"`
	}

	// TemplateID represents a template identifier
	TemplateID int

	// TemplateRepository represents the git repository configuration for a template
	TemplateRepository struct {
		URL       string `json:"url"`
		StackFile string `json:"stackfile"`
	}

	// TemplateType represents the type of a template
	TemplateType int

	// TemplateVolume represents a template volume configuration
	TemplateVolume struct {
		Container string `json:"container"`
		Bind      string `json:"bind,omitempty"`
		ReadOnly  bool   `json:"readonly,omitempty"`
	}

	// TLSConfiguration represents a TLS configuration
	TLSConfiguration struct {
		TLS           bool   `json:"TLS"`
		TLSSkipVerify bool   `json:"TLSSkipVerify"`
		TLSCACertPath string `json:"TLSCACert,omitempty"`
		TLSCertPath   string `json:"TLSCert,omitempty"`
		TLSKeyPath    string `json:"TLSKey,omitempty"`
	}

	// TLSFileType represents a type of TLS file required to connect to a Docker endpoint.
	// It can be either a TLS CA file, a TLS certificate file or a TLS key file
	TLSFileType int

	// TokenData represents the data embedded in a JWT token
	TokenData struct {
		ID       UserID
		Username string
		Role     UserRole
	}

	// TunnelDetails represents information associated to a tunnel
	TunnelDetails struct {
		Status       string
		LastActivity time.Time
		Port         int
		Jobs         []EdgeJob
		Credentials  string
	}

	// TunnelServerInfo represents information associated to the tunnel server
	TunnelServerInfo struct {
		PrivateKeySeed string `json:"PrivateKeySeed"`
	}

	// User represents a user account
	User struct {
		ID       UserID   `json:"Id"`
		Username string   `json:"Username"`
		Password string   `json:"Password,omitempty"`
		Role     UserRole `json:"Role"`

		// Deprecated fields
		// Deprecated in DBVersion == 25
		PortainerAuthorizations Authorizations         `json:"PortainerAuthorizations"`
		EndpointAuthorizations  EndpointAuthorizations `json:"EndpointAuthorizations"`
	}

	// UserAccessPolicies represent the association of an access policy and a user
	UserAccessPolicies map[UserID]AccessPolicy

	// UserID represents a user identifier
	UserID int

	// UserResourceAccess represents the level of control on a resource for a specific user
	UserResourceAccess struct {
		UserID      UserID              `json:"UserId"`
		AccessLevel ResourceAccessLevel `json:"AccessLevel"`
	}

	// UserRole represents the role of a user. It can be either an administrator
	// or a regular user
	UserRole int

	// Webhook represents a url webhook that can be used to update a service
	Webhook struct {
		ID          WebhookID   `json:"Id"`
		Token       string      `json:"Token"`
		ResourceID  string      `json:"ResourceId"`
		EndpointID  EndpointID  `json:"EndpointId"`
		WebhookType WebhookType `json:"Type"`
	}

	// WebhookID represents a webhook identifier.
	WebhookID int

	// WebhookType represents the type of resource a webhook is related to
	WebhookType int

	// CLIService represents a service for managing CLI
	CLIService interface {
		ParseFlags(version string) (*CLIFlags, error)
		ValidateFlags(flags *CLIFlags) error
	}

	// ComposeStackManager represents a service to manage Compose stacks
	ComposeStackManager interface {
		ComposeSyntaxMaxVersion() string
		Up(stack *Stack, endpoint *Endpoint) error
		Down(stack *Stack, endpoint *Endpoint) error
	}

	// CryptoService represents a service for encrypting/hashing data
	CryptoService interface {
		Hash(data string) (string, error)
		CompareHashAndData(hash string, data string) error
	}

	// CustomTemplateService represents a service to manage custom templates
	CustomTemplateService interface {
		GetNextIdentifier() int
		CustomTemplates() ([]CustomTemplate, error)
		CustomTemplate(ID CustomTemplateID) (*CustomTemplate, error)
		CreateCustomTemplate(customTemplate *CustomTemplate) error
		UpdateCustomTemplate(ID CustomTemplateID, customTemplate *CustomTemplate) error
		DeleteCustomTemplate(ID CustomTemplateID) error
	}

	// DataStore defines the interface to manage the data
	DataStore interface {
		Open() error
		Init() error
		Close() error
		IsNew() bool
		MigrateData() error

		DockerHub() DockerHubService
		CustomTemplate() CustomTemplateService
		EdgeGroup() EdgeGroupService
		EdgeJob() EdgeJobService
		EdgeStack() EdgeStackService
		Endpoint() EndpointService
		EndpointGroup() EndpointGroupService
		EndpointRelation() EndpointRelationService
		Registry() RegistryService
		ResourceControl() ResourceControlService
		Role() RoleService
		Settings() SettingsService
		Stack() StackService
		Tag() TagService
		TeamMembership() TeamMembershipService
		Team() TeamService
		TunnelServer() TunnelServerService
		User() UserService
		Version() VersionService
		Webhook() WebhookService
	}

	// DigitalSignatureService represents a service to manage digital signatures
	DigitalSignatureService interface {
		ParseKeyPair(private, public []byte) error
		GenerateKeyPair() ([]byte, []byte, error)
		EncodedPublicKey() string
		PEMHeaders() (string, string)
		CreateSignature(message string) (string, error)
	}

	// DockerHubService represents a service for managing the DockerHub object
	DockerHubService interface {
		DockerHub() (*DockerHub, error)
		UpdateDockerHub(registry *DockerHub) error
	}

	// DockerSnapshotter represents a service used to create Docker endpoint snapshots
	DockerSnapshotter interface {
		CreateSnapshot(endpoint *Endpoint) (*DockerSnapshot, error)
	}

	// EdgeGroupService represents a service to manage Edge groups
	EdgeGroupService interface {
		EdgeGroups() ([]EdgeGroup, error)
		EdgeGroup(ID EdgeGroupID) (*EdgeGroup, error)
		CreateEdgeGroup(group *EdgeGroup) error
		UpdateEdgeGroup(ID EdgeGroupID, group *EdgeGroup) error
		DeleteEdgeGroup(ID EdgeGroupID) error
	}

	// EdgeJobService represents a service to manage Edge jobs
	EdgeJobService interface {
		EdgeJobs() ([]EdgeJob, error)
		EdgeJob(ID EdgeJobID) (*EdgeJob, error)
		CreateEdgeJob(edgeJob *EdgeJob) error
		UpdateEdgeJob(ID EdgeJobID, edgeJob *EdgeJob) error
		DeleteEdgeJob(ID EdgeJobID) error
		GetNextIdentifier() int
	}

	// EdgeStackService represents a service to manage Edge stacks
	EdgeStackService interface {
		EdgeStacks() ([]EdgeStack, error)
		EdgeStack(ID EdgeStackID) (*EdgeStack, error)
		CreateEdgeStack(edgeStack *EdgeStack) error
		UpdateEdgeStack(ID EdgeStackID, edgeStack *EdgeStack) error
		DeleteEdgeStack(ID EdgeStackID) error
		GetNextIdentifier() int
	}

	// EndpointService represents a service for managing endpoint data
	EndpointService interface {
		Endpoint(ID EndpointID) (*Endpoint, error)
		Endpoints() ([]Endpoint, error)
		CreateEndpoint(endpoint *Endpoint) error
		UpdateEndpoint(ID EndpointID, endpoint *Endpoint) error
		DeleteEndpoint(ID EndpointID) error
		Synchronize(toCreate, toUpdate, toDelete []*Endpoint) error
		GetNextIdentifier() int
	}

	// EndpointGroupService represents a service for managing endpoint group data
	EndpointGroupService interface {
		EndpointGroup(ID EndpointGroupID) (*EndpointGroup, error)
		EndpointGroups() ([]EndpointGroup, error)
		CreateEndpointGroup(group *EndpointGroup) error
		UpdateEndpointGroup(ID EndpointGroupID, group *EndpointGroup) error
		DeleteEndpointGroup(ID EndpointGroupID) error
	}

	// EndpointRelationService represents a service for managing endpoint relations data
	EndpointRelationService interface {
		EndpointRelation(EndpointID EndpointID) (*EndpointRelation, error)
		CreateEndpointRelation(endpointRelation *EndpointRelation) error
		UpdateEndpointRelation(EndpointID EndpointID, endpointRelation *EndpointRelation) error
		DeleteEndpointRelation(EndpointID EndpointID) error
	}

	// FileService represents a service for managing files
	FileService interface {
		GetFileContent(filePath string) ([]byte, error)
		Rename(oldPath, newPath string) error
		RemoveDirectory(directoryPath string) error
		StoreTLSFileFromBytes(folder string, fileType TLSFileType, data []byte) (string, error)
		GetPathForTLSFile(folder string, fileType TLSFileType) (string, error)
		DeleteTLSFile(folder string, fileType TLSFileType) error
		DeleteTLSFiles(folder string) error
		GetStackProjectPath(stackIdentifier string) string
		StoreStackFileFromBytes(stackIdentifier, fileName string, data []byte) (string, error)
		GetEdgeStackProjectPath(edgeStackIdentifier string) string
		StoreEdgeStackFileFromBytes(edgeStackIdentifier, fileName string, data []byte) (string, error)
		StoreRegistryManagementFileFromBytes(folder, fileName string, data []byte) (string, error)
		KeyPairFilesExist() (bool, error)
		StoreKeyPair(private, public []byte, privatePEMHeader, publicPEMHeader string) error
		LoadKeyPair() ([]byte, []byte, error)
		WriteJSONToFile(path string, content interface{}) error
		FileExists(path string) (bool, error)
		StoreEdgeJobFileFromBytes(identifier string, data []byte) (string, error)
		GetEdgeJobFolder(identifier string) string
		ClearEdgeJobTaskLogs(edgeJobID, taskID string) error
		GetEdgeJobTaskLogFileContent(edgeJobID, taskID string) (string, error)
		StoreEdgeJobTaskLogFileFromBytes(edgeJobID, taskID string, data []byte) error
		GetBinaryFolder() string
		StoreCustomTemplateFileFromBytes(identifier, fileName string, data []byte) (string, error)
		GetCustomTemplateProjectPath(identifier string) string
		GetTemporaryPath() (string, error)
	}

	// GitService represents a service for managing Git
	GitService interface {
		ClonePublicRepository(repositoryURL, referenceName string, destination string) error
		ClonePrivateRepositoryWithBasicAuth(repositoryURL, referenceName string, destination, username, password string) error
	}

	// JWTService represents a service for managing JWT tokens
	JWTService interface {
		GenerateToken(data *TokenData) (string, error)
		ParseAndVerifyToken(token string) (*TokenData, error)
		SetUserSessionDuration(userSessionDuration time.Duration)
	}

	// KubeClient represents a service used to query a Kubernetes environment
	KubeClient interface {
		SetupUserServiceAccount(userID int, teamIDs []int) error
		GetServiceAccountBearerToken(userID int) (string, error)
		StartExecProcess(namespace, podName, containerName string, command []string, stdin io.Reader, stdout io.Writer) error
	}

	// KubernetesDeployer represents a service to deploy a manifest inside a Kubernetes endpoint
	KubernetesDeployer interface {
		Deploy(endpoint *Endpoint, data string, composeFormat bool, namespace string) ([]byte, error)
	}

	// KubernetesSnapshotter represents a service used to create Kubernetes endpoint snapshots
	KubernetesSnapshotter interface {
		CreateSnapshot(endpoint *Endpoint) (*KubernetesSnapshot, error)
	}

	// LDAPService represents a service used to authenticate users against a LDAP/AD
	LDAPService interface {
		AuthenticateUser(username, password string, settings *LDAPSettings) error
		TestConnectivity(settings *LDAPSettings) error
		GetUserGroups(username string, settings *LDAPSettings) ([]string, error)
	}

	// OAuthService represents a service used to authenticate users using OAuth
	OAuthService interface {
		Authenticate(code string, configuration *OAuthSettings) (string, error)
	}

	// RegistryService represents a service for managing registry data
	RegistryService interface {
		Registry(ID RegistryID) (*Registry, error)
		Registries() ([]Registry, error)
		CreateRegistry(registry *Registry) error
		UpdateRegistry(ID RegistryID, registry *Registry) error
		DeleteRegistry(ID RegistryID) error
	}

	// ResourceControlService represents a service for managing resource control data
	ResourceControlService interface {
		ResourceControl(ID ResourceControlID) (*ResourceControl, error)
		ResourceControlByResourceIDAndType(resourceID string, resourceType ResourceControlType) (*ResourceControl, error)
		ResourceControls() ([]ResourceControl, error)
		CreateResourceControl(rc *ResourceControl) error
		UpdateResourceControl(ID ResourceControlID, resourceControl *ResourceControl) error
		DeleteResourceControl(ID ResourceControlID) error
	}

	// ReverseTunnelService represensts a service used to manage reverse tunnel connections.
	ReverseTunnelService interface {
		StartTunnelServer(addr, port string, snapshotService SnapshotService) error
		GenerateEdgeKey(url, host string, endpointIdentifier int) string
		SetTunnelStatusToActive(endpointID EndpointID)
		SetTunnelStatusToRequired(endpointID EndpointID) error
		SetTunnelStatusToIdle(endpointID EndpointID)
		GetTunnelDetails(endpointID EndpointID) *TunnelDetails
		AddEdgeJob(endpointID EndpointID, edgeJob *EdgeJob)
		RemoveEdgeJob(edgeJobID EdgeJobID)
	}

	// RoleService represents a service for managing user roles
	RoleService interface {
		Role(ID RoleID) (*Role, error)
		Roles() ([]Role, error)
		CreateRole(role *Role) error
		UpdateRole(ID RoleID, role *Role) error
	}

	// SettingsService represents a service for managing application settings
	SettingsService interface {
		Settings() (*Settings, error)
		UpdateSettings(settings *Settings) error
	}

	// Server defines the interface to serve the API
	Server interface {
		Start() error
	}

	// StackService represents a service for managing stack data
	StackService interface {
		Stack(ID StackID) (*Stack, error)
		StackByName(name string) (*Stack, error)
		Stacks() ([]Stack, error)
		CreateStack(stack *Stack) error
		UpdateStack(ID StackID, stack *Stack) error
		DeleteStack(ID StackID) error
		GetNextIdentifier() int
	}

	// StackService represents a service for managing endpoint snapshots
	SnapshotService interface {
		Start()
		SetSnapshotInterval(snapshotInterval string) error
		SnapshotEndpoint(endpoint *Endpoint) error
	}

	// SwarmStackManager represents a service to manage Swarm stacks
	SwarmStackManager interface {
		Login(dockerhub *DockerHub, registries []Registry, endpoint *Endpoint)
		Logout(endpoint *Endpoint) error
		Deploy(stack *Stack, prune bool, endpoint *Endpoint) error
		Remove(stack *Stack, endpoint *Endpoint) error
	}

	// TagService represents a service for managing tag data
	TagService interface {
		Tags() ([]Tag, error)
		Tag(ID TagID) (*Tag, error)
		CreateTag(tag *Tag) error
		UpdateTag(ID TagID, tag *Tag) error
		DeleteTag(ID TagID) error
	}

	// TeamService represents a service for managing user data
	TeamService interface {
		Team(ID TeamID) (*Team, error)
		TeamByName(name string) (*Team, error)
		Teams() ([]Team, error)
		CreateTeam(team *Team) error
		UpdateTeam(ID TeamID, team *Team) error
		DeleteTeam(ID TeamID) error
	}

	// TeamMembershipService represents a service for managing team membership data
	TeamMembershipService interface {
		TeamMembership(ID TeamMembershipID) (*TeamMembership, error)
		TeamMemberships() ([]TeamMembership, error)
		TeamMembershipsByUserID(userID UserID) ([]TeamMembership, error)
		TeamMembershipsByTeamID(teamID TeamID) ([]TeamMembership, error)
		CreateTeamMembership(membership *TeamMembership) error
		UpdateTeamMembership(ID TeamMembershipID, membership *TeamMembership) error
		DeleteTeamMembership(ID TeamMembershipID) error
		DeleteTeamMembershipByUserID(userID UserID) error
		DeleteTeamMembershipByTeamID(teamID TeamID) error
	}

	// TunnelServerService represents a service for managing data associated to the tunnel server
	TunnelServerService interface {
		Info() (*TunnelServerInfo, error)
		UpdateInfo(info *TunnelServerInfo) error
	}

	// UserService represents a service for managing user data
	UserService interface {
		User(ID UserID) (*User, error)
		UserByUsername(username string) (*User, error)
		Users() ([]User, error)
		UsersByRole(role UserRole) ([]User, error)
		CreateUser(user *User) error
		UpdateUser(ID UserID, user *User) error
		DeleteUser(ID UserID) error
	}

	// VersionService represents a service for managing version data
	VersionService interface {
		DBVersion() (int, error)
		StoreDBVersion(version int) error
		InstanceID() (string, error)
		StoreInstanceID(ID string) error
	}

	// WebhookService represents a service for managing webhook data.
	WebhookService interface {
		Webhooks() ([]Webhook, error)
		Webhook(ID WebhookID) (*Webhook, error)
		CreateWebhook(portainer *Webhook) error
		WebhookByResourceID(resourceID string) (*Webhook, error)
		WebhookByToken(token string) (*Webhook, error)
		DeleteWebhook(serviceID WebhookID) error
	}
)

const (
	// APIVersion is the version number of the Portainer API
	APIVersion = "2.1.0"
	// DBVersion is the version number of the Portainer database
	DBVersion = 26
	// ComposeSyntaxMaxVersion is a maximum supported version of the docker compose syntax
	ComposeSyntaxMaxVersion = "3.9"
	// AssetsServerURL represents the URL of the Portainer asset server
	AssetsServerURL = "https://portainer-io-assets.sfo2.digitaloceanspaces.com"
	// MessageOfTheDayURL represents the URL where Portainer MOTD message can be retrieved
	MessageOfTheDayURL = AssetsServerURL + "/motd.json"
	// VersionCheckURL represents the URL used to retrieve the latest version of Portainer
	VersionCheckURL = "https://api.github.com/repos/portainer/portainer/releases/latest"
	// PortainerAgentHeader represents the name of the header available in any agent response
	PortainerAgentHeader = "Portainer-Agent"
	// PortainerAgentEdgeIDHeader represent the name of the header containing the Edge ID associated to an agent/agent cluster
	PortainerAgentEdgeIDHeader = "X-PortainerAgent-EdgeID"
	// HTTPResponseAgentPlatform represents the name of the header containing the Agent platform
	HTTPResponseAgentPlatform = "Portainer-Agent-Platform"
	// PortainerAgentTargetHeader represent the name of the header containing the target node name
	PortainerAgentTargetHeader = "X-PortainerAgent-Target"
	// PortainerAgentSignatureHeader represent the name of the header containing the digital signature
	PortainerAgentSignatureHeader = "X-PortainerAgent-Signature"
	// PortainerAgentPublicKeyHeader represent the name of the header containing the public key
	PortainerAgentPublicKeyHeader = "X-PortainerAgent-PublicKey"
	// PortainerAgentKubernetesSATokenHeader represent the name of the header containing a Kubernetes SA token
	PortainerAgentKubernetesSATokenHeader = "X-PortainerAgent-SA-Token"
	// PortainerAgentSignatureMessage represents the message used to create a digital signature
	// to be used when communicating with an agent
	PortainerAgentSignatureMessage = "Portainer-App"
	// DefaultEdgeAgentCheckinIntervalInSeconds represents the default interval (in seconds) used by Edge agents to checkin with the Portainer instance
	DefaultEdgeAgentCheckinIntervalInSeconds = 5
	// DefaultTemplatesURL represents the URL to the official templates supported by Portainer
	DefaultTemplatesURL = "https://raw.githubusercontent.com/portainer/templates/master/templates-2.0.json"
	// DefaultUserSessionTimeout represents the default timeout after which the user session is cleared
	DefaultUserSessionTimeout = "8h"
)

const (
	_ AuthenticationMethod = iota
	// AuthenticationInternal represents the internal authentication method (authentication against Portainer API)
	AuthenticationInternal
	// AuthenticationLDAP represents the LDAP authentication method (authentication against a LDAP server)
	AuthenticationLDAP
	//AuthenticationOAuth represents the OAuth authentication method (authentication against a authorization server)
	AuthenticationOAuth
)

const (
	_ AgentPlatform = iota
	// AgentPlatformDocker represent the Docker platform (Standalone/Swarm)
	AgentPlatformDocker
	// AgentPlatformKubernetes represent the Kubernetes platform
	AgentPlatformKubernetes
)

const (
	_ EdgeJobLogsStatus = iota
	// EdgeJobLogsStatusIdle represents an idle log collection job
	EdgeJobLogsStatusIdle
	// EdgeJobLogsStatusPending represents a pending log collection job
	EdgeJobLogsStatusPending
	// EdgeJobLogsStatusCollected represents a completed log collection job
	EdgeJobLogsStatusCollected
)

const (
	_ CustomTemplatePlatform = iota
	// CustomTemplatePlatformLinux represents a custom template for linux
	CustomTemplatePlatformLinux
	// CustomTemplatePlatformWindows represents a custom template for windows
	CustomTemplatePlatformWindows
)

const (
	_ EdgeStackStatusType = iota
	//StatusOk represents a successfully deployed edge stack
	StatusOk
	//StatusError represents an edge endpoint which failed to deploy its edge stack
	StatusError
	//StatusAcknowledged represents an acknowledged edge stack
	StatusAcknowledged
)

const (
	_ EndpointExtensionType = iota
	// StoridgeEndpointExtension represents the Storidge extension
	StoridgeEndpointExtension
)

const (
	_ EndpointStatus = iota
	// EndpointStatusUp is used to represent an available endpoint
	EndpointStatusUp
	// EndpointStatusDown is used to represent an unavailable endpoint
	EndpointStatusDown
)

const (
	_ EndpointType = iota
	// DockerEnvironment represents an endpoint connected to a Docker environment
	DockerEnvironment
	// AgentOnDockerEnvironment represents an endpoint connected to a Portainer agent deployed on a Docker environment
	AgentOnDockerEnvironment
	// AzureEnvironment represents an endpoint connected to an Azure environment
	AzureEnvironment
	// EdgeAgentOnDockerEnvironment represents an endpoint connected to an Edge agent deployed on a Docker environment
	EdgeAgentOnDockerEnvironment
	// KubernetesLocalEnvironment represents an endpoint connected to a local Kubernetes environment
	KubernetesLocalEnvironment
	// AgentOnKubernetesEnvironment represents an endpoint connected to a Portainer agent deployed on a Kubernetes environment
	AgentOnKubernetesEnvironment
	// EdgeAgentOnKubernetesEnvironment represents an endpoint connected to an Edge agent deployed on a Kubernetes environment
	EdgeAgentOnKubernetesEnvironment
)

const (
	_ JobType = iota
	// SnapshotJobType is a system job used to create endpoint snapshots
	SnapshotJobType = 2
)

const (
	_ MembershipRole = iota
	// TeamLeader represents a leader role inside a team
	TeamLeader
	// TeamMember represents a member role inside a team
	TeamMember
)

const (
	_ RegistryType = iota
	// QuayRegistry represents a Quay.io registry
	QuayRegistry
	// AzureRegistry represents an ACR registry
	AzureRegistry
	// CustomRegistry represents a custom registry
	CustomRegistry
	// GitlabRegistry represents a gitlab registry
	GitlabRegistry
)

const (
	_ ResourceAccessLevel = iota
	// ReadWriteAccessLevel represents an access level with read-write permissions on a resource
	ReadWriteAccessLevel
)

const (
	_ ResourceControlType = iota
	// ContainerResourceControl represents a resource control associated to a Docker container
	ContainerResourceControl
	// ServiceResourceControl represents a resource control associated to a Docker service
	ServiceResourceControl
	// VolumeResourceControl represents a resource control associated to a Docker volume
	VolumeResourceControl
	// NetworkResourceControl represents a resource control associated to a Docker network
	NetworkResourceControl
	// SecretResourceControl represents a resource control associated to a Docker secret
	SecretResourceControl
	// StackResourceControl represents a resource control associated to a stack composed of Docker services
	StackResourceControl
	// ConfigResourceControl represents a resource control associated to a Docker config
	ConfigResourceControl
	// CustomTemplateResourceControl  represents a resource control associated to a custom template
	CustomTemplateResourceControl
)

const (
	_ StackType = iota
	// DockerSwarmStack represents a stack managed via docker stack
	DockerSwarmStack
	// DockerComposeStack represents a stack managed via docker-compose
	DockerComposeStack
	// KubernetesStack represents a stack managed via kubectl
	KubernetesStack
)

// StackStatus represents a status for a stack
const (
	_ StackStatus = iota
	StackStatusActive
	StackStatusInactive
)

const (
	_ TemplateType = iota
	// ContainerTemplate represents a container template
	ContainerTemplate
	// SwarmStackTemplate represents a template used to deploy a Swarm stack
	SwarmStackTemplate
	// ComposeStackTemplate represents a template used to deploy a Compose stack
	ComposeStackTemplate
	// EdgeStackTemplate represents a template used to deploy an Edge stack
	EdgeStackTemplate
)

const (
	// TLSFileCA represents a TLS CA certificate file
	TLSFileCA TLSFileType = iota
	// TLSFileCert represents a TLS certificate file
	TLSFileCert
	// TLSFileKey represents a TLS key file
	TLSFileKey
)

const (
	_ UserRole = iota
	// AdministratorRole represents an administrator user role
	AdministratorRole
	// StandardUserRole represents a regular user role
	StandardUserRole
)

const (
	_ WebhookType = iota
	// ServiceWebhook is a webhook for restarting a docker service
	ServiceWebhook
)

const (
	// EdgeAgentIdle represents an idle state for a tunnel connected to an Edge endpoint.
	EdgeAgentIdle string = "IDLE"
	// EdgeAgentManagementRequired represents a required state for a tunnel connected to an Edge endpoint
	EdgeAgentManagementRequired string = "REQUIRED"
	// EdgeAgentActive represents an active state for a tunnel connected to an Edge endpoint
	EdgeAgentActive string = "ACTIVE"
)

const (
	OperationDockerContainerArchiveInfo         Authorization = "DockerContainerArchiveInfo"
	OperationDockerContainerList                Authorization = "DockerContainerList"
	OperationDockerContainerExport              Authorization = "DockerContainerExport"
	OperationDockerContainerChanges             Authorization = "DockerContainerChanges"
	OperationDockerContainerInspect             Authorization = "DockerContainerInspect"
	OperationDockerContainerTop                 Authorization = "DockerContainerTop"
	OperationDockerContainerLogs                Authorization = "DockerContainerLogs"
	OperationDockerContainerStats               Authorization = "DockerContainerStats"
	OperationDockerContainerAttachWebsocket     Authorization = "DockerContainerAttachWebsocket"
	OperationDockerContainerArchive             Authorization = "DockerContainerArchive"
	OperationDockerContainerCreate              Authorization = "DockerContainerCreate"
	OperationDockerContainerPrune               Authorization = "DockerContainerPrune"
	OperationDockerContainerKill                Authorization = "DockerContainerKill"
	OperationDockerContainerPause               Authorization = "DockerContainerPause"
	OperationDockerContainerUnpause             Authorization = "DockerContainerUnpause"
	OperationDockerContainerRestart             Authorization = "DockerContainerRestart"
	OperationDockerContainerStart               Authorization = "DockerContainerStart"
	OperationDockerContainerStop                Authorization = "DockerContainerStop"
	OperationDockerContainerWait                Authorization = "DockerContainerWait"
	OperationDockerContainerResize              Authorization = "DockerContainerResize"
	OperationDockerContainerAttach              Authorization = "DockerContainerAttach"
	OperationDockerContainerExec                Authorization = "DockerContainerExec"
	OperationDockerContainerRename              Authorization = "DockerContainerRename"
	OperationDockerContainerUpdate              Authorization = "DockerContainerUpdate"
	OperationDockerContainerPutContainerArchive Authorization = "DockerContainerPutContainerArchive"
	OperationDockerContainerDelete              Authorization = "DockerContainerDelete"
	OperationDockerImageList                    Authorization = "DockerImageList"
	OperationDockerImageSearch                  Authorization = "DockerImageSearch"
	OperationDockerImageGetAll                  Authorization = "DockerImageGetAll"
	OperationDockerImageGet                     Authorization = "DockerImageGet"
	OperationDockerImageHistory                 Authorization = "DockerImageHistory"
	OperationDockerImageInspect                 Authorization = "DockerImageInspect"
	OperationDockerImageLoad                    Authorization = "DockerImageLoad"
	OperationDockerImageCreate                  Authorization = "DockerImageCreate"
	OperationDockerImagePrune                   Authorization = "DockerImagePrune"
	OperationDockerImagePush                    Authorization = "DockerImagePush"
	OperationDockerImageTag                     Authorization = "DockerImageTag"
	OperationDockerImageDelete                  Authorization = "DockerImageDelete"
	OperationDockerImageCommit                  Authorization = "DockerImageCommit"
	OperationDockerImageBuild                   Authorization = "DockerImageBuild"
	OperationDockerNetworkList                  Authorization = "DockerNetworkList"
	OperationDockerNetworkInspect               Authorization = "DockerNetworkInspect"
	OperationDockerNetworkCreate                Authorization = "DockerNetworkCreate"
	OperationDockerNetworkConnect               Authorization = "DockerNetworkConnect"
	OperationDockerNetworkDisconnect            Authorization = "DockerNetworkDisconnect"
	OperationDockerNetworkPrune                 Authorization = "DockerNetworkPrune"
	OperationDockerNetworkDelete                Authorization = "DockerNetworkDelete"
	OperationDockerVolumeList                   Authorization = "DockerVolumeList"
	OperationDockerVolumeInspect                Authorization = "DockerVolumeInspect"
	OperationDockerVolumeCreate                 Authorization = "DockerVolumeCreate"
	OperationDockerVolumePrune                  Authorization = "DockerVolumePrune"
	OperationDockerVolumeDelete                 Authorization = "DockerVolumeDelete"
	OperationDockerExecInspect                  Authorization = "DockerExecInspect"
	OperationDockerExecStart                    Authorization = "DockerExecStart"
	OperationDockerExecResize                   Authorization = "DockerExecResize"
	OperationDockerSwarmInspect                 Authorization = "DockerSwarmInspect"
	OperationDockerSwarmUnlockKey               Authorization = "DockerSwarmUnlockKey"
	OperationDockerSwarmInit                    Authorization = "DockerSwarmInit"
	OperationDockerSwarmJoin                    Authorization = "DockerSwarmJoin"
	OperationDockerSwarmLeave                   Authorization = "DockerSwarmLeave"
	OperationDockerSwarmUpdate                  Authorization = "DockerSwarmUpdate"
	OperationDockerSwarmUnlock                  Authorization = "DockerSwarmUnlock"
	OperationDockerNodeList                     Authorization = "DockerNodeList"
	OperationDockerNodeInspect                  Authorization = "DockerNodeInspect"
	OperationDockerNodeUpdate                   Authorization = "DockerNodeUpdate"
	OperationDockerNodeDelete                   Authorization = "DockerNodeDelete"
	OperationDockerServiceList                  Authorization = "DockerServiceList"
	OperationDockerServiceInspect               Authorization = "DockerServiceInspect"
	OperationDockerServiceLogs                  Authorization = "DockerServiceLogs"
	OperationDockerServiceCreate                Authorization = "DockerServiceCreate"
	OperationDockerServiceUpdate                Authorization = "DockerServiceUpdate"
	OperationDockerServiceDelete                Authorization = "DockerServiceDelete"
	OperationDockerSecretList                   Authorization = "DockerSecretList"
	OperationDockerSecretInspect                Authorization = "DockerSecretInspect"
	OperationDockerSecretCreate                 Authorization = "DockerSecretCreate"
	OperationDockerSecretUpdate                 Authorization = "DockerSecretUpdate"
	OperationDockerSecretDelete                 Authorization = "DockerSecretDelete"
	OperationDockerConfigList                   Authorization = "DockerConfigList"
	OperationDockerConfigInspect                Authorization = "DockerConfigInspect"
	OperationDockerConfigCreate                 Authorization = "DockerConfigCreate"
	OperationDockerConfigUpdate                 Authorization = "DockerConfigUpdate"
	OperationDockerConfigDelete                 Authorization = "DockerConfigDelete"
	OperationDockerTaskList                     Authorization = "DockerTaskList"
	OperationDockerTaskInspect                  Authorization = "DockerTaskInspect"
	OperationDockerTaskLogs                     Authorization = "DockerTaskLogs"
	OperationDockerPluginList                   Authorization = "DockerPluginList"
	OperationDockerPluginPrivileges             Authorization = "DockerPluginPrivileges"
	OperationDockerPluginInspect                Authorization = "DockerPluginInspect"
	OperationDockerPluginPull                   Authorization = "DockerPluginPull"
	OperationDockerPluginCreate                 Authorization = "DockerPluginCreate"
	OperationDockerPluginEnable                 Authorization = "DockerPluginEnable"
	OperationDockerPluginDisable                Authorization = "DockerPluginDisable"
	OperationDockerPluginPush                   Authorization = "DockerPluginPush"
	OperationDockerPluginUpgrade                Authorization = "DockerPluginUpgrade"
	OperationDockerPluginSet                    Authorization = "DockerPluginSet"
	OperationDockerPluginDelete                 Authorization = "DockerPluginDelete"
	OperationDockerSessionStart                 Authorization = "DockerSessionStart"
	OperationDockerDistributionInspect          Authorization = "DockerDistributionInspect"
	OperationDockerBuildPrune                   Authorization = "DockerBuildPrune"
	OperationDockerBuildCancel                  Authorization = "DockerBuildCancel"
	OperationDockerPing                         Authorization = "DockerPing"
	OperationDockerInfo                         Authorization = "DockerInfo"
	OperationDockerEvents                       Authorization = "DockerEvents"
	OperationDockerSystem                       Authorization = "DockerSystem"
	OperationDockerVersion                      Authorization = "DockerVersion"

	OperationDockerAgentPing         Authorization = "DockerAgentPing"
	OperationDockerAgentList         Authorization = "DockerAgentList"
	OperationDockerAgentHostInfo     Authorization = "DockerAgentHostInfo"
	OperationDockerAgentBrowseDelete Authorization = "DockerAgentBrowseDelete"
	OperationDockerAgentBrowseGet    Authorization = "DockerAgentBrowseGet"
	OperationDockerAgentBrowseList   Authorization = "DockerAgentBrowseList"
	OperationDockerAgentBrowsePut    Authorization = "DockerAgentBrowsePut"
	OperationDockerAgentBrowseRename Authorization = "DockerAgentBrowseRename"

	OperationPortainerDockerHubInspect        Authorization = "PortainerDockerHubInspect"
	OperationPortainerDockerHubUpdate         Authorization = "PortainerDockerHubUpdate"
	OperationPortainerEndpointGroupCreate     Authorization = "PortainerEndpointGroupCreate"
	OperationPortainerEndpointGroupList       Authorization = "PortainerEndpointGroupList"
	OperationPortainerEndpointGroupDelete     Authorization = "PortainerEndpointGroupDelete"
	OperationPortainerEndpointGroupInspect    Authorization = "PortainerEndpointGroupInspect"
	OperationPortainerEndpointGroupUpdate     Authorization = "PortainerEndpointGroupEdit"
	OperationPortainerEndpointGroupAccess     Authorization = "PortainerEndpointGroupAccess "
	OperationPortainerEndpointList            Authorization = "PortainerEndpointList"
	OperationPortainerEndpointInspect         Authorization = "PortainerEndpointInspect"
	OperationPortainerEndpointCreate          Authorization = "PortainerEndpointCreate"
	OperationPortainerEndpointExtensionAdd    Authorization = "PortainerEndpointExtensionAdd"
	OperationPortainerEndpointJob             Authorization = "PortainerEndpointJob"
	OperationPortainerEndpointSnapshots       Authorization = "PortainerEndpointSnapshots"
	OperationPortainerEndpointSnapshot        Authorization = "PortainerEndpointSnapshot"
	OperationPortainerEndpointUpdate          Authorization = "PortainerEndpointUpdate"
	OperationPortainerEndpointUpdateAccess    Authorization = "PortainerEndpointUpdateAccess"
	OperationPortainerEndpointDelete          Authorization = "PortainerEndpointDelete"
	OperationPortainerEndpointExtensionRemove Authorization = "PortainerEndpointExtensionRemove"
	OperationPortainerExtensionList           Authorization = "PortainerExtensionList"
	OperationPortainerExtensionInspect        Authorization = "PortainerExtensionInspect"
	OperationPortainerExtensionCreate         Authorization = "PortainerExtensionCreate"
	OperationPortainerExtensionUpdate         Authorization = "PortainerExtensionUpdate"
	OperationPortainerExtensionDelete         Authorization = "PortainerExtensionDelete"
	OperationPortainerMOTD                    Authorization = "PortainerMOTD"
	OperationPortainerRegistryList            Authorization = "PortainerRegistryList"
	OperationPortainerRegistryInspect         Authorization = "PortainerRegistryInspect"
	OperationPortainerRegistryCreate          Authorization = "PortainerRegistryCreate"
	OperationPortainerRegistryConfigure       Authorization = "PortainerRegistryConfigure"
	OperationPortainerRegistryUpdate          Authorization = "PortainerRegistryUpdate"
	OperationPortainerRegistryUpdateAccess    Authorization = "PortainerRegistryUpdateAccess"
	OperationPortainerRegistryDelete          Authorization = "PortainerRegistryDelete"
	OperationPortainerResourceControlCreate   Authorization = "PortainerResourceControlCreate"
	OperationPortainerResourceControlUpdate   Authorization = "PortainerResourceControlUpdate"
	OperationPortainerResourceControlDelete   Authorization = "PortainerResourceControlDelete"
	OperationPortainerRoleList                Authorization = "PortainerRoleList"
	OperationPortainerRoleInspect             Authorization = "PortainerRoleInspect"
	OperationPortainerRoleCreate              Authorization = "PortainerRoleCreate"
	OperationPortainerRoleUpdate              Authorization = "PortainerRoleUpdate"
	OperationPortainerRoleDelete              Authorization = "PortainerRoleDelete"
	OperationPortainerScheduleList            Authorization = "PortainerScheduleList"
	OperationPortainerScheduleInspect         Authorization = "PortainerScheduleInspect"
	OperationPortainerScheduleFile            Authorization = "PortainerScheduleFile"
	OperationPortainerScheduleTasks           Authorization = "PortainerScheduleTasks"
	OperationPortainerScheduleCreate          Authorization = "PortainerScheduleCreate"
	OperationPortainerScheduleUpdate          Authorization = "PortainerScheduleUpdate"
	OperationPortainerScheduleDelete          Authorization = "PortainerScheduleDelete"
	OperationPortainerSettingsInspect         Authorization = "PortainerSettingsInspect"
	OperationPortainerSettingsUpdate          Authorization = "PortainerSettingsUpdate"
	OperationPortainerSettingsLDAPCheck       Authorization = "PortainerSettingsLDAPCheck"
	OperationPortainerStackList               Authorization = "PortainerStackList"
	OperationPortainerStackInspect            Authorization = "PortainerStackInspect"
	OperationPortainerStackFile               Authorization = "PortainerStackFile"
	OperationPortainerStackCreate             Authorization = "PortainerStackCreate"
	OperationPortainerStackMigrate            Authorization = "PortainerStackMigrate"
	OperationPortainerStackUpdate             Authorization = "PortainerStackUpdate"
	OperationPortainerStackDelete             Authorization = "PortainerStackDelete"
	OperationPortainerTagList                 Authorization = "PortainerTagList"
	OperationPortainerTagCreate               Authorization = "PortainerTagCreate"
	OperationPortainerTagDelete               Authorization = "PortainerTagDelete"
	OperationPortainerTeamMembershipList      Authorization = "PortainerTeamMembershipList"
	OperationPortainerTeamMembershipCreate    Authorization = "PortainerTeamMembershipCreate"
	OperationPortainerTeamMembershipUpdate    Authorization = "PortainerTeamMembershipUpdate"
	OperationPortainerTeamMembershipDelete    Authorization = "PortainerTeamMembershipDelete"
	OperationPortainerTeamList                Authorization = "PortainerTeamList"
	OperationPortainerTeamInspect             Authorization = "PortainerTeamInspect"
	OperationPortainerTeamMemberships         Authorization = "PortainerTeamMemberships"
	OperationPortainerTeamCreate              Authorization = "PortainerTeamCreate"
	OperationPortainerTeamUpdate              Authorization = "PortainerTeamUpdate"
	OperationPortainerTeamDelete              Authorization = "PortainerTeamDelete"
	OperationPortainerTemplateList            Authorization = "PortainerTemplateList"
	OperationPortainerTemplateInspect         Authorization = "PortainerTemplateInspect"
	OperationPortainerTemplateCreate          Authorization = "PortainerTemplateCreate"
	OperationPortainerTemplateUpdate          Authorization = "PortainerTemplateUpdate"
	OperationPortainerTemplateDelete          Authorization = "PortainerTemplateDelete"
	OperationPortainerUploadTLS               Authorization = "PortainerUploadTLS"
	OperationPortainerUserList                Authorization = "PortainerUserList"
	OperationPortainerUserInspect             Authorization = "PortainerUserInspect"
	OperationPortainerUserMemberships         Authorization = "PortainerUserMemberships"
	OperationPortainerUserCreate              Authorization = "PortainerUserCreate"
	OperationPortainerUserUpdate              Authorization = "PortainerUserUpdate"
	OperationPortainerUserUpdatePassword      Authorization = "PortainerUserUpdatePassword"
	OperationPortainerUserDelete              Authorization = "PortainerUserDelete"
	OperationPortainerWebsocketExec           Authorization = "PortainerWebsocketExec"
	OperationPortainerWebhookList             Authorization = "PortainerWebhookList"
	OperationPortainerWebhookCreate           Authorization = "PortainerWebhookCreate"
	OperationPortainerWebhookDelete           Authorization = "PortainerWebhookDelete"

	OperationIntegrationStoridgeAdmin Authorization = "IntegrationStoridgeAdmin"

	OperationDockerUndefined      Authorization = "DockerUndefined"
	OperationDockerAgentUndefined Authorization = "DockerAgentUndefined"
	OperationPortainerUndefined   Authorization = "PortainerUndefined"

	EndpointResourcesAccess Authorization = "EndpointResourcesAccess"
)
