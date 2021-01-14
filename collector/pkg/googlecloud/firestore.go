package googlecloud

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
const masterAlertingCollection = "gcp-projects-dashboard-gke-master-version-alerts"
const nodeAlertingCollection = "gcp-projects-dashboard-gke-node-version-alerts"
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

// GetAvailableGKEVersions returns the lists of available GKE master and node versions from Firestore.
func (c FirestoreClient) GetAvailableGKEVersions() GKEVersions {
	doc := c.client.Collection(versionsCollection).Doc(london)
	snapshot, err := doc.Get(*c.context)
	if err != nil {
		log.Fatalf("Failed to read Firestore document %s in collection %s: %v", london, versionsCollection, err)
	}

	var versions GKEVersions
	if err := snapshot.DataTo(&versions); err != nil {
		log.Fatalf("Failed to populate GKEVersions struct: %v", err)
	}

	return versions
}

// GKEMasterVersionAlert returns any current GKE master version alert set for the passed project.
func (c FirestoreClient) GKEMasterVersionAlert(projectName string) string {
	return c.versionAlert(masterAlertingCollection, projectName)
}

// GKENodeVersionAlert returns any current GKE node version alert set for the passed project.
func (c FirestoreClient) GKENodeVersionAlert(projectName string) string {
	return c.versionAlert(nodeAlertingCollection, projectName)
}

// SaveGKEClusterDetails creates or updates the Firestore document with the passed name, setting its contents to the passed cluster details.
func (c FirestoreClient) SaveGKEClusterDetails(projectName string, clusterDetails map[string]interface{}) {
	doc := c.client.Collection(clustersCollection).Doc(projectName)
	_, err := doc.Set(*c.context, map[string]interface{}{
		"cluster": clusterDetails,
		"updated": now(),
	})
	if err != nil {
		log.Fatalf("Failed to update Firestore document %s in collection %s: %v", projectName, clustersCollection, err)
	}
}

// SaveGKEMasterVersionAlert updates the Firestore document with the passed GKE master version for the passed project.
func (c FirestoreClient) SaveGKEMasterVersionAlert(masterVersion, projectName string) {
	c.saveVersionAlert(masterAlertingCollection, masterVersion, projectName)
}

// SaveGKENodeVersionAlert updates the Firestore document within the passed GKE node version for the passed project.
func (c FirestoreClient) SaveGKENodeVersionAlert(nodeVersion, projectName string) {
	c.saveVersionAlert(nodeAlertingCollection, nodeVersion, projectName)
}

func now() string {
	return time.Now().Format("Monday 02 Jan 2006 15:04:05 MST")
}

func (c FirestoreClient) saveVersionAlert(collection, version, projectName string) {
	doc := c.client.Collection(collection).Doc(projectName)
	_, err := doc.Set(*c.context, map[string]interface{}{
		"version": version,
		"posted":  now(),
	})
	if err != nil {
		log.Fatalf("Failed to update Firestore document %s in collection %s: %v", projectName, collection, err)
	}
}

func (c FirestoreClient) versionAlert(collection, projectName string) string {
	doc := c.client.Collection(collection).Doc(projectName)
	snapshot, err := doc.Get(*c.context)
	if !snapshot.Exists() {
		return ""
	}

	if err != nil {
		if status.Code(err) == codes.NotFound {
			return ""
		} else {
			log.Fatalf("Failed to read Firestore document %s in collection %s: %v", projectName, collection, err)
		}
	}

	version, _ := snapshot.DataAt("version")
	if version == nil {
		return ""
	}

	return version.(string)
}
