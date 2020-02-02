package cmd

import (
	// "context"
	// "database/sql"
	// "net/http"
  "github.com/svennela/sample-broker/utils"
)

const (
	apiUserProp     = "api.user"
	apiPasswordProp = "api.password"
	apiPortProp     = "api.port"
)

func serve() {
  logger := utils.NewLogger("sample-broker")
  logger.Info(" Server Function" )
}
