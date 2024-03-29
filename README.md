# kettle-templates

This repository contains machine learning pipeline and microservice templates that you can use to bootstrap your projects, with the [kettle CLI](https://github.com/nlathia/kettle-cli).

## How to use a template

1. Install the [kettle CLI](https://github.com/nlathia/kettle-cli);
2. Use `kettle create <template-name>` where the template name is in the first column below 👇

## AWS Lambda

| kettle create      | Description |
| ----------- | ----------- |
| pyenv-aws-lambda      | Create a Python AWS Lambda function and set it up locally with [pyenv](https://github.com/pyenv/pyenv) and [pyenv-virtualenv](https://github.com/pyenv/pyenv-virtualenv)      |
| conda-aws-lambda      | Create a Python AWS Lambda function and set it up locally with [conda](https://docs.conda.io/en/latest/)      |
| golang-aws-lambda   | A Go template for a basic AWS Lambda function        |
| golang-aws-lambda-dynamo-s3   | A Go template for a AWS Lambda function that includes initialising code for Dynamo DB and s3.        |


## Google Cloud Functions

| kettle create      | Description |
| ----------- | ----------- |
| pyenv-gcloud-function      | Create a Python Cloud Function function and set it up locally with [pyenv](https://github.com/pyenv/pyenv) and [pyenv-virtualenv](https://github.com/pyenv/pyenv-virtualenv)       |
| golang-gcloud-function   | A Go template for a Google Cloud function        |

## Google Cloud Run

| kettle create      | Description |
| ----------- | ----------- |
| pyenv-gcloud-run      | Create a Python Cloud Run container and set it up locally with [pyenv](https://github.com/pyenv/pyenv) and [pyenv-virtualenv](https://github.com/pyenv/pyenv-virtualenv)       |
| golang-gcloud-run   | A Go template for a Google Cloud Run container        |
