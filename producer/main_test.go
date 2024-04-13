package main

// import (
// 	"fmt"

// 	"github.com/IBM/sarama"
// )

// func main() {
// 	servers := []string{"localhost:9092"}

// 	producer, err := sarama.NewSyncProducer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer producer.Close()

// 	msg := sarama.ProducerMessage{
// 		Topic: "firsthello",
// 		Value: sarama.StringEncoder("Hello World"),
// 	}
// 	p, o, err := producer.SendMessage(&msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("partition=%v, offset=%v", p, o)
// }
