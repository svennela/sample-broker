package brokers

import (
	"fmt"

	"github.com/svennela/sample-broker/pkg/broker"
//	"github.com/svennela/sample-broker/pkg/brokerpak"
	"github.com/svennela/sample-broker/pkg/providers/builtin"
	"github.com/svennela/sample-broker/utils"
	"golang.org/x/oauth2/jwt"
)

type BrokerConfig struct {
	HttpConfig *jwt.Config
	ProjectId  string
	Registry   broker.BrokerRegistry
}

func NewBrokerConfigFromEnv() (*BrokerConfig, error) {
	fmt.Println("NewBrokerConfigFromEnv")
	projectId, err := utils.GetDefaultProjectId()
	if err != nil {
		return nil, err
	}

	conf, err := utils.GetAuthedConfig()
	if err != nil {
		return nil, err
	}

	registry := builtin.BuiltinBrokerRegistry()
	// if err := brokerpak.RegisterAll(registry); err != nil {
	// 	return nil, fmt.Errorf("Error loading brokerpaks: %v", err)
	// }

	return &BrokerConfig{
		ProjectId:  projectId,
		HttpConfig: conf,
		Registry:   registry,
	}, nil
}
