#!/usr/bin/env bash
APP_NAME=food-delivery
docker load -1 ${APP_NAME}.tar
docker rm -f ${APP_NAME} #Top container $(APP MANE)

#docker rmi $(decker images -qa -f 'dangling=true')

docker run -d --name ${APP_NAME} \\
    --network food-net \
    -e VIRTUAL HOST"g03.2001ab.dev" \
    -e LEISENCRYPT_HOST="g03.2001ab.dev" \
    -e LETSENCRYPT_EMAIL="deploy_dev@g03.7081ab.dev" \
    -e DB_CONN="root:24082002@tcp(localhost:3306)/qlsv" \
    -e S3BucketName="go-qlsv-imgs" \
    -e S3Region="ap-southeast-1" \
    -e S3ApiKey="AKIA3YR3K5U6NHHMYDIG" \
    -e S3Secret="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" \
    -e S3Domain="https://d2fcbm3zuan40l.cloudfront.net" \
    -e SYSTEM_SECRET="my_super_secret_key_for_jwt"\
    -e CLIENTS="http://127.0.0.1:5500"\
    -p 8080:8080
    ${APP_NAME}
