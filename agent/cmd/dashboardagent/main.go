package main

import (
	"fmt"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/agent/pkg/googlecloud"
)

func main() {
	project := googlecloud.CurrentProject()
	client := googlecloud.NewGKEClient(project.Name)
	clusters := client.ListClusters()
	cluster := client.GetCluster(clusters.Clusters[0].Name)

	fmt.Printf("Cluster %s within project %s has master version %s", cluster.Name, project.Name, cluster.CurrentMasterVersion)
}
