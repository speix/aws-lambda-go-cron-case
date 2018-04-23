# aws-lambda-go-cron-case

This is a showcase of a golang cron service that runs on AWS Lambda with the help of AWS CloudWatch for scheduling.   
Its purpose is to collect database records with a given time interval and forward them to a message brokering system.

# Building your function

Preparing a binary to deploy to AWS Lambda requires that it is compiled for Linux and placed into a .zip file.  
For more information check out the official aws package: https://github.com/aws/aws-lambda-go

## For the sake of reference: Linux and macOS
``` shell
# Remember to build your handler executable for Linux!
GOOS=linux GOARCH=amd64 go build -o main .
zip main.zip main
```

### Required packages for this service:
```
- github.com/aws/aws-lambda-go/lambda
- github.com/lib/pq
- github.com/jmoiron/sqlx
```

### Environment variables
```
DB_HOST, DB_NAME, DB_USER, DB_PASS, DB_SSL_MODE
Q_HOST, Q_NAME, Q_TASK_NAME, Q_DELAY
```

### AWS CloudWatch Rules
```
Scheduled to run every half hour of every day from 09:00 to 21:00

Cron expression 0/30 9-21 ? * * *
```