# THIS FILE INSTRUCTS EVERYONE TO BUILD A EXECUTE FILE FROM GOLANG FILE AND BUILD WITH DOCKER

## Step 1: To build execute file
```
$env:CGO_ENABLED=0
$env:GOOS="linux"
go build -a -installsuffix cgo -o app
```

## Step 2: Run bitnami/mysql image in your computer
```
docker pull bitnami/mysql
docker run -d --name my-mysql-container -e MYSQL_ROOT_PASSWORD=mysecretpassword bitnami/mysql
docker exec -it my-mysql-container mysql -uroot -p
``` 

Stop and remove containers (if you want):
```
docker stop my-mysql-container
docker rm my-mysql-container
```
You can use `docker rm -f CONTAINER` to remove the container 

## Step 3: Create a network
```
docker network create sm-network
```
Check network
```
docker network ls
```

## Step 4: To build image for your project
```
docker build -t student-management .
docker images 
docker run -v D:\CODE\Golang\list_tutorial_project\project3_REST\.env:/app/.env -d --name student-management -p 3500:8080 --network=sm-network student-management
docker logs {...name...}
```

I use -e to set the environment for my project

## Step 5: Connect MySQL container and your project container to your network
```
docker network connect sm-network my-mysql-container 
```

* `sm-network`: is my network
* `my-mysql-container`: is mysql container
* `student-management`: is student management container

## Step 6: Pull your image to docker hub

```
docker login
docker push your-username/your-repo:tag
```