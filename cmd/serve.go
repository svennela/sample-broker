package cmd

import (
	// "context"
	// "database/sql"
	// "net/http"
  "github.com/svennela/sample-broker/utils"
  "github.com/svennela/sample-broker/pkg/toggles"
  "github.com/pivotal-cf/brokerapi"
  "github.com/spf13/viper"
  "code.cloudfoundry.org/lager"
  "github.com/spf13/cobra"
)

const (
	apiUserProp     = "api.user"
	apiPasswordProp = "api.password"
	apiPortProp     = "api.port"
)

var cfCompatibilityToggle = toggles.Features.Toggle("enable-cf-sharing", false, `Set all services to have the Sharable flag so they can be shared
	across spaces in PCF.`)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "serve",
		Short: "Start the service broker",
		Long: `Starts the service broker listening on a port defined by the
	PORT environment variable.`,
		Run: func(cmd *cobra.Command, args []string) {
			serve()
		},
	})

	viper.BindEnv(apiUserProp, "SECURITY_USER_NAME")
	viper.BindEnv(apiPasswordProp, "SECURITY_USER_PASSWORD")
	viper.BindEnv(apiPortProp, "PORT")
}

func serve() {
  logger := utils.NewLogger("sample-broker")
  logger.Info(" Server Function" )

  credentials := brokerapi.BrokerCredentials{
		Username: viper.GetString(apiUserProp),
		Password: viper.GetString(apiPasswordProp),
	}
  logger.Info("credentials catalog", lager.Data{"credentials": credentials})

	if cfCompatibilityToggle.IsActive() {
		logger.Info("Enabling Cloud Foundry service sharing")
    logger.Info("credentials catalog", lager.Data{"credentials": credentials})
		//serviceBroker = server.NewCfSharingWrapper(serviceBroker)
	}

}
