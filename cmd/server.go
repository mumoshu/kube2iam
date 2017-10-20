package cmd

import (
	"github.com/spf13/pflag"

	"github.com/jtblin/kube2iam/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run a kube2iam server serves kube2iam agents so that they can provide pods creds for requested roles",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		s.Run()
	},
}

// addServerFlags adds the command line flags.
func addServerFlags(s *server.Server, fs *pflag.FlagSet) {
	fs.StringVar(&s.APIServer, "api-server", s.APIServer, "Endpoint for the api server")
	fs.StringVar(&s.APIToken, "api-token", s.APIToken, "Token to authenticate with the api server")
	fs.StringVar(&s.AppPort, "app-port", s.AppPort, "Http port")
	fs.StringVar(&s.BaseRoleARN, "base-role-arn", s.BaseRoleARN, "Base role ARN")
	fs.BoolVar(&s.Debug, "debug", s.Debug, "Enable debug features")
	fs.StringVar(&s.DefaultIAMRole, "default-role", s.DefaultIAMRole, "Fallback role to use when annotation is not set")
	fs.StringVar(&s.IAMRoleKey, "iam-role-key", s.IAMRoleKey, "Pod annotation key used to retrieve the IAM role")
	fs.BoolVar(&s.Insecure, "insecure", false, "Kubernetes server should be accessed without verifying the TLS. Testing only")
	fs.StringVar(&s.MetadataAddress, "metadata-addr", s.MetadataAddress, "Address for the ec2 metadata")
	fs.BoolVar(&s.AddIPTablesRule, "iptables", false, "Add iptables rule (also requires --host-ip)")
	fs.BoolVar(&s.AutoDiscoverBaseArn, "auto-discover-base-arn", false, "Queries EC2 Metadata to determine the base ARN")
	fs.BoolVar(&s.AutoDiscoverDefaultRole, "auto-discover-default-role", false, "Queries EC2 Metadata to determine the default Iam Role and base ARN, cannot be used with --default-role, overwrites any previous setting for --base-role-arn")
	fs.StringVar(&s.HostInterface, "host-interface", "docker0", "Host interface for proxying AWS metadata")
	fs.BoolVar(&s.NamespaceRestriction, "namespace-restrictions", false, "Enable namespace restrictions")
	fs.StringVar(&s.NamespaceKey, "namespace-key", s.NamespaceKey, "Namespace annotation key used to retrieve the IAM roles allowed (value in annotation should be json array)")
	fs.StringVar(&s.HostIP, "host-ip", s.HostIP, "IP address of host")
	fs.DurationVar(&s.BackoffMaxInterval, "backoff-max-interval", s.BackoffMaxInterval, "Max interval for backoff when querying for role.")
	fs.DurationVar(&s.BackoffMaxElapsedTime, "backoff-max-elapsed-time", s.BackoffMaxElapsedTime, "Max elapsed time for backoff when querying for role.")
	fs.StringVar(&s.LogLevel, "log-level", s.LogLevel, "Log level")
	fs.BoolVar(&s.Verbose, "verbose", false, "Verbose")
	fs.BoolVar(&s.Version, "version", false, "Print the version and exits")
}

var s *server.Server

func init() {
	s = server.NewServer()
	addServerFlags(s, serverCmd.Flags())

	rootCmd.AddCommand(serverCmd)
}
