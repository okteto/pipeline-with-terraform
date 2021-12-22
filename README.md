# Okteto Pipeline with Terraform (PubSub)

This sample covers a producer/consumer application in Golang using Pub/Sub.

The okteto pipeline uses terraform to create a Pub/Sub topic on deployment, and delete it on pipeline destroy.
The terraform is store in a google bucket.

### Deploy the sample app

- [Create a secret](https://okteto.com/docs/cloud/secrets/) `GOOGLE_CLOUD_SERVICE_ACCOUNT_KEY` with your Google Cloud credentials. If you want to make the secret available for every namespace, create an [admin secret](https://okteto.com/docs/enterprise/administration/dashboard/#secrets-section).

- [Create a secret](https://okteto.com/docs/cloud/secrets/) `GOOGLE_CLOUD_PROJECT` with your Google Cloud project. If you want to make the secret available for every namespace, create an [admin secret](https://okteto.com/docs/enterprise/administration/dashboard/#secrets-section).

- Login to the Okteto UI and deploy the https://github.com/okteto/pipeline-with-terraform

In less than a minute, the application should be available!

### Develop the "pub" microservice

- [Install](https://okteto.com/docs/getting-started/installation/) the Okteto CLI.

- Configure the Okteto CLI context to your Okteto Cluster: `okteto login https://your-company.okteto.dev`

- Start developing with `okteto up -f pub/okteto.yml`.

> [This tutorial](https://okteto.com/docs/samples/golang/) shows the full golang dev flow including how to configure your debugger

Happy coding!


