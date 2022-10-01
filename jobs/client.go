package jobs

import (
	"os"

	"github.com/hibiken/asynq"
)

var client *asynq.Client

func setUpClient() {

	redisAddr := os.Getenv("REDIS_URL")

	client = asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
}

func GetTaskClient() *asynq.Client {

	if client == nil {
		setUpClient()
	}
	return client
}