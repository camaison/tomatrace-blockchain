docker-compose -f network.yaml down -v
docker rm $(docker ps -aq)
docker rmi $(docker images dev-* -q)
