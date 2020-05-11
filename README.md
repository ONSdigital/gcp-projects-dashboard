# GCP Projects Dashboard
This repository contains a dashboard that displays useful information from multiple [GCP](https://cloud.google.com/) projects, with a particular focus on [GKE](https://cloud.google.com/kubernetes-engine) clusters.

## Organisation
This repository contains the following sub-directories:

* [collector](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/collector) - [Go](https://golang.org/) application that runs as a Kubernetes [CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) and calls the [Kubernetes Engine API](https://cloud.google.com/kubernetes-engine/docs/reference/rest) to collect information about each GKE cluster of interest. A [Cloud Firestore](https://cloud.google.com/firestore/) database is used as persistent storage
* [webapp](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/dashboard) - Dashboard web application that displays the information held in Firestore

## Building
For the agent and API server, use `make` to compile binaries for macOS and Linux.

## Copyright
Copyright (C) 2020 Crown Copyright (Office for National Statistics)