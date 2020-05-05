package project

import (
	"log"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	cloudresourcemanager "google.golang.org/api/cloudresourcemanager/v1"
	compute "google.golang.org/api/compute/v1"
)

// CurrentProject returns information about the current project.
func CurrentProject() *cloudresourcemanager.Project {
	ctx := context.Background()

	client, err := google.DefaultClient(ctx, cloudresourcemanager.CloudPlatformScope)
	if err != nil {
		log.Fatal(err)
	}

	service, err := cloudresourcemanager.New(client)
	if err != nil {
		log.Fatal(err)
	}

	projectID := currentProjectID()
	project, err := service.Projects.Get(projectID).Context(ctx).Do()
	if err != nil {
		log.Fatal(err)
	}

	return project
}

func currentProjectID() string {
	ctx := context.Background()
	credentials, err := google.FindDefaultCredentials(ctx, compute.ComputeScope)
	if err != nil {
		log.Fatal(err)
	}

	return credentials.ProjectID
}
