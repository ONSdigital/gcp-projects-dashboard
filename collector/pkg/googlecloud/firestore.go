package googlecloud

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
)

type (

	// FirestoreClient wraps a Google Firestore client.
	FirestoreClient struct {
		projectName string
		context     *context.Context
		client      *firestore.Client
	}

	// GKEVersions represents available GKE version information.
	GKEVersions struct {
		MasterVersions []string `firestore:"validMasterVersions"`
		NodeVersions   []string `firestore:"validNodeVersions"`
	}
)

const clustersCollection = "gcp-projects-dashboard"
const versionsCollection = "gcp-projects-dashboard-gke-versions"

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

// MasterVersion returns the GKE master version for the passed GCP project name.
func (c FirestoreClient) MasterVersion(projectName string) string {
	doc := c.client.Collection(clustersCollection).Doc(projectName)
	snapshot, err := doc.Get(*c.context)
	if err != nil {
		log.Fatalf("Failed to read Firestore document %s in collection %s: %v", projectName, clustersCollection, err)
	}
	masterVersion, err := snapshot.DataAt("cluster.currentMasterVersion")
	if err != nil {
		log.Fatalf("Failed to read value of cluster.currentMasterVersion field in Firestore document %s in collection %s: %v", projectName, clustersCollection, err)
	}

	return masterVersion.(string)
}
// NodeVersion returns the GKE node version for the passed GCP project name.
func (c FirestoreClient) NodeVersion(projectName string) string {
	doc := c.client.Collection(clustersCollection).Doc(projectName)
	snapshot, err := doc.Get(*c.context)
	if err != nil {
		log.Fatalf("Failed to read Firestore document %s in collection %s: %v", projectName, clustersCollection, err)
	}
	nodeVersion, err := snapshot.DataAt("cluster.currentNodeVersion")
	if err != nil {
		log.Fatalf("Failed to read value of cluster.currentNodeVersion field in Firestore document %s in collection %s: %v", projectName, clustersCollection, err)
	}

	return nodeVersion.(string)
}
// SaveDoc creates or updates the Firestore document with the passed name, setting its contents to the passed cluster details.
func (c FirestoreClient) SaveDoc(projectName string, clusterDetails map[string]interface{}) error {
	doc := c.client.Collection(clustersCollection).Doc(projectName)
	_, err := doc.Set(*c.context, map[string]interface{}{
		"cluster": clusterDetails,
		"updated": now(),
	})

	return err
}

func now() string {
	return time.Now().Format("Monday 02 Jan 2006 15:04:05 MST")
}
