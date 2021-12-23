# Okteto Pipeline with Terraform (PubSub)

This sample covers a producer/consumer application in Golang using Pub/Sub.

The okteto pipeline uses terraform to create a Pub/Sub topic on deployment, and delete it on pipeline destroy.
The terraform state is stored in a Google Cloud Bucket.

### Configure Okteto Secrets

By default, a fake Google Cloud account is used to access PubSub. If you want to use your own account, define the following secrets:

- [Create a secret](https://okteto.com/docs/cloud/secrets/) `GCP_CREDENTIALS` with your Google Cloud credentials

- [Create a secret](https://okteto.com/docs/cloud/secrets/) `GCP_PROJECT_ID` with your Google Cloud project

- [Create a secret](https://okteto.com/docs/cloud/secrets/) `BUCKET` with the bucket name to persist Terraform states


### Deploy the PubSub sample app

Login to the Okteto UI and deploy the https://github.com/okteto/pipeline-with-terraform git repository.

In less than a minute, the application should be available!

The pipeline is defined in [this file](https://github.com/okteto/pipeline-with-terraform/blob/main/okteto-pipeline.yml).

A summary of the pipeline is:

- Install the right version of terraform using `tenv`
- Generate configuration files rendering the Okteto Secrets
- Apply the Terraform plan to create a PubSub topic
- Build the coniner images for the `pub` and `sub` microservices
- Render the kustomization file
- Deploy the `pub` and `sub` Kubernetes manifests

### Develop the "pub" microservice

- [Install](https://okteto.com/docs/getting-started/installation/) the Okteto CLI

- Configure the Okteto CLI context to your Okteto Cluster: `okteto login https://your-company.okteto.dev`

- Start developing with `okteto up -f pub/okteto.yml`.

> [This tutorial](https://okteto.com/docs/samples/golang/) shows the full golang dev flow including how to configure your debugger

Happy coding!


