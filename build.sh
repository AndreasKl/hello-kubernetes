#!/bin/bash
set -euo pipefail
IFS=$'\n\t'

gcloud config  set project powerpuffgirls

docker build -t eu.gcr.io/powerpuffgirls/hello-kubernetes:latest .
docker push eu.gcr.io/powerpuffgirls/hello-kubernetes:latest
