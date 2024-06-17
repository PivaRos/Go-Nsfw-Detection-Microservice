package utils

import (
	"database/sql"

	"github.com/IBM/sarama"
	"github.com/pivaros/go-image-recognition/kafka/nsfw"
)

type AppState struct {
	Env                *Env
	Db                 *sql.DB
	ProduceMessage     func(message *sarama.ProducerMessage)
	ClassificationFunc func(string) (nsfw.Result, error)
}
