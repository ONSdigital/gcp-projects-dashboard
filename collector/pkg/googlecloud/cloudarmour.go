package googlecloud

import (
	"log"

	"golang.org/x/net/context"
	compute "google.golang.org/api/compute/v1"
)

type (

	// ComputeClient wraps a Google Compute Engine API client.
	ComputeClient struct {
		projectName             string
		context                 *context.Context
		computeService          *compute.Service
		securityPoliciesService *compute.SecurityPoliciesService
	}
)

// NewComputeClient instantiates a Google Compute Engine API client.
func NewComputeClient(projectName string) *ComputeClient {
	ctx := context.Background()
	computeService, err := compute.NewService(ctx)
	if err != nil {
		log.Fatalf("Failed to instantiate compute service: %v", err)
	}

	securityPoliciesService := compute.NewSecurityPoliciesService(computeService)

	return &ComputeClient{
		projectName:             projectName,
		context:                 &ctx,
		computeService:          computeService,
		securityPoliciesService: securityPoliciesService,
	}
}

// ListSecurityPolicies returns a list of Cloud Armour security policies within the project set on the client.
func (c ComputeClient) ListSecurityPolicies() *compute.SecurityPolicyList {
	// name := fmt.Sprintf("projects/%s/global/securityPolicies", c.projectName)
	securityPolicies, err := c.securityPoliciesService.List(c.projectName).Context(*c.context).Do()
	if err != nil {
		log.Fatalf("Failed to list security policies in %s: %v", c.projectName, err)
	}

	return securityPolicies
}
