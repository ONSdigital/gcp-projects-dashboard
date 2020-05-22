# GCP Projects Dashboard
This repository contains a dashboard that displays useful information from multiple [GCP](https://cloud.google.com/) projects, with a particular focus on [GKE](https://cloud.google.com/kubernetes-engine) clusters.

## Organisation
This repository contains the following sub-directories:

* [collector](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/collector) - [Go](https://golang.org/) application that runs as a Kubernetes [CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) and calls the [Kubernetes Engine API](https://cloud.google.com/kubernetes-engine/docs/reference/rest) to collect information about each GKE cluster of interest. A [Cloud Firestore](https://cloud.google.com/firestore/) database is used as persistent storage. **Note that it is assumed there is only one GKE cluster per GCP project**

* [webapp](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/webapp) - [Ruby](https://ruby-lang.org/) [Sinatra](http://sinatrarb.com/) dashboard application that displays the information held in Firestore

## Building
For the collector, use `make` to compile binaries for macOS and Linux.

## Environment Variables
foo

| Component | Variable               | Purpose                                                                                                |
|-----------|------------------------|--------------------------------------------------------------------------------------------------------|
| collector | `FIRESTORE_PROJECT`    | Name of the GCP project containing the Firestore project.                                              |
|           | `GCP_PROJECTS`         | List of GCP projects containing the GKE clusters to collect information for (one cluster per project). |
| webapp    | `FIRESTORE_PROJECT`    | Name of the GCP project containing the Firestore project.                                              |
|           | `GCP_CONSOLE_BASE_URL` | Base URL to use for the project hyperlinks. The project name is appended to this URL.                  |
|           | `GCP_ORGANISATION`     | Name of the GCP organisation the deployed dashboard is reporting against. Displayed in the heading.    |

## IAM Roles
The following [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) are required to run this software:

| Component | IAM Role                        |
|-----------|---------------------------------|
| collector | `roles/browser`                 |
|           | `roles/container.clusterViewer` |
|           | `roles/datastore.user`          |
| webapp    | `roles/datastore.user`          |

## Copyright
Copyright (C) 2020 Crown Copyright (Office for National Statistics)