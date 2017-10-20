package cmd

import (
	"errors"

	"github.com/jtblin/kube2iam/proxy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// DefaultPort is the default localhost port (chosen randomly).
const DefaultProxyPort = 8181

// serverCmd represents the server command
var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run a kube2iam proxy proxies credential requests from your pods to a kube2iam server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := getProxyConfig()

		if err != nil {
			logrus.Fatalf("%s", err)
		}

		proxy.New(config).Run()
	},
}

func getProxyConfig() (proxy.Config, error) {
	config := proxy.Config{
		Kube2IamServer: viper.GetString("kube2iam.server"),
		StateDir:       viper.GetString("server.state-dir"),
	}
	if config.Kube2IamServer == "" {
		return config, errors.New("kube2iam server cannot be empty")
	}

	return config, nil
}

func init() {
	proxyCmd.Flags().String("kube2iam-server",
		"kube2iam",
		"The endpoint of kube2iam server.")
	viper.BindPFlag("kube2iam.server", proxyCmd.Flags().Lookup("kube2iam-server"))

	proxyCmd.Flags().String("state-dir",
		"/var/kube2iam",
		"State `directory` for certificate and private key (should be a hostPath mount).")
	viper.BindPFlag("server.stateDir", proxyCmd.Flags().Lookup("state-dir"))

	rootCmd.AddCommand(proxyCmd)
}
