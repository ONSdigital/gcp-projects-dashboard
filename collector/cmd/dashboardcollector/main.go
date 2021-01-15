package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/collector/pkg/googlecloud"
	"github.com/ONSdigital/gcp-projects-dashboard/collector/pkg/slack"
)

const rateLimitPause = 5 * time.Second
const supportedGKEVersionWindow = 2

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
			clusterDetails := redactSensitiveFields(cluster, "masterAuth")
			firestoreClient.SaveGKEClusterDetails(projectName, clusterDetails)

			for nodePoolIndex, nodePool := range cluster.NodePools {
				autoUpgrade := nodePool.Management.AutoUpgrade

				if !autoUpgrade {

					// Note that the returned arrays of available GKE master and node versions are ordered latest to earliest.
					versions := firestoreClient.GetAvailableGKEVersions()

					// Post a Slack alert if the current master version is the penultimate or last version in the list of available GKE master versions
					// and if an alert hasn't already been posted for this version. Then save the details to Firestore.
					masterVersion := cluster.CurrentMasterVersion
					if indexOf(versions.MasterVersions, masterVersion) >= len(versions.MasterVersions)-supportedGKEVersionWindow {
						if firestoreClient.GKEMasterVersionAlert(projectName) != masterVersion {
							postMasterSlackMessage(masterVersion, projectName, slackAlertsChannel, slackWebHookURL)
							firestoreClient.SaveGKEMasterVersionAlert(masterVersion, projectName)
						}
					}

					// Do the same for the node version.
					nodeVersion := cluster.CurrentNodeVersion
					if indexOf(versions.NodeVersions, nodeVersion) >= len(versions.NodeVersions)-supportedGKEVersionWindow {
						if firestoreClient.GKENodeVersionAlert(projectName) != nodeVersion {
							postNodeSlackMessage(nodePoolIndex, nodeVersion, projectName, slackAlertsChannel, slackWebHookURL)
							firestoreClient.SaveGKENodeVersionAlert(nodeVersion, projectName)
						}
					}
				}

				time.Sleep(rateLimitPause)
			}
		}
	}
}

func indexOf(versions []string, version string) int {
	for i, v := range versions {
		if v == version {
			return i
		}
	}
	return -1
}

func postMasterSlackMessage(version, projectName, slackAlertsChannel, slackWebHookURL string) {
	text := fmt.Sprintf("GKE master version *%s* in cluster *%s* will go out of support soon. Automatic upgrades are disabled for this cluster. Please upgrade.", version, projectName)
	postSlackMessage(text, slackAlertsChannel, slackWebHookURL)
}

func postNodeSlackMessage(nodePoolIndex int, version, projectName, slackAlertsChannel, slackWebHookURL string) {
	text := fmt.Sprintf("GKE node version *%s* for *node pool %d* in cluster *%s* will go out of support soon. Automatic upgrades are disabled for this cluster. Please upgrade.", version, nodePoolIndex, projectName)
	postSlackMessage(text, slackAlertsChannel, slackWebHookURL)
}

func postSlackMessage(text, slackAlertsChannel, slackWebHookURL string) {
	payload := slack.Payload{
		Text:      fmt.Sprintf(text),
		Username:  "GCP Projects Dashboard",
		Channel:   slackAlertsChannel,
		IconEmoji: ":gke:",
	}

	time.Sleep(rateLimitPause)

	err := slack.Send(slackWebHookURL, payload)
	if err != nil {
		log.Fatalf("Failed to send Slack message: %v", err)
	}
}

func redactSensitiveFields(obj interface{}, redactedFields ...string) map[string]interface{} {
	jsonString, err := json.Marshal(obj)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	jsonMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(jsonString)), &jsonMap)

	for _, field := range redactedFields {
		delete(jsonMap, field)
	}

	return jsonMap
}
