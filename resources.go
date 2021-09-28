package insightcloudsec

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// STRUCTS
///////////////////////////////////////////
type Query struct {
	Scopes               []string `json:"scopes"`
	Filters              []string `json:"filters"`
	Offset               int      `json:"offset"`
	Limit                int      `json:"limit"`
	OrderBy              string   `json:"order_by"`
	SelectedResourceType string   `json:"selected_resource_type"`
}

type QueryResult struct {
	Scopes               []string       `json:"scopes"`
	Filters              []QueryFilter  `json:"filters"`
	Offset               int            `json:"offset"`
	OrderBy              string         `json:"order_by"`
	Counts               map[string]int `json:"counts"`
	SelectedResourceType string         `json:"selected_resource_type"`
	Resources            []Resource     `json:"resources"`
	SupportedTypes       []string       `json:"supported_resources"`
}

type QueryFilter struct {
	Config QueryFilterConfig
	Name   string
}

type QueryFilterConfig struct {
}

type Resource struct {
	Type                    string                          `json:"resource_type"`
	Instance                InstanceResource                `json:"instance,omitempty"`
	Volume                  VolumeResource                  `json:"volume,omitempty"`
	AccessFlowLog           AccessFlowLogResource           `json:"accesslistflowlog,omitempty"`
	AccessAnalyzer          AccessAnalyzerResource          `json:"accessanalyzer,omitempty"`
	AirflowEnvironment      AirflowEnvironmentResource      `json:"airflowenvironment,omitempty"`
	APIAccountingConfig     APIAccountingConfigResource     `json:"apiaccountingconfig,omitempty"`
	AppRunnerService        AppRunnerServiceResource        `json:"apprunnerservice,omitempty"`
	AppServer               AppServerResource               `json:"appserver,omitempty"`
	AutoscalingGroup        AutoscalingGroupResource        `json:"autoscalinggroup,omitempty"`
	AutoscalingLaunchConfig AutoscalingLaunchConfigResource `json:"autoscalinglaunchconfiguration,omitempty"`
	AWSConfig               AWSConfigResource               `json:"awsconfig,omitempty"`
	BackendService          BackendServiceResource          `json:"backendservice,omitempty"`
	BackupVault             BackupVaultResource             `json:"backupvault,omitempty"`
	BatchEnvironment        BatchEnvironmentResource        `json:"batchenvironment,omitempty"`
	BatchPool               BatchPoolResource               `json:"batchpool,omitempty"`
}

type CommonResourceValues struct {
	//Common Attributes
	ID               string            `json:"resource_id"`
	Name             string            `json:"resource_name"`
	Type             string            `json:"type"`
	Cloud            string            `json:"cloud"`
	Account          string            `json:"account"`
	AccountID        string            `json:"account_id"`
	AccountStatus    string            `json:"account_status"`
	OrgServiceID     int               `json:"organization_service_id"`
	AvailabilityZone string            `json:"availablility_zone"`
	Region           string            `json:"region"`
	Created          string            `json:"creation_timestamp"`
	Discovered       string            `json:"discovered_timestamp"`
	Modified         string            `json:"modified_timestamp"`
	NamespaceID      string            `json:"namespace_id"`
	Tags             map[string]string `json:"tags"`
}

type AccessAnalyzerResource struct {
	Common                   CommonResourceValues `json:"common"`
	ID                       string               `json:"analyzer_id"`
	Mapping                  string               `json:"account_mapping"`
	ActiveFindingCount       int                  `json:"active_finding_count"`
	PublicFindingCount       int                  `json:"public_finding_count"`
	CrossAccountFindingCount int                  `json:"cross_account_count"`
	ThirdPartyFindingCount   int                  `json:"unknown_account_count"`
}

type AccessFlowLogResource struct {
	Common                   CommonResourceValues `json:"common"`
	Name                     string               `json:"name"`
	RegionName               string               `json:"region_name"`
	StorageID                string               `json:"storage_id"`
	TargetResourceID         string               `json:"target_resource_id"`
	ProvisioningState        string               `json:"provisioning_state"`
	Enabled                  bool                 `json:"enabled"`
	Retention                bool                 `json:"retention_enabled"`
	RetentionTime            string               `json:"retention_time"`
	TrafficAnalytics         bool                 `json:"traffic_analytics_enabled"`
	TrafficAnalyticsInterval string               `json:"raffic_analytics_interval"`
	Version                  string               `json:"version"`
}

