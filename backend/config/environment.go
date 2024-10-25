package config

import "errors"

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

func NewEnvironment(env string) (Environment, error) {
	switch env {
	case string(Dev):
		return Dev, nil
	case string(Prod):
		return Prod, nil
	default:
		return "", errors.New("invalid environment value")
	}
}
