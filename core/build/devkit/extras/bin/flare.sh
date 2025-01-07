#!/usr/bin/env sh

# Run commands in the openwrt container
docker compose run -it --rm app sh -c "./bin/flare $1"
