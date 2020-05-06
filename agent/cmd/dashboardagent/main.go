package main

import (
	"encoding/json"
	"fmt"
	"log"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/agent/pkg/googlecloud"
	response "github.com/ONSdigital/gcp-projects-dashboard/agent/pkg/response"
)

func main() {
	project := googlecloud.CurrentProject()
	client := googlecloud.NewGKEClient(project.Name)
	clusters := client.ListClusters()
	cluster := client.GetCluster(clusters.Clusters[0].Name)

	projectJSON := response.Project{
		ID:         project.ProjectId,
		Labels:     project.Labels,
		Name:       project.Name,
		Number:     project.ProjectNumber,
		CreateTime: project.CreateTime}

	clusterJSON := response.Cluster{
		Name:                   cluster.Name,
		Description:            cluster.Description,
		CreateTime:             cluster.CreateTime,
		InitialClusterVersion:  cluster.InitialClusterVersion,
		CurrentMasterVersion:   cluster.CurrentMasterVersion,
		CurrentNodeVersion:     cluster.CurrentNodeVersion,
		InitialNodeCount:       cluster.InitialNodeCount,
		CurrentNodeCount:       cluster.CurrentNodeCount,
		MaximumPodsPerNode:     cluster.DefaultMaxPodsConstraint.MaxPodsPerNode,
		ShieldedNodes:          cluster.ShieldedNodes.Enabled,
		LoggingService:         cluster.LoggingService,
		MonitoringService:      cluster.MonitoringService,
		DatabaseEncryption:     cluster.DatabaseEncryption.State,
		DatabaseEncryptionKey:  cluster.DatabaseEncryption.KeyName,
		Endpoint:               cluster.Endpoint,
		Location:               cluster.Location,
		Network:                cluster.Network,
		NodeIPv4CIDRSize:       cluster.NodeIpv4CidrSize,
		ServicesIPv4CIDR:       cluster.ServicesIpv4Cidr,
		Status:                 cluster.Status,
		StatusMessage:          cluster.StatusMessage,
		Subnetwork:             cluster.Subnetwork,
		WorkloadIdentityConfig: cluster.WorkloadIdentityConfig.WorkloadPool}

	responseJSON := response.Response{Project: projectJSON, Cluster: clusterJSON}

	response, err := json.MarshalIndent(responseJSON, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", response)
}
