#!bin/bash
 GOOS=linux CGO_ENABLED=0 go build -o docker-file
 sudo docker build -t docker-file .
 sudo docker run docker-file
