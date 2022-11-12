#!/bin/bash

set -e

podman run -d --rm \
    -p 5432:5432 \
    -e POSTGRES_PASSWORD=password \
    --name postgres \
    postgres