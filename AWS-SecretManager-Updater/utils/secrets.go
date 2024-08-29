package utils

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/tidwall/sjson"
)

func NewSecretManagerClient() (*SecretManagerClient, error) {
	client, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config,%v", err)
	}

	svc := secretsmanager.NewFromConfig(client)
	return &SecretManagerClient{Service: svc}, nil
}

func (sm *SecretManagerClient) GetSecretValue(secretID string) (string, error) {
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretID),
	}
	result, err := sm.Service.GetSecretValue(context.TODO(), input)
	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}

func (sm *SecretManagerClient) UpdateSecretValue(secretID string, updatedJSON string) error {
	input := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(secretID),
		SecretString: aws.String(updatedJSON),
	}

	_, err := sm.Service.UpdateSecret(context.TODO(), input)
	return err
}

func (sm *SecretManagerClient) UpdateSecretKey(secretID string, key, newValue string) (string, error) {
	secretJSON, err := sm.GetSecretValue(secretID)
	if err != nil {
		return "", err
	}

	updatedJSON, err := sjson.Set(secretJSON, key, newValue)
	if err != nil {
		return "", fmt.Errorf("Failed to update JSON key: %v", err)
	}

	err = sm.UpdateSecretValue(secretID, updatedJSON)
	if err != nil {
		return "", fmt.Errorf("Failed to update sercert: %v", err)
	}

	return updatedJSON, nil
}
