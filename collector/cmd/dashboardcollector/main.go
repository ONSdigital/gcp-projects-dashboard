package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/collector/pkg/googlecloud"
)

const rateLimitPause = 5 * time.Second

func main() {
	firestoreProject := ""
	if firestoreProject = os.Getenv("FIRESTORE_PROJECT"); len(firestoreProject) == 0 {
		log.Fatal("Missing FIRESTORE_PROJECT environment variable")
	}

	firestoreClient := googlecloud.NewFirestoreClient(firestoreProject)

	projects := ""
	if projects = os.Getenv("GCP_PROJECTS"); len(projects) == 0 {
		log.Fatal("Missing GCP_PROJECTS environment variable")
	}

	slackAlertsChannel := ""
	if slackAlertsChannel = os.Getenv("SLACK_CHANNEL"); len(slackAlertsChannel) == 0 {
		log.Fatal("Missing SLACK_CHANNEL environment variable")
	}

	slackWebHookURL := ""
	if slackWebHookURL = os.Getenv("SLACK_WEBHOOK"); len(slackWebHookURL) == 0 {
		log.Fatal("Missing SLACK_WEBHOOK environment variable")
	}

	projectNames := strings.Split(projects, "\n")
	for _, projectName := range projectNames {
		fmt.Printf("Getting GKE cluster details for %s\n", projectName)

		client := googlecloud.NewGKEClient(projectName)
		cluster := client.GetFirstCluster()

		if cluster != nil {
			clusterDetails, err := redactSensitiveFields(cluster, "masterAuth")
			if err != nil {
				log.Fatal(err)
			}

			err = firestoreClient.SaveDoc(projectName, clusterDetails)
			if err != nil {
				log.Fatalf("Failed to save document to Firestore: %v", err)
			}

			time.Sleep(rateLimitPause)
		}
	}
}

func redactSensitiveFields(obj interface{}, redactedFields ...string) (map[string]interface{}, error) {
	jsonString, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	jsonMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(jsonString)), &jsonMap)

	for _, field := range redactedFields {
		delete(jsonMap, field)
	}

	return jsonMap, nil
}
