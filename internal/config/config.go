package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Mysql struct {
		DataSource string
	}
	Kafka KafkaConfig
}

type KafkaConfig struct {
	Hosts []string
}
