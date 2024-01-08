#!/bin/bash
[[ -z "${GIT_COMMIT}" ]] && Tag='local' || Tag="${GIT_COMMIT::8}"
[[ -z "${docker_username}" ]] && DockerRepo='' || DockerRepo="ghcr.io/${docker_username}/"
docker build -t "${DockerRepo}migrator:latest" -t "${DockerRepo}migrator:1.0-$Tag" migrator/
docker build -t "${DockerRepo}restapi:latest" -t "${DockerRepo}restapi:1.0-$Tag" restapi/
