package main

import (
	"encoding/json"
	"fmt"
	"log"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/collector/pkg/googlecloud"
)

func main() {
	client := googlecloud.NewGKEClient("census-rh-jt-gke")
	cluster := client.GetFirstCluster()
	json, err := redactJSON(cluster, "masterAuth")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", json)
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
