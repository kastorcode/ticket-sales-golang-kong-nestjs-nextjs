FROM node:21-slim
RUN apt update && apt install -y openssl procps
RUN yarn global add @nestjs/cli@10.3.2
WORKDIR /home/node/app
USER node
CMD tail -f /dev/null