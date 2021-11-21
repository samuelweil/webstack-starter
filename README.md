# Full Stack Starter

This project is designed as a full-stack web application running in Kubernetes.

## Developing

This project requires the kubernetes supported
[ingress-nginx](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/) ingress
controller. Make sure to install it in the Kubernetes cluster prior to deploying the project. See
[the docs](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start) for more info.

In order to run this project locally in development, install:

- [Skaffold](https://skaffold.dev/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)

A skaffold configuration is set up to allow for local development.

1. Start minikube using `minikube start`
2. Enable the ingress controller `minikube addons enable ingress`
3. Run the application in dev mode using `skaffold dev`

Routes will be enabled via the portforward at localhost:3000. Go to localhost:3000 or
localhost:3000/api to see.

## Configuration

The following configuration variables must be set in a root level file called env.config using.
`<parameter>=<value>` notation

| Parameter | Description                                                      | Required |
| --------- | ---------------------------------------------------------------- | -------- |
| client_id | Client id of the Google API credential to use for the OAuth flow | Y        |
