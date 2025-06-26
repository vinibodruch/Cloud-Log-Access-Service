# Cloud-Log-Access-Service

![GitHub release](https://img.shields.io/github/v/release/vinibodruch/Cloud-Log-Access-Service.svg) ![Keycloak Version](https://img.shields.io/badge/Keycloak-v26.2-blue) ![Postgres Version](https://img.shields.io/badge/PostgresDB-v17-cyan) ![Docker Version](https://img.shields.io/badge/Docker-v28-darkblue) ![Docker Compose Version](https://img.shields.io/badge/Docker_Compose-v2.36-darkblue)
![Terraform version](https://img.shields.io/badge/Docker_Compose-v2.36-darkblue)

A BFF service to access log files across multiple Cloud Providers

## Requirements

- An Azure Az subscription logged in your terminal, with Az CLI and an active subscription. See [Manage Azure Subscriptions](https://learn.microsoft.com/en-us/cli/azure/manage-azure-subscriptions-azure-cli?view=azure-cli-latest&tabs=bash) and [az cli](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest) documentation
- An AWS CLI account logged in your terminal. See [Configuring using AWS CLI commands](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-quickstart.html#getting-started-quickstart-new-command)
- Terraform installed. See [Instalation](https://developer.hashicorp.com/terraform/install)
- Docker and docker compose installed. See [Instalation](https://docs.docker.com/engine/install/)
- Golang installed. Check [this](https://go.dev/doc/install)

## Preparation

### Terraform environment

- `cd terraform && terraform init && terraform plan`
- `terraform apply # If it all ok`

### Docker local environment

- `docker compose up`

### Authenticate on keycloak

    curl -X POST \
    "http://localhost:8082/realms/SAP/protocol/openid-connect/token" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=password" \
    -d "client_id=cloud-log-access-service" \
    -d "client_secret=secret-key-for-cloud-log-access-service" \
    -d "username=john.doe" \
    -d "password=john.doe" \
    -d "scope=openid profile email"
    curl: (56) Recv failure: Connection r

### Uninstall

`docker compose down -v`
