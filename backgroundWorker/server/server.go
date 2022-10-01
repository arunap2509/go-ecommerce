package main

import (
	"log"
	"os"

	"github.com/arunap2509/ecommerce/jobs"
	"github.com/hibiken/asynq"
)

func main() {

	redisAddr := os.Getenv("REDIS_URL")
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(jobs.TypeOrderedToTransit, jobs.InvokeOrderedToInTransit)
	mux.HandleFunc(jobs.TypeInTransitToDelivered, jobs.InvokeInTransitToDelivered)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
