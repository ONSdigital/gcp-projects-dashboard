package main

import (
	"fmt"

	project "github.com/ONSdigital/gcp-projects-dashboard/pkg/googlecloud"
)

func main() {
	fmt.Println(project.CurrentProject())
}
