#!usr/bin/env bash

# build and transfer to Virtual machine
APP_NAME=student-management
DEPLOY_CONNECT=root@<domain or ip adress of virtual machine>
go mod download
echo "Downloading..."
CGO_ENABLED=0 GOOS=Linux go build -a -installsuffix cgo -o app

echo "Docker building..."
docker build -t ${APP_NAME} -f ./Dockerfile .
echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}


echo "Deploying..."
scp -o StrictHostKeyChecking=no ./${APP_NAME}.tar ${DEPLOY_CONNECT}:~
ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/stg.sh
#

echo "Cleaning..."
rm -f ./${APP_NAME}.tar
#docker rmi ${docker images -qa -f 'danling=true'}
echo "Done!"