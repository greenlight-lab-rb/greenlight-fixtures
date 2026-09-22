# Deliberately outdated base image so the docker ecosystem has something to do.
FROM node:27.0.0-alpine
WORKDIR /app
COPY package.json .
CMD ["node", "--version"]
