# Bus Library: Implementing the Publish/Subscribe Paradigm in Golang

[![Go Report Card](https://goreportcard.com/badge/github.com/e154/bus)](https://goreportcard.com/report/github.com/e154/bus)
[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)
[![codecov](https://codecov.io/github/e154/bus/graph/badge.svg?token=9DQPJQ4OW5)](https://codecov.io/github/e154/bus)

| Branch | Status                                                                                           |
|--------|--------------------------------------------------------------------------------------------------|
| master | ![Build Status](https://github.com/e154/bus/actions/workflows/test.yml/badge.svg?branch=main)    |
| dev    | ![Build Status](https://github.com/e154/bus/actions/workflows/test.yml/badge.svg?branch=develop) |

### Overview

The Bus library is a tool designed for convenient implementation of the publish/subscribe paradigm in the Go programming language. The publish/subscribe paradigm is widely used to facilitate efficient message exchange between system components, distinguishing message producers (publishers) and message consumers (subscribers).

### Features

1. The Bus library supports a powerful topic system. Topics can contain `+` and `#` symbols, allowing for flexible topic hierarchies. For example, the topic "home/+/temperature" can subscribe to all temperature changes in different rooms.

2. One of the key advantages of the Bus library is its ability to transmit any data type based on structures. This provides flexibility and allows for the transmission of complex data between system components.

### Library Interface

```go
type Bus interface {
    Publish(topic string, args ...interface{})
    CloseTopic(topic string)
    Subscribe(topic string, fn interface{}, options ...interface{}) error
    Unsubscribe(topic string, fn interface{}) error
    Stat(ctx context.Context, limit, offset int64, orderBy, sort string) (stats Stats, total int64, err error)
    Purge()
}
```

### Core Methods

- **Publish(topic string, args ...interface{})**: Publishes a message to the specified topic with arbitrary arguments.

- **CloseTopic(topic string)**: Closes a topic, causing all subscribers to unsubscribe from that topic.

- **Subscribe(topic string, fn interface{}, options ...interface{}) error**: Subscribes to a topic with a specified handler function and additional options.

- **Unsubscribe(topic string, fn interface{}) error**: Unsubscribes from a topic for the specified handler function.

### Additional Methods

- **Stat(ctx context.Context, limit, offset int64, orderBy, sort string) (stats Stats, total int64, err error)**: Retrieves statistics on the usage of topics with pagination and sorting options.

- **Purge()**: Clears all topics and unsubscribes from all subscriptions.

### Example

```go
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
```

### Conclusion

The Bus library provides efficient tools for implementing the publish/subscribe paradigm in the Go programming language. Its use simplifies communication between system components, providing flexibility and ease of use.

### Contributors

- [Alex Filippov](https://github.com/e154)

All contributors are welcome. If you would like to contribute, please adhere to the following rules.

- Pull requests will be accepted only in the "develop" branch.
- All modifications or additions should be tested.

Thank you for your understanding!

### LICENSE

[GPLv3 Public License](https://github.com/e154/bus/blob/master/LICENSE)
