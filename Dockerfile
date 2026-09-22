# Deliberately outdated base image so the docker ecosystem has something to do.
FROM node:26.8.6-alpine
WORKDIR /app
COPY package.json .
CMD ["node", "--version"]
