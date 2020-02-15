package brokers

import (
	"fmt"

	"github.com/svennela/sample-broker/pkg/broker"
	"github.com/svennela/sample-broker/pkg/brokerpak"
)

type BrokerConfig struct {
	Registry   broker.BrokerRegistry
}

func NewBrokerConfigFromEnv() (*BrokerConfig, error) {
	registry := broker.BrokerRegistry{}
	if err := brokerpak.RegisterAll(registry); err != nil {
		return nil, fmt.Errorf("Error loading brokerpaks: %v", err)
	}

	return &BrokerConfig{
		Registry:   registry,
	}, nil
}
