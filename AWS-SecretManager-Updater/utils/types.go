package utils

import "github.com/aws/aws-sdk-go-v2/service/secretsmanager"

type SecretManagerClient struct {
	Service *secretsmanager.Client
}

type SecretData struct {
	SecretID   string
	SecretJSON string
}
