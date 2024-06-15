package utils

import (
	"database/sql"

	"github.com/IBM/sarama"
)

type AppState struct {
	Env            *Env
	Db             *sql.DB
	ProduceMessage func(message *sarama.ProducerMessage)
}
