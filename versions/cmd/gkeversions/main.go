package main

import (
	"log"
	"os"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/versions/pkg/googlecloud"
)

func main() {
	firestoreProject := ""
	if firestoreProject = os.Getenv("FIRESTORE_PROJECT"); len(firestoreProject) == 0 {
		log.Fatal("Missing FIRESTORE_PROJECT environment variable")
	}

	project := ""
	if project = os.Getenv("GCP_PROJECT"); len(project) == 0) {
		log.Fatal("Missing GCP_PROJECT environment variable")
	}

	firestoreClient := googlecloud.NewFirestoreClient(firestoreProject)
	client := googlecloud.NewGKEClient(project)
	serverConfig := client.ListVersions()

	if serverConfig != nil {
		err := firestoreClient.SaveDoc(serverConfig)
		if err != nil {
			log.Fatalf("Failed to save document to Firestore: %v", err)
		}
	}
}
