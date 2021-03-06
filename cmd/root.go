package cmd

import (
	"fmt"
	"os"
  "log"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
	"github.com/svennela/sample-broker/utils"
)
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "sample-broker",
	Short: "Service Broker is an OSB compatible service broker",
	Long:  `An OSB compatible service broker for GCP/AZURE/AWS.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WARNING: In the future running the broker from the root..")
		fmt.Println("WARNING: command will show help instead...")
		fmt.Println("WARNING: Update your scripts to run service-broker serve..")

		serve()
	},
}

func init() {
  cobra.OnInitialize(initConfig)
	fmt.Println("config file",cfgFile)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Configuration file to be read")
	viper.SetEnvPrefix(utils.EnvironmentVarPrefix)
	viper.SetEnvKeyReplacer(utils.PropertyToEnvReplacer)
	viper.AutomaticEnv()

}

func initConfig() {
  if cfgFile == "" {
  		return
  	}

		fmt.Println("setting config file",cfgFile)

  	viper.SetConfigFile(cfgFile)

  	if err := viper.ReadInConfig(); err != nil {
  		log.Fatalf("Can't read config: %v\n", err)
  	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
