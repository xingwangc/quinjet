FROM nginx:latest

RUN apt-get update && apt-get install -y python git

RUN curl -sL https://deb.nodesource.com/setup_0.12 | bash -

RUN apt-get update && apt-get install -y nodejs

COPY . .
