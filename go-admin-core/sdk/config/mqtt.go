package config

type Mqtt struct {
	Host    string
	Port string
	Prefix string
}

var MqttConfig = new(Mqtt)

