#!/bin/bash
echo "$docker_password" | docker login ghcr.io --username "$docker_username" --password-stdin
docker push "ghcr.io/$docker_username/restapi:1.0-${GIT_COMMIT::8}"
docker push "ghcr.io/$docker_username/restapi:latest"
docker push "ghcr.io/$docker_username/migrator:1.0-${GIT_COMMIT::8}"
docker push "ghcr.io/$docker_username/migrator:latest" &
wait