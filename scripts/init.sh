#!/usr/local/bin/bash

IMAGE_TAG=$1

function buildMd5tool() {
    cd tools/quinjetmd5/       
    GOOS=linux GOARCH=amd64 go build -o ../../initenv/quinjetmd5 .
    cd -
}

function copyRabbitadmin() {
    cp tools/rabbitmqadmin ./initenv
}

function copyJobScript() {
    cp scripts/job.sh ./initenv
    chmod u+x ./initenv/job.sh
}

function copyDockerfile() {
    cp dockerfile/quinjetjob.dockerfile ./initenv/Dockerfile
}

function buildJobImage() {
    cd initenv
    docker build -t ${IMAGE_TAG} .
    cd -
}

function pushImage() {
    docker push ${IMAGE_TAG}
}

function main() {
    buildMd5tool
    copyRabbitadmin
    copyJobScript
    copyDockerfile
    buildJobImage
    pushImage
}

main
