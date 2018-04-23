package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/speix/aws-lambda-go-cron-case/app/config"
	"github.com/speix/aws-lambda-go-cron-case/app/data"
)

func main() {
	lambda.Start(Handler)
}

func Handler() {
	database := config.NewDB(config.GetDBConfig())
	defer database.DB.Close()

	reservations, err := data.ReservationData{Database: database}.GetAll()
	if err != nil {
		log.Printf(err.Error())
		return
	}

	request := data.RequestData{Reservations: reservations, QueueConfig: config.GetQueueConfig()}
	if err = request.QueueTask(); err != nil {
		log.Printf(err.Error())
		return
	}
}
