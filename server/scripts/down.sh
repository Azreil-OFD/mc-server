docker stop $(docker ps -a -q --filter "name=minecraft_container")
docker rm $(docker ps -a -q --filter "name=minecraft_container")
docker rmi $(docker images --filter=reference="*m-server*" -q)