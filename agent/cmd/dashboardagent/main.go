package main

import (
	googlecloud "github.com/ONSdigital/gcp-projects-dashboard/agent/pkg/googlecloud"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Dump(googlecloud.CurrentProject())
}
