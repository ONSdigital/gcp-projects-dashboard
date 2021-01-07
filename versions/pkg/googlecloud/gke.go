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
		projectName      string
		context          *context.Context
		containerService *container.Service
	}
)

// NewGKEClient instantiates a Google Kubernetes Engine API client.
func NewGKEClient(projectName string) *GKEClient {
	ctx := context.Background()
	containerService, err := container.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to instantiate container service: %v", err)
	}

	return &GKEClient{
		projectName:      projectName,
		context:          &ctx,
		containerService: containerService,
	}
}

// ListVersions returns a list of GKE versions valid for the project set on the client.
func (c GKEClient) ListVersions() *container.ServerConfig {
	name := fmt.Sprintf("projects/%s/locations/%s", c.projectName, london)
	config, err := c.containerService.Projects.Locations.GetServerConfig(name).Context(*c.context).Do()
	if err != nil {
		log.Fatalf("Failed to get server config for %s: %v", name, err)
	}

	return config
}
