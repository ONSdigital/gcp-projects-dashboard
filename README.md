# GCP Projects Dashboard
This repository contains a dashboard that displays useful information from multiple [GCP](https://cloud.google.com/) projects, with a particular focus on [GKE](https://cloud.google.com/kubernetes-engine) clusters.

## Organisation
This repository contains the following sub-directories:

* [agent](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/agent) - Lightweight [Go](https://golang.org/) agent application that runs as a Kubernetes [CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) on each GKE cluster that's included in the dashboard
* [apiserver](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/apiserver) - Central Go API server that each agent posts collected information to. A [Cloud Firestore](https://cloud.google.com/firestore/) database is used as persistent storage
* [dashboard](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/dashboard) - Dashboard web application that displays the information held in Firestore

## Building
For the agent and API server, use `make` to compile binaries for macOS and Linux.

## Copyright
Copyright (C) 2020 Crown Copyright (Office for National Statistics)