type AirflowEnvironmentResource struct {
	Common              CommonResourceValues `json:"common"`
	WebserverAccessMode string               `json:"webserver_access_mode"`
	WebserverURL        string               `json:"webserver_url"`
	Status              string               `json:"status"`
	Class               string               `json:"environment_class"`
	MaxWorkers          int                  `json:"max_workers"`
	Encrypted           bool                 `json:"encrypted"`
	EncryptionKeyID     string               `json:"key_resource_id"`
	LoggingConfig       string               `json:"logging_config"`
	ExecutionRoleID     string               `json:"execution_role_resource_id"`
	ServiceRoleID       string               `json:"service_role_resource_id"`
}

type APIAccountingConfigResource struct {
	Common                     CommonResourceValues `json:"common"`
	ID                         string               `json:"accounting_config_id"`
	Name                       string               `json:"name"`
	MultiRegion                bool                 `json:"multi_region"`
	ParentResourceID           string               `json:"parent_resource_id"`
	IsLogging                  bool                 `json:"is_logging"`
	IsOrgTrail                 bool                 `json:"is_organization_trail"`
	IncludeGlobalServiceEvents bool                 `json:"include_global_service_events"`
	StorageContainerName       string               `json:"storage_container_name"`
	ManagementEvents           string               `json:"management_events,omitempty"`
	SNSTopic                   string               `json:"sns_topic_name"`
	LogGroup                   string               `json:"cloud_watch_group_arn"`
	RetentionDays              int                  `json:"retention_days"`
	Locked                     bool                 `json:"locked"`
}

type AppRunnerServiceResource struct {
	Common         CommonResourceValues `json:"commmon"`
	ID             string               `json:"service_id"`
	Status         string               `json:"status"`
	Repository     string               `json:"repository"`
	RepositoryType string               `json:"repository_type"`
	Cores          int                  `json:"cores"`
	Memory         int                  `json:"memory"`
}

type AppServerResource struct {
	Common           CommonResourceValues `json:"commmon"`
	ID               string               `json:"app_server_id"`
	State            string               `json:"state"`
	Type             string               `json:"server_type"`
	InstanceCount    int                  `json:"instance_count"`
	MaxInstanceCount int                  `json:"max_instance_count"`
	AppCount         int                  `json:"app_count"`
}

type AutoscalingGroupResource struct {
	Common                 CommonResourceValues `json:"commmon"`
	ID                     string               `json:"group_id"`
	Created                string               `json:"create_time"`
	HealthCheckGracePeriod int                  `json:"health_check_grace_period"`
	MultiAZ                bool                 `json:"multi_az"`
	MinSize                int                  `json:"min_size"`
	MaxSize                int                  `json:"max_size"`
	DesiredCapacity        int                  `json:"desired_capacity"`
	NewInstanceProtection  bool                 `json:"new_instance_protection"`
	DefaultCooldown        int                  `json:"default_cooldown"`
	UpgradePolicy          string               `json:"upgrade_policy"`
	SuspendedProcesses     []string             `json:"suspended_processes"`
}

type AutoscalingLaunchConfigResource struct {
	Common                CommonResourceValues `json:"commmon"`
	Name                  string               `json:"name"`
	ImageID               string               `json:"image_id"`
	InstanceType          string               `json:"instance_type"`
	IAMRole               string               `json:"identity_management_role"`
	Region                string               `json:"region_name"`
	Created               string               `json:"create_time"`
	Monitoring            bool                 `json:"monitoring"`
	BlockStorageOptimized bool                 `json:"block_storage_optimized"`
	AssociateIP           string               `json:"associate_ip"`
	RAMID                 string               `json:"ram_id"`
	KernelID              string               `json:"kernel_id"`
}

type AWSConfigResource struct {
	Common                      CommonResourceValues `json:"commmon"`
	ID                          string               `json:"resource_id"`
	DeliveryChannelCreated      bool                 `json:"delivery_channel_created"`
	ConfiurationRecorderCreated bool                 `json:"configuration_recorder_created"`
	AuditingBegun               bool                 `json:"auditing_has_begun"`
	AuditingEnabled             bool                 `json:"auditing_enabled"`
	CrossAccount                bool                 `json:"cross_account"`
	UnknownAccount              bool                 `json:"unknown_account"`
}

type BackendServiceResource struct {
	Common             CommonResourceValues `json:"common"`
	Kind               string               `json:"kind"`
	StorageContainerID string               `json:"storage_container_resource_id"`
	PortName           string               `json:"port_name"`
	Port               string               `json:"port"`
	Created            string               `json:"created_time"`
	Scheme             string               `json:"scheme"`
}

type BackupVaultResource struct {
	Common          CommonResourceValues `json:"common"`
	Name            string               `json:"name"`
	Created         string               `json:"create_time"`
	RecoveryPoints  int                  `json:"recovery_points"`
	Policy          string               `json:"policy"`
	Public          bool                 `json:"public"`
	EncryptionKeyID string               `json:"key_resource_id"`
}

