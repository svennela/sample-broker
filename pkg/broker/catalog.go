package broker

import (
	"github.com/pivotal-cf/brokerapi"
)

// Service overrides the canonical Service Broker service type using a custom
// type for Plans, everything else is the same.
type Service struct {
	brokerapi.Service

	Plans []ServicePlan `json:"plans"`
}

// ToPlain converts this service to a plain PCF Service definition.
func (s Service) ToPlain() brokerapi.Service {
	plain := s.Service
	plainPlans := []brokerapi.ServicePlan{}

	for _, plan := range s.Plans {
		plainPlans = append(plainPlans, plan.ServicePlan)
	}

	plain.Plans = plainPlans

	return plain
}

// ServicePlan extends the OSB ServicePlan by including a map of key/value
// pairs that can be used to pass additional information to the back-end.
type ServicePlan struct {
	brokerapi.ServicePlan

	ServiceProperties  map[string]string      `json:"service_properties"`
	ProvisionOverrides map[string]interface{} `json:"provision_overrides,omitempty"`
	BindOverrides      map[string]interface{} `json:"bind_overrides,omitempty"`
}

// GetServiceProperties gets the plan settings variables as a string->interface map.
func (sp *ServicePlan) GetServiceProperties() map[string]interface{} {
	props := make(map[string]interface{})
	for k, v := range sp.ServiceProperties {
		props[k] = v
	}

	return props
}
