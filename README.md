# GCP Projects Dashboard
This repository contains a dashboard that displays useful information from multiple [GCP](https://cloud.google.com/) projects, with a particular focus on [GKE](https://cloud.google.com/kubernetes-engine) clusters.

## Organisation
This repository contains the following sub-directories:

* [collector](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/collector) - [Go](https://golang.org/) application that runs as a Kubernetes [CronJob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/) and calls the [Kubernetes Engine API](https://cloud.google.com/kubernetes-engine/docs/reference/rest) to collect information about each GKE cluster of interest. A [Cloud Firestore](https://cloud.google.com/firestore/) database is used as persistent storage. **Note that it is assumed there is only one GKE cluster per GCP project**

* [versions](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/versions) - Go application that runs as a Kubernetes CronJob and calls the Kubernetes Engine API to retrieve inforation about available GKE versions. A Cloud Firestore database is used as persistent storage.

* [webapp](https://github.com/ONSdigital/gcp-projects-dashboard/tree/master/webapp) - [Ruby](https://ruby-lang.org/) [Sinatra](http://sinatrarb.com/) dashboard application that displays the information held in Firestore

## Building
For the collector and versions applications, use `make` to compile binaries for macOS and Linux.

## Environment Variables
The environment variables below are required:

| Component | Variable               | Purpose                                                                                                |
|-----------|------------------------|--------------------------------------------------------------------------------------------------------|
| collector | `FIRESTORE_PROJECT`    | Name of the GCP project containing the Firestore database.                                             |
|           | `GCP_PROJECTS`         | List of GCP projects containing the GKE clusters to collect information for (one cluster per project). |
|           | `SLACK_CHANNEL`        | Name of the Slack channel for post expiring GKE master/node version alerts to.                         |
|           | `SLACK_WEBHOOK`        | Slack webhook for posting expiring GKE master/node version alerts to.                                  |
| versions  | `FIRESTORE_PROJECT`    | Name of the GCP project containing the Firestore database.                                             |
|           | `GCP_PROJECT`          | Name of the GCP project to use when invoking the Kubernetes Engine API.                                |
| webapp    | `FIRESTORE_PROJECT`    | Name of the GCP project containing the Firestore database.                                             |
|           | `GCP_CONSOLE_BASE_URL` | Base URL to use for the project hyperlinks. The project name is appended to this URL.                  |
|           | `GCP_ORGANISATION`     | Name of the GCP organisation the deployed dashboard is reporting against. Displayed in the heading.    |

## IAM Roles
The following [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) are required to run this software:

| Component | IAM Role                        |
|-----------|---------------------------------|
| collector | `roles/browser`                 |
|           | `roles/container.clusterViewer` |
|           | `roles/datastore.user`          |
| versions  | `roles/datastore.user`          |
|           | `roles/container.clusterViewer` |
| webapp    | `roles/datastore.user`          |

## Copyright
Copyright (C) 2020-2021 Crown Copyright (Office for National Statistics)