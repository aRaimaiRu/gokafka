package main

import (
	"fmt"
	"producer/controllers"
	"producer/services"
	"strings"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("kafka.server", viper.GetStringSlice("db.driver"))
	producer, err := sarama.NewSyncProducer(viper.GetStringSlice("kafka.servers"), nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	eventProducer := services.NewEventProducer(producer)
	accountService := services.NewAccountService(eventProducer)
	accountController := controllers.NewOpenAccountController(accountService)

	app := fiber.New()

	app.Post("/openaccount", accountController.OpenAccount)
	app.Post("/depositfund", accountController.DepositFund)
	app.Post("/withdrawfund", accountController.WithdrawFund)
	app.Post("/closeaccount", accountController.CloseAccount)

	app.Listen(":8080")
}
