package main

import (
	"fmt"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/agent/pkg/googlecloud"
)

func main() {
	fmt.Println(googlecloud.CurrentProject())
}