type BatchEnvironmentResource struct {
	Common         CommonResourceValues `json:"common"`
	Name           string               `json:"name"`
	Region         string               `json:"region_name"`
	Endpoint       string               `json:"endpoint"`
	State          string               `json:"state"`
	AllocationType string               `json:"allocation_type"`
	PublicAccess   bool                 `json:"public_access"`
	MinVCPUs       int                  `json:"minimum_cpus"`
	MaxVCPUs       int                  `json:"maximum_cpus"`
	PoolType       string               `json:"pool_type"`
}

type BatchPoolResource struct {
	Common                 CommonResourceValues `json:"common"`
	Name                   string               `json:"name"`
	Region                 string               `json:"region"`
	State                  string               `json:"state"`
	InstanceSize           string               `json:"vm_size"`
	Autoscaling            string               `json:"autoscaling"`
	InterNodeCommunication string               `json:"inter_node_communication"`
}

type InstanceResource struct {
	Common                                         CommonResourceValues `json:"common"`
	ID                                             string               `json:"instance_id"`
	Type                                           string               `json:"instance_type"`
	LaunchTime                                     string               `json:"launch_time"`
	Platform                                       string               `json:"platform"`
	State                                          string               `json:"state"`
	ImageID                                        string               `json:"image_id"`
	PublicIPAddress                                string               `json:"public_ip_address,omitempty"`
	PrivateIPAddress                               string               `json:"private_ip_address"`
	NetworkResourceID                              string               `json:"network_resource_id"`
	SubnetResourceID                               string               `json:"subnet_resource_id"`
	ObjectID                                       string               `json:"object_id"`
	KeyName                                        string               `json:"key_name,omitempty"`
	RoleName                                       string               `json:"role_name,omitempty"`
	TerminationProtection                          string               `json:"termination_protection,omitempty"`
	VMExtensions                                   string               `json:"vm_extensions,omitempty"`
	JITAccessPolicy                                string               `json:"jit_access_policy"`
	RootDeviceType                                 string               `json:"root_device_type,omitempty"`
	Tenancy                                        string               `json:"tenancy,omitempty"`
	DetailedMonitoring                             bool                 `json:"detailed_monitoring,omitempty"`
	SecondaryPrivateIPAddresses                    []string             `json:"secondary_private_ip_addresses,omitempty"`
	SecondaryPublicIPAddresses                     []string             `json:"secondary_public_ip_addresses,omitempty"`
	Architecture                                   string               `json:"architecture,omitempty"`
	AWSInstanceMetadataServiceV2Required           bool                 `json:"aws_instance_metadata_service_v2_required,omitempty"`
	AWSInstanceMetadataServiceHopLimit             int                  `json:"aws_instance_metadata_hop_limit,omitempty"`
	AWSInstanceMetadataServiceEndpointEnabled      string               `json:"aws_instance_metadata_service_endpoint_enabled,omitempty"`
	AWSInstanceMetadataServiceEndpointConfigStatus string               `json:"aws_instance_metadata_service_endpoint_config_status,omitempty"`
	StateTransitionReason                          string               `json:"state_transition_reason,omitempty"`
}

type VolumeResource struct {
	Common              CommonResourceValues `json:"common"`
	ID                  string               `json:"volume_id"`
	Type                string               `json:"volume_type"`
	Size                int                  `json:"size"`
	State               string               `json:"state"`
	IOPS                string               `json:"rated_iops,omitempty"`
	Encrypted           bool                 `json:"encrypted"`
	DeleteOnTermination bool                 `json:"delete_on_termination"`
	AttachState         string               `json:"attach_state"`
	AttachDeviceName    string               `json:"attach_device_name,omitempty"`
	InstanceAssociation InstanceAssociation  `json:"instance_association"`
	Created             string               `json:"creation_time"`
}

type InstanceAssociation struct {
	Common InstanceAssociationDetails `json:"common"`
}

type InstanceAssociationDetails struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
}

// QUERY FUNCTIONS
///////////////////////////////////////////
func (c Client) Query(q *Query) (*QueryResult, error) {
	if q.Filters == nil {
		q.Filters = make([]string, 0)
	}
	if q.Scopes == nil {
		q.Scopes = make([]string, 0)
	}

	data, err := json.Marshal(q)
	if err != nil {
		return nil, err
	}

	resp, err := c.makeRequest(http.MethodPost, "/v2/public/resource/query", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	var ret *QueryResult
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// RESOURCE FUNCTIONS
///////////////////////////////////////////