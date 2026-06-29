package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type TelemetryPayload struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Light       int     `json:"light"`
}

// TelemetrySaver menerima data telemetri yang diterima dari MQTT.
type TelemetrySaver func(temp float32, humidity float32, light int) error

// InitMQTT connects to the MQTT broker and starts the subscriber worker
func InitMQTT(saveTelemetry TelemetrySaver) mqtt.Client {
	host := os.Getenv("MQTT_HOST")
	port := os.Getenv("MQTT_PORT")
	clientID := os.Getenv("MQTT_CLIENT_ID")
	user := os.Getenv("MQTT_USER")
	password := os.Getenv("MQTT_PASSWORD")

	broker := fmt.Sprintf("tcp://%s:%s", host, port)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	
	if user != "" {
		opts.SetUsername(user)
		opts.SetPassword(password)
	}

	opts.OnConnect = func(c mqtt.Client) {
		log.Println("Connected to MQTT Broker!")
		// Subscribe to telemetry topic
		topic := "kolam/aquarium/telemetry"
		if token := c.Subscribe(topic, 0, getMessageHandler(saveTelemetry)); token.Wait() && token.Error() != nil {
			log.Printf("Failed to subscribe to %s: %v", topic, token.Error())
		} else {
			log.Printf("Subscribed to topic: %s", topic)
		}
	}

	opts.OnConnectionLost = func(c mqtt.Client, err error) {
		log.Printf("MQTT connection lost: %v", err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Printf("Error connecting to MQTT Broker: %v", token.Error())
	}

	return client
}

// getMessageHandler returns the callback function when a message is received
func getMessageHandler(saveTelemetry TelemetrySaver) mqtt.MessageHandler {
	return func(client mqtt.Client, msg mqtt.Message) {
		go func() {
			log.Printf("Received message on topic: %s\nPayload: %s", msg.Topic(), msg.Payload())

			var payload TelemetryPayload
			err := json.Unmarshal(msg.Payload(), &payload)
			if err != nil {
				log.Printf("Error parsing MQTT payload: %v", err)
				return
			}

			// Call the provided save callback to store data
			saveTelemetry(payload.Temperature, payload.Humidity, payload.Light)
		}()
	}
}
