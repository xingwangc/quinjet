#!/usr/local/bin/bash

MSG_QUEUE=$(./quinjetmd5 ${GIT_REPO} ${DEPENDANCE} ${BUILD})

./rabbitmqadmin -H ${RABBITMQ_HOST} -P ${RABBITMQ_PORT} -u ${RABBITMQ_USER} -p ${RABBITMQ_PWD} list vhosts
./rabbitmqadmin declare queue name=${MSG_QUEUE} durable=false

#start ci
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="Start CI"

git clone ${GIT_REPO} work_dir
cd work_dir
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="clone repo from ${GIT_REPO} done"


npm install ${DEPENDANCE}
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="install DEPENDANCE done"

make
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="build project done"

docker build ${IMAGE_TAG} .
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="build docker image done"

docker push ${IMAGE_TAG}
./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="push docker image done"

./rabbitmqadmin publish routing_key=${MSG_QUEUE} payload="complete"

./rabbitmqadmin -f tsv -q list connections name | while read conn ; do rabbitmqadmin -q close connection name="${conn}" ; done
