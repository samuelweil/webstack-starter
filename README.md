# Full Stack Starter

This project is designed as a full-stack web application running in Kubernetes. 

## Developing

In order to run this project locally in development, install:

- [Skaffold](https://skaffold.dev/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)

A skaffold configuration is set up to allow for local development.

1. Start minikube using `minikube start`
2. Run the application in dev mode using `skaffold dev`