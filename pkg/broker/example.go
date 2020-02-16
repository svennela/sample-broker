package broker

import "github.com/svennela/sample-broker/pkg/validation"

// ServiceExample holds example configurations for a service that _should_
// work.
type ServiceExample struct {
	// Name is a human-readable name of the example.
	Name string `json:"name" yaml:"name"`
	// Description is a long-form description of what this example is about.
	Description string `json:"description" yaml:"description"`
	// PlanId is the plan this example will run against.
	PlanId string `json:"plan_id" yaml:"plan_id"`

	// ProvisionParams is the JSON object that will be passed to provision.
	ProvisionParams map[string]interface{} `json:"provision_params" yaml:"provision_params"`

	// BindParams is the JSON object that will be passed to bind. If nil,
	// this example DOES NOT include a bind portion.
	BindParams map[string]interface{} `json:"bind_params" yaml:"bind_params"`
}

var _ validation.Validatable = (*ServiceExample)(nil)

// Validate implements validation.Validatable.
func (action *ServiceExample) Validate() (errs *validation.FieldError) {
	return errs.Also(
		validation.ErrIfBlank(action.Name, "name"),
		validation.ErrIfBlank(action.Description, "description"),
		validation.ErrIfBlank(action.PlanId, "plan_id"),
	)
}
