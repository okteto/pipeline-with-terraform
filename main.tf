terraform {
  backend "gcs" {
  }
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = ">= 4.5.0"
    }
  }
  required_version = ">= 1.1.2"
}
provider "google" {
  credentials = base64decode(var.credentials)
  project     = local.credentials.project_id
}

resource "google_pubsub_topic" "this" {
  name = format("topic-%s", var.name)
}

locals {
  credentials = jsondecode(base64decode(var.credentials))
}

variable "credentials" {
  type        = string
  description = "google cloud credentials"
}

variable "name" {
  type        = string
  description = "pubsub topic name"
}

output "topicID" {
  value = google_pubsub_topic.this.id
}
