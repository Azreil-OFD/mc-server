docker stop minecraft_container
docker rm minecraft_container
docker run -d \
  -p 25565:25565 \
  -p 8281:8281
  -v $(pwd)/:/server/ \
  --name minecraft_container \
  m-server:latest