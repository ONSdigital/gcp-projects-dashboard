package response

type (

	// Response represents a JSON response containing GCP project and GKE cluster details.
	Response struct {
		Project Project `json:"project"`
		Cluster Cluster `json:"cluster"`
	}

	// Project represents a GCP project.
	Project struct {
		ID         string            `json:"id,omitempty"`
		Labels     map[string]string `json:"labels,omitempty"`
		Name       string            `json:"name,omitempty"`
		Number     int64             `json:"number,omitempty,string"`
		CreateTime string            `json:"createTime,omitempty"`
	}

	// Cluster represents a GKE cluster.
	Cluster struct {
		Name                  string `json:"name,omitempty"`
		Description           string `json:"description,omitempty"`
		CreateTime            string `json:"createTime,omitempty"`
		InitialClusterVersion string `json:"initialClusterVersion,omitempty"`
		CurrentMasterVersion  string `json:"currentMasterVersion,omitempty"`
		CurrentNodeVersion    string `json:"currentNodeVersion,omitempty"`
		InitialNodeCount      int64  `json:"initialNodeCount,omitempty"`
		CurrentNodeCount      int64  `json:"currentNodeCount,omitempty"`
		MaximumPodsPerNode    int64  `json:"maximumPodsPerNode,omitempty"`
		LoggingService        string `json:"loggingService,omitempty"`
		MonitoringService     string `json:"monitoringService,omitempty"`
		DatabaseEncryption    string `json:"databaseEncryption,omitempty"`
		DatabaseEncryptionKey string `json:"databaseEncryptionKey,omitempty"`
	}
)
