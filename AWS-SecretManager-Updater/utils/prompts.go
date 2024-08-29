package utils

import (
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/tidwall/gjson"
)

func PromptSecretID() string {
	prompt := promptui.Prompt{
		Label: "Enter the Secret ID",
	}
	secretID, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return secretID
}

func SelectKeyFromJSON(secretJSON string) string {
	keys := gjson.Get(secretJSON, "@this").Map()

	var keyOptions []string
	for key := range keys {
		keyOptions = append(keyOptions, key)
	}

	keyPrompt := promptui.Select{
		Label: "Select key to Update",
		Items: keyOptions,
	}
	_, selectedKey, err := keyPrompt.Run()
	if err != nil {
		log.Fatalf("Failed to select key: %v", err)
	}
	return selectedKey
}

func PromptNewValue(key string) string {
	valuePrompt := promptui.Prompt{
		Label: fmt.Sprintf("Enter new value for %s", key),
	}
	newValue, err := valuePrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed: %v\n", err)
	}
	return newValue
}

func PromptConfirmation(key string, newValue string) string {
	conformationPrompt := promptui.Prompt{
		Label: fmt.Sprintf("[y/n] Are you sure you want to update value for key %s with %s", key, newValue),
		Validate: func(s string) error {
			if s == "y" || s == "n" {
				return nil
			}
			return fmt.Errorf("invalid input: please enter 'y' or 'n'")
		},
	}
	for {
		result, err := conformationPrompt.Run()
		if err != nil {
			return ""
		}

		if result == "y" || result == "n" {
			return result
		}
		//If not valid the loop continues
		fmt.Println("Invalid input. Please enter 'y' or 'n'.")
	}

}
