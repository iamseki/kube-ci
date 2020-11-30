# K8s Continuous Deployment

## Cluster Authentication

For now, the strategy to allow the container runned based on this image to access your clusters  is as the follow:

-   Generate your own kube config file named as ***kubeconfig*** in root folder, besides Dockerfile.
-   Build and Push this image to your own registry or another one such as docker hub and use it in ci/cd `deployment pipeline`.

>In the future it'll be possible to download kubeconfig for some storage cloud provider such as AWS or GCP

## Build this image to use in your ci/cd pipeline

-   `docker build --build-arg CONFIG_FILE_URL=bukcets3testing --build-arg CONFIG_PROVIDER=aws -t kube-ci .`

- `docker build -t kube-ci .` ensure that already exists a kubeconfig in root folder

## Commands Available