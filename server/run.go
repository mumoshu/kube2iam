package server

import (
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/jtblin/kube2iam/iam"
	"github.com/jtblin/kube2iam/iptables"
	"github.com/jtblin/kube2iam/version"
)

func (s *Server) Run() {
	logLevel, err := log.ParseLevel(s.LogLevel)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if s.Verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(logLevel)
	}

	if s.Version {
		version.PrintVersionAndExit()
	}

	if s.BaseRoleARN != "" {
		if !iam.IsValidBaseARN(s.BaseRoleARN) {
			log.Fatalf("Invalid --base-role-arn specified, expected: %s", iam.ARNRegexp.String())
		}
		if !strings.HasSuffix(s.BaseRoleARN, "/") {
			s.BaseRoleARN += "/"
		}
	}

	if s.AutoDiscoverBaseArn {
		if s.BaseRoleARN != "" {
			log.Fatal("--auto-discover-base-arn cannot be used if --base-role-arn is specified")
		}
		arn, err := iam.GetBaseArn()
		if err != nil {
			log.Fatalf("%s", err)
		}
		log.Infof("base ARN autodetected, %s", arn)
		s.BaseRoleARN = arn
	}

	if s.AutoDiscoverDefaultRole {
		if s.DefaultIAMRole != "" {
			log.Fatalf("You cannot use --default-role and --auto-discover-default-role at the same time")
		}
		arn, err := iam.GetBaseArn()
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.BaseRoleARN = arn
		instanceIAMRole, err := iam.GetInstanceIAMRole()
		if err != nil {
			log.Fatalf("%s", err)
		}
		s.DefaultIAMRole = instanceIAMRole
		log.Infof("Using instance IAMRole %s%s as default", s.BaseRoleARN, s.DefaultIAMRole)
	}

	if s.AddIPTablesRule {
		if err := iptables.AddRule(s.AppPort, s.MetadataAddress, s.HostInterface, s.HostIP); err != nil {
			log.Fatalf("%s", err)
		}
	}

	if err := s.Serve(s.APIServer, s.APIToken, s.Insecure); err != nil {
		log.Fatalf("%s", err)
	}
}
