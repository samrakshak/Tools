package main

import (
	"fmt"
	"log"
	"tool/utils"
)

// ANSI color codes
const (
	green = "\033[32m"
	blue  = "\033[34m"
	red   = "\033[0m"
)

func main() {

	//Mapping the options to a shorter value
	options := []struct {
		Display string
		Value   string
	}{
		{fmt.Sprintf("1. Update Existing Key/Value", blue), "update"},
		{fmt.Sprintf("2. Create New Key/Value", green), "create"},
		{fmt.Sprintf("3. Delete Existing Key/Value", red), "delete"},
	}

	// Extract the display options for the prompt
	items := make([]string, len(options))
	for i, option := range options {
		items[i] = option.Display
	}

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

	//Prompt to select Edit Delete or Update Json.
	index, _ := utils.SelectConfigOption()
	// Get the selected short value based on the index
	selectedValue := options[index].Value

	// Cases for updates
	if selectedValue == "create" {
		//Primpt new key and value
		newKey := utils.PromptNewKey()
		newValue := utils.PromptNewValue(newKey)
		createdSecretJSON, err := smClient.UpdateSecretKey(secretID, newKey, newValue)
		if err != nil {
			log.Fatalf("Failed to update secret: %v", err)
		}
		fmt.Println("Key/Value updated succesfully.")
		fmt.Printf("Updated Secret JSON: %s\n", createdSecretJSON)
	} else if selectedValue == "update" {
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

	} else if selectedValue == "delete" {
		fmt.Println("We can delete it later")
	}

}
