package api

// Cluster represents high-level information about a LXD cluster.
//
// swagger:model
//
// API extension: clustering
type Cluster struct {
	// Name of the cluster member answering the request
	// Example: lxd01
	ServerName string `json:"server_name" yaml:"server_name"`

	// Whether clustering is enabled
	// Example: true
	Enabled bool `json:"enabled" yaml:"enabled"`

	// List of member configuration keys (used during join)
	//
	// API extension: clustering_join
	MemberConfig []ClusterMemberConfigKey `json:"member_config" yaml:"member_config"`
}

// ClusterMemberConfigKey represents a single config key that a new member of
// the cluster is required to provide when joining.
//
// The Value field is empty when getting clustering information with GET
// /1.0/cluster, and should be filled by the joining node when performing a PUT
// /1.0/cluster join request.
//
// swagger:model
//
// API extension: clustering_join
type ClusterMemberConfigKey struct {
	// The kind of configuration key (network, storage-pool, ...)
	// Example: storage-pool
	Entity string `json:"entity" yaml:"entity"`

	// The name of the object requiring this key
	// Example: local
	Name string `json:"name" yaml:"name"`

	// The name of the key
	// Example: source
	Key string `json:"key" yaml:"key"`

	// The value on the answering cluster member
	// Example: /dev/sdb
	Value string `json:"value" yaml:"value"`

	// A human friendly description key
	// Example: "source" property for storage pool "local"
	Description string `json:"description" yaml:"description"`
}

// ClusterPut represents the fields required to bootstrap or join a LXD
// cluster.
//
// swagger:model
//
// API extension: clustering
type ClusterPut struct {
	Cluster `yaml:",inline"`

	// The address of the cluster you wish to join
	// Example: 10.0.0.1:8443
	ClusterAddress string `json:"cluster_address" yaml:"cluster_address"`

	// The expected certificate (X509 PEM encoded) for the cluster
	// Example: X509 PEM certificate
	ClusterCertificate string `json:"cluster_certificate" yaml:"cluster_certificate"`

	// The local address to use for cluster communication
	// Example: 10.0.0.2:8443
	//
	// API extension: clustering_join
	ServerAddress string `json:"server_address" yaml:"server_address"`

	// The trust password of the cluster you're trying to join
	// Example: blah
	//
	// API extension: clustering_join
	ClusterPassword string `json:"cluster_password" yaml:"cluster_password"`
}

// ClusterMemberPost represents the fields required to rename a LXD node.
//
// swagger:model
//
// API extension: clustering
type ClusterMemberPost struct {
	// The new name of the cluster member
	// Example: lxd02
	ServerName string `json:"server_name" yaml:"server_name"`
}

// ClusterMember represents the a LXD node in the cluster.
//
// swagger:model
//
// API extension: clustering
type ClusterMember struct {
	ClusterMemberPut `yaml:",inline"`

	// Name of the cluster member
	// Example: lxd01
	ServerName string `json:"server_name" yaml:"server_name"`

	// URL at which the cluster member can be reached
	// Example: https://10.0.0.1:8443
	URL string `json:"url" yaml:"url"`

	// Whether the cluster member is a database server
	// Example: true
	Database bool `json:"database" yaml:"database"`

	// Current status
	// Example: Online
	Status string `json:"status" yaml:"status"`

	// Additional status information
	// Example: fully operational
	Message string `json:"message" yaml:"message"`

	// The primary architecture of the cluster member
	//
	// API extension: clustering_architecture
	Architecture string `json:"architecture" yaml:"architecture"`
}

// Writable converts a full Profile struct into a ProfilePut struct (filters read-only fields)
func (member *ClusterMember) Writable() ClusterMemberPut {
	return member.ClusterMemberPut
}

// ClusterMemberPut represents the the modifiable fields of a LXD cluster member
//
// swagger:model
//
// API extension: clustering_edit_roles
type ClusterMemberPut struct {
	// List of roles held by this cluster member
	// Example: ["database"]
	//
	// API extension: clustering_roles
	Roles []string `json:"roles" yaml:"roles"`

	// Name of the failure domain for this cluster member
	// Example: rack1
	//
	// API extension: clustering_failure_domains
	FailureDomain string `json:"failure_domain" yaml:"failure_domain"`
}
