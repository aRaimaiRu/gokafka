package main
// package example

// import (
// 	"fmt"

// 	"github.com/IBM/sarama"
// )

// func main() {
// 	fmt.Println("main")
// 	servers := []string{"localhost:9092"}

// 	consumer, err := sarama.NewConsumer(servers, nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer consumer.Close()
// 	partitionConsumer, err := consumer.ConsumePartition("firsthello", 0, sarama.OffsetNewest)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer partitionConsumer.Close()
// 	fmt.Println("start")
// 	for {
// 		select {
// 		case err := <-partitionConsumer.Errors():
// 			fmt.Println(err)
// 		case msg := <-partitionConsumer.Messages():
// 			fmt.Println(string(msg.Value))
// 		}
// 	}

// }