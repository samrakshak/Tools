package main

import (
	"fmt"
	"log"
	"tool/utils"
)

func main() {
	smClient, err := utils.NewSecretManagerClient()
	if err != nil {
		log.Fatalf("Failed to create Secrets Manager client: %v", err)
	}

	//Prompt for Secret ID
	secretID := utils.PromptSecretID()

	//Get the current secret value
	secretJSON, err := smClient.GetSecretValue(secretID)
	if err != nil {
		log.Fatalf("Failed to get secret value: %v", err)
	}

	//Prompt to select a key
	selectedKey := utils.SelectKeyFromJSON(secretJSON)

	//Prompt to the new value
	newValue := utils.PromptNewValue(selectedKey)

	//Prompt Are you sure you want to update?
	validationResopnse := utils.PromptConfirmation(selectedKey, newValue)
	if validationResopnse == "y" {
		//Update the secret
		updatedSecretJSON, err := smClient.UpdateSecretKey(secretID, selectedKey, newValue)
		if err != nil {
			log.Fatalf("Failed to update secret: %v", err)
		}

		fmt.Println("Secret updated succesfully.")
		fmt.Printf("Updated Secret JSON: %s\n", updatedSecretJSON)
	}
}
