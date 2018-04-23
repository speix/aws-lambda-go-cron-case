package data

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/speix/aws-lambda-go-cron-case/app/config"

	"github.com/speix/aws-lambda-go-cron-case/app"
)

type RequestData struct {
	app.Reservations
	*config.QueueConfig
}

type responseBody struct {
	Message string
}

func (r *RequestData) QueueTask() error {

	if len(r.Reservations) == 0 {
		return errors.New("No records found ")
	}

	client := &http.Client{
		Timeout: time.Duration(60 * time.Second),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := r.QueueConfig.Host
	container := &app.RequestContainer{}
	container.QueueName = r.QueueConfig.QueueName
	container.TaskName = r.QueueConfig.TaskName
	container.Delay = r.QueueConfig.Delay

	for item := range r.Reservations {
		container.Tasks = append(container.Tasks,
			app.Message{
				Template:   r.Reservations[item].Template,
				Phone:      r.Reservations[item].Phone,
				Attributes: r.Reservations[item]})
	}

	body := new(bytes.Buffer)
	err := json.NewEncoder(body).Encode(container)
	if err != nil {
		return nil
	}

	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		b := responseBody{}
		json.NewDecoder(response.Body).Decode(&b)

		return errors.New(b.Message)
	}

	return nil
}
