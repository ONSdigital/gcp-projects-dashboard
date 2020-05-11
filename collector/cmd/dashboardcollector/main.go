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
	projects := ""
	if projects = os.Getenv("GCP_PROJECTS"); len(projects) == 0 {
		log.Fatal("Missing GCP_PROJECTS environment variable")
	}

	projectNames := strings.Split(projects, "\n")
	for _, projectName := range projectNames {
		client := googlecloud.NewGKEClient(projectName)
		cluster := client.GetFirstCluster()
		json, err := redactJSON(cluster, "masterAuth")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", json)
		time.Sleep(rateLimitPause)
	}
}

func redactJSON(obj interface{}, redactedFields ...string) (string, error) {
	redactedJSON, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	if len(redactedFields) == 0 {
		return string(redactedJSON), nil
	}

	jsonMap := map[string]interface{}{}
	json.Unmarshal([]byte(string(redactedJSON)), &jsonMap)

	for _, field := range redactedFields {
		delete(jsonMap, field)
	}

	redactedJSON, err = json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	return string(redactedJSON), nil
}
