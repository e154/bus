package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/e154/bus"
)

type TemperatureReading struct {
	RoomID      string
	Temperature float64
}

type HumidityReading struct {
	RoomID   string
	Humidity float64
}

func main() {

	var messageBus = bus.NewBus()

	var eventHandler = func(topic string, msg interface{}) {
		switch v := msg.(type) {
		case TemperatureReading:
			fmt.Printf("topic: \"%s\", message: %v\n", topic, v)
		case HumidityReading:
			fmt.Printf("topic: \"%s\", message: %v\n", topic, v)
		}
	}

	messageBus.Subscribe("home/#", eventHandler)
	defer messageBus.Unsubscribe("home/#", eventHandler)

	go func() {
		for i := 0; i < 5; i++ {
			messageBus.Publish("home/living_room/temperature", TemperatureReading{
				RoomID:      "living_room",
				Temperature: 20.0 + float64(i),
			})
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			messageBus.Publish("home/kitchen/humidity", HumidityReading{
				RoomID:   "kitchen",
				Humidity: 40.0 + float64(i),
			})
			time.Sleep(time.Second * 3)
		}
	}()

	// ctrl + C
	var gracefulStop = make(chan os.Signal, 10)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	<-gracefulStop

}
