package config

import (
	"os"
	"strconv"
)

type QueueConfig struct {
	Host      string
	QueueName string
	TaskName  string
	Delay     int
}

func GetQueueConfig() *QueueConfig {

	delay, _ := strconv.Atoi(os.Getenv("Q_DELAY"))

	return &QueueConfig{
		Host:      os.Getenv("Q_HOST"),
		QueueName: os.Getenv("Q_NAME"),
		TaskName:  os.Getenv("Q_TASK_NAME"),
		Delay:     delay,
	}
}
