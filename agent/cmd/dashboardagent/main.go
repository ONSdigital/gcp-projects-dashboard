package main

import (
	"fmt"

	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/pkg/googlecloud"
)

func main() {
	fmt.Println(googlecloud.CurrentProject())
}
