#!/bin/bash

set -e

podman run -d \
    -e KEYCLOAK_USER=admin \
    -e KEYCLOAK_PASSWORD=password \
    -p 8180:8080 \
    --name keycloak \
    jboss/keycloak 