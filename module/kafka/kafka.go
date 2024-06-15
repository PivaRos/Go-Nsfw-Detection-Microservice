package kafka

import "log"

func Run() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("kafka thread: top level panic:", r)
		} else {
			log.Println("kafka: couldn't recover panic in top level")
		}
	}()

}
