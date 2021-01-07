package googlecloud

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	container "google.golang.org/api/container/v1"
)

type (

	// FirestoreClient wraps a Google Firestore client.
	FirestoreClient struct {
		projectName string
		context     *context.Context
		client      *firestore.Client
	}
)

const firestoreCollection = "gcp-projects-dashboard-gke-versions"
const london = "europe-west2"

// NewFirestoreClient instantiates a new Firestore client for the passed GCP project name.
func NewFirestoreClient(projectName string) *FirestoreClient {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectName)
	if err != nil {
		log.Fatalf("Failed to instantiate Firestore client in project %s: %v", projectName, err)
	}

	return &FirestoreClient{
		projectName: projectName,
		context:     &ctx,
		client:      client,
	}
}

// SaveDoc creates or updates the Firestore document with the passed name, setting its contents to the passed cluster details.
func (c FirestoreClient) SaveDoc(serverConfig *container.ServerConfig) error {
	doc := c.client.Collection(firestoreCollection).Doc(london)
	_, err := doc.Set(*c.context, map[string]interface{}{
		"versions": serverConfig,
		"updated":  now(),
	})

	return err
}

func now() string {
	return time.Now().Format("Monday 02 Jan 2006 15:04:05 MST")
}
