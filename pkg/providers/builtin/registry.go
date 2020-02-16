package builtin

import (
	"github.com/svennela/sample-broker/pkg/broker"
)

// NOTE(josephlewis42) unless there are extenuating circumstances, as of 2019
// no new builtin providers should be added. Instead, providers should be
// added using downloadable brokerpaks.

// BuiltinBrokerRegistry creates a new registry with all the built-in brokers
// added to it.
func BuiltinBrokerRegistry() broker.BrokerRegistry {
	out := broker.BrokerRegistry{}
	RegisterBuiltinBrokers(out)
	return out
}

// RegisterBuiltinBrokers adds the built-in brokers to the given registry.
func RegisterBuiltinBrokers(registry broker.BrokerRegistry) {
}
