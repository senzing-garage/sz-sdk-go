# g2-sdk-go

## :warning: WARNING: g2-sdk-go is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

The Senzing g2-sdk-go packages provide a Software Development Kit that wraps the
Senzing C SDK APIs.

[![GoReportCard example](https://goreportcard.com/badge/github.com/senzing/g2-sdk-go)](https://goreportcard.com/report/github.com/senzing/g2-sdk-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/g2-sdk-go.svg)](https://pkg.go.dev/github.com/senzing/g2-sdk-go)

## Overview

The Senzing g2-sdk-go packages enable Go programs to call Senzing library functions.
Under the covers, Golang's CGO is used by the g2-sdk-go packages to make the calls
to the Senzing functions.

## Developing with g2-sdk-go

### Install Senzing library

Since the Senzing library is a pre-requisite, it must be installed first.
This can be done by installing the Senzing package using `apt`, `yum`,
or a technique using Docker containers.
Once complete, the Senzing library will be installed in the `/opt/senzing` directory.

- Using `apt`:

    ```console
    wget https://senzing-production-apt.s3.amazonaws.com/senzingrepo_1.0.0-1_amd64.deb
    sudo apt install ./senzingrepo_1.0.0-1_amd64.deb
    sudo apt update
    sudo apt install senzingapi

    ```

- Using `yum`:

    ```console
    sudo yum install https://senzing-production-yum.s3.amazonaws.com/senzingrepo-1.0.0-1.x86_64.rpm
    sudo yum install senzingapi

    ```

- Using Docker:

    1. Build Senzing installer.

        ```console
        curl -X GET \
            --output /tmp/senzing-versions-latest.sh \
            https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/senzing-versions-latest.sh
        source /tmp/senzing-versions-latest.sh

        sudo docker build \
            --build-arg SENZING_ACCEPT_EULA=I_ACCEPT_THE_SENZING_EULA \
            --build-arg SENZING_APT_INSTALL_PACKAGE=senzingapi=${SENZING_VERSION_SENZINGAPI_BUILD} \
            --build-arg SENZING_DATA_VERSION=${SENZING_VERSION_SENZINGDATA} \
            --no-cache \
            --tag senzing/installer:${SENZING_VERSION_SENZINGAPI} \
            https://github.com/senzing/docker-installer.git#main

        ```

    1. Install Senzing.

        ```console
            curl -X GET \
                --output /tmp/senzing-versions-latest.sh \
                https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/senzing-versions-latest.sh
            source /tmp/senzing-versions-latest.sh

            sudo rm -rf /opt/senzing
            sudo mkdir -p /opt/senzing

            sudo docker run \
                --rm \
                --user 0 \
                --volume /opt/senzing:/opt/senzing \
                senzing/installer:${SENZING_VERSION_SENZINGAPI}

        ```

### Create a stack used in testing

The following instructions show how to bring up a test stack to be used
in testing the `g2-sdk-go` packages.

1. Bring up Senzing stack:

    ```console
    export DOCKER_COMPOSE_DIR=~/my-senzing-stack
    export SENZING_DOCKER_COMPOSE_YAML=postgresql/docker-compose-rabbitmq-postgresql-minimal.yaml

    rm -rf ${DOCKER_COMPOSE_DIR:-/tmp/nowhere/for/safety}
    mkdir -p ${DOCKER_COMPOSE_DIR}

    curl -X GET \
        --output ${DOCKER_COMPOSE_DIR}/docker-compose.yaml \
        "https://raw.githubusercontent.com/Senzing/docker-compose-demo/main/resources/${SENZING_DOCKER_COMPOSE_YAML}"

    curl -X GET \
        --output ${DOCKER_COMPOSE_DIR}/docker-versions-latest.sh \
        https://raw.githubusercontent.com/Senzing/knowledge-base/main/lists/docker-versions-latest.sh
    source ${DOCKER_COMPOSE_DIR}/docker-versions-latest.sh

    export SENZING_DATA_VERSION_DIR=/opt/senzing/data
    export SENZING_ETC_DIR=/etc/opt/senzing
    export SENZING_G2_DIR=/opt/senzing/g2
    export SENZING_VAR_DIR=/var/opt/senzing

    export PGADMIN_DIR=${DOCKER_COMPOSE_DIR}/pgadmin
    export POSTGRES_DIR=${DOCKER_COMPOSE_DIR}/postgres
    export RABBITMQ_DIR=${DOCKER_COMPOSE_DIR}/rabbitmq

    sudo mkdir -p ${PGADMIN_DIR} ${POSTGRES_DIR} ${RABBITMQ_DIR}
    sudo chown $(id -u):$(id -g) -R ${DOCKER_COMPOSE_DIR}
    sudo chmod -R 770 ${DOCKER_COMPOSE_DIR}
    sudo chmod -R 777 ${PGADMIN_DIR}

    cd ${DOCKER_COMPOSE_DIR}
    sudo --preserve-env docker-compose up

    ```

### Run test cases

1. Identify git repository.

    ```console
    export GIT_ACCOUNT=senzing
    export GIT_REPOSITORY=g2-sdk-go
    export GIT_ACCOUNT_DIR=~/${GIT_ACCOUNT}.git
    export GIT_REPOSITORY_DIR="${GIT_ACCOUNT_DIR}/${GIT_REPOSITORY}"

    ```

1. Using the environment variables values just set, follow steps in
   [clone-repository](https://github.com/Senzing/knowledge-base/blob/main/HOWTO/clone-repository.md) to install the Git repository.

1. Set environment variables.
   Identify Database URL of database in docker-compose stack.
   Example:

    ```console
    export LOCAL_IP_ADDRESS=$(curl --silent https://raw.githubusercontent.com/Senzing/knowledge-base/main/gists/find-local-ip-address/find-local-ip-address.py | python3 -)
    export SENZING_TOOLS_DATABASE_URL=postgresql://postgres:postgres@${LOCAL_IP_ADDRESS}:5432/G2

    ```

1. Run tests.

    ```console
    cd ${GIT_REPOSITORY_DIR}
    make test

    ```

## Error prefixes

Error identifiers are in the format `senzing-PPPPnnnn` where:

`P` is a prefix used to identify the package.
`n` is a location within the package.

Prefixes:

1. `6001` - g2config
1. `6002` - g2configmgr
1. `6003` - g2diagnostic
1. `6004` - g2engine
1. `6005` - g2hasher
1. `6006` - g2product
1. `6007` - g2ssadm
