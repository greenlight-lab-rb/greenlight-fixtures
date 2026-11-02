# Deliberately outdated base image so the docker ecosystem has something to do.
FROM node:26.13.0-alpine
WORKDIR /app
COPY package.json .
CMD ["node", "--version"]
