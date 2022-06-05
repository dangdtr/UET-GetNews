# Overview
A simple tool that detects new notices on the university website and send it to my personal mail.

## Prepare
```shell
cp .env.example .env 
```

## Normal run
```shell
go run *.go 
```
## Run with docker-file
```shell
sudo sh ./docker-build.sh  
```
## Run with docker-compose
```shell
sudo docker-compose up
```

## Deploy to heroku
```shell
sudo sh ./heroku-deploy-docker.sh  
```
### Stopping one-off dynos (in heroku)
```shell
heroku ps:stop [NAME_OF_APP]
#restart
heroku restart [NAME_OF_APP]
```
This project has no Graceful Shutdown
