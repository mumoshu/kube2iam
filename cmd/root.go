package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/kubernetes-incubator/kube-aws/core/root/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kubernetes-aws-authenticator",
	Short: "A tool to authenticate to Kubernetes using AWS IAM credentials",
}

// Execute the CLI entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Load configuration from `filename`")

	rootCmd.PersistentFlags().StringP(
		"cluster-id",
		"i",
		"",
		"Specify the cluster `ID`, a unique-per-cluster identifier for your kubernetes-aws-authenticator installation.",
	)
	viper.BindPFlag("clusterID", rootCmd.PersistentFlags().Lookup("cluster-id"))
	viper.BindEnv("clusterID", "KUBERNETES_AWS_AUTHENTICATOR_CLUSTER_ID")
}

func initConfig() {
	if cfgFile == "" {
		return
	}
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Can't read configuration file %q: %v\n", cfgFile, err)
		os.Exit(1)
	}
}

func getConfig() (config.Config, error) {
	config := config.Config{
		ClusterID:              viper.GetString("clusterID"),
		LocalhostPort:          viper.GetInt("server.port"),
		GenerateKubeconfigPath: viper.GetString("server.generateKubeconfig"),
		StateDir:               viper.GetString("server.stateDir"),
	}
	if err := viper.UnmarshalKey("server.mapRoles", &config.StaticRoleMappings); err != nil {
		return config, fmt.Errorf("invalid server role mappings: %v", err)
	}

	if config.ClusterID == "" {
		return config, errors.New("cluster ID cannot be empty")
	}

	return config, nil
}
