#Overview
A simple tool that detects new notices on the university website and send it to my personal mail.![image](https://user-images.githubusercontent.com/78289033/172062020-8610304f-b51e-42a3-bf18-4b5e802f4c4e.png)


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
