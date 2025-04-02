docker run -d \
  -p 25565:25565 \
  -v $(pwd)/world:/server/world \
  --name minecraft_container \
  m-server:latest