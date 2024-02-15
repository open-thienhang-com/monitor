package inventory

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	configs "mono.thienhang.com/services/shop/configs"

	"github.com/spf13/cobra"
)

var (
	cfg        *configs.AppConfig
	rootCmd    *cobra.Command
	configPath string
	taskToRun  string
)

func init() {
	rootCmd = &cobra.Command{
		Short: "Example-Admin",
	}
	rootCmd.PersistentFlags().StringVar(&configPath, "configPath", ".", "Path to look for the config.yaml file in")
	rootCmd.PersistentFlags().StringVar(&taskToRun, "task", "", "Task to run")
	rootCmd.AddCommand(cmdServer)
}

func initializeCobra() {
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		zap.S().Fatal(fmt.Errorf("Error loading the configuration file: %s \n", err))
	}
	viper.SetEnvPrefix("app")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	overrideWithEnvVars()

	cfg = new(configs.AppConfig)
	if err = viper.Unmarshal(cfg); err != nil {
		zap.S().Fatal(fmt.Errorf("Error unmarshal the configuration file: %s \n", err))
	}
}

func overrideWithEnvVars() {
	for _, key := range viper.AllKeys() {
		viper.Set(key, viper.Get(key))
	}
}

func Execute() {
	cobra.OnInitialize(initializeCobra)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
