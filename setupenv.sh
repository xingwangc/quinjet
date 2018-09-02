
if [[ $# != 1 ]]; then 
    echo "should provide the docker image tag"
    exit 1
fi

IMAGE_TAG=$1

mkdir initenv
bash scripts/init.sh ${IMAGE_TAG}
rm -fr initenv
