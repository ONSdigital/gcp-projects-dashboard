package googlecloud

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
	container "google.golang.org/api/container/v1"
)

type (

	// GKEClient wraps a Google Kubernetes Engine API client.
	GKEClient struct {
		projectName                      string
		context                          *context.Context
		containerService                 *container.Service
		projectsLocationsClustersService *container.ProjectsLocationsClustersService
	}
)

const london = "europe-west2"

// NewGKEClient instantiates a Google Kubernetes Engine API client.
func NewGKEClient(projectName string) *GKEClient {
	ctx := context.Background()
	containerService, err := container.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to instantiate container service: %v", err)
	}

	projectLocationsClustersService := container.NewProjectsLocationsClustersService(containerService)

	return &GKEClient{
		projectName:                      projectName,
		context:                          &ctx,
		containerService:                 containerService,
		projectsLocationsClustersService: projectLocationsClustersService,
	}
}

// GetCluster returns the details of the specified cluster within the project set on the client.
func (c GKEClient) GetCluster(clusterName string) *container.Cluster {
	name := fmt.Sprintf("projects/%s/locations/%s/clusters/%s", c.projectName, london, clusterName)
	cluster, err := c.projectsLocationsClustersService.Get(name).Context(*c.context).Do()
	if err != nil {
		log.Fatalf("Failed to get cluster %s: %v", name, err)
	}

	return cluster
}

// ListClusters returns a list of GKE clusters within the project set on the client.
func (c GKEClient) ListClusters() *container.ListClustersResponse {
	name := fmt.Sprintf("projects/%s/locations/%s", c.projectName, london)
	clusters, err := c.projectsLocationsClustersService.List(name).Context(*c.context).Do()
	if err != nil {
		log.Fatalf("Failed to list clusters in %s: %v", name, err)
	}

	return clusters
}
