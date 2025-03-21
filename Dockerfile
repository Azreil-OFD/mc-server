FROM openjdk:8-jre

RUN apt-get update && apt-get install -y curl \
    && curl -fsSL https://deb.nodesource.com/setup_16.x | bash - \
    && apt-get install -y nodejs \
    && npm install -g pm2 \
    && curl -sSLf https://get.tuna.am | sh

RUN tuna config save-token tt_ubx4yk96ygl7n8lt0uh2npkk703aow1x

WORKDIR /server

COPY . /server

EXPOSE 25565

CMD ["pm2-runtime", "start", "pm2-minecraft.json"]