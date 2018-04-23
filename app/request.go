package app

type RequestContainer struct {
	QueueName string    `json:"queue"`
	TaskName  string    `json:"task"`
	Tasks     []Message `json:"messages"`
	Delay     int       `json:"delay"`
}

type Message struct {
	Template   string      `json:"template"`
	Phone      string      `json:"to"`
	Attributes Reservation `json:"attributes"`
}

type RequestService interface {
	QueueTask() error
}
