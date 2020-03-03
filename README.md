# Kafka Broker Environment


## Install
`go get github.com/abhibvp003/docker-kafka`

## Set up kafka broker
`docker-compose up -d`

## Producer
`go run main.go producer`

## Consumer
`go run main.go consumer` 

**Note:** *To start multiples consumers, run this command in different terminals concurrently*.
