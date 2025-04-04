docker run -d \
  -p 25565:25565 \
  -v $(pwd)/:/server/ \
  --name minecraft_container \
  m-server:latest