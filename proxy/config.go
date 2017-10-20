package proxy

import "errors"

type Config struct {
	Kube2IamServer string
	StateDir       string
}

func (c Config) Validate() error {
	if c.Kube2IamServer == "" {
		return errors.New("kube2iam server cannot be empty")
	}
	return nil
}
