/*
main.go

This file is part of the nictestcasegen project.
It contains the main entry point for the application.

The main package prints "Hello, World!" to the console.
Additional functionality can be added where indicated.

Usage:
	go run main.go

Author: Kishore
Created: 05-Oct-2023
*/

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type ModelConfig struct {
	name         string
	temperature  float64
	topP         float64
	topK         int
	maxOutTokens int
}

func main() {

	var ctx context.Context
	var client *genai.Client

	ModelConfig := ModelConfig{"gemini-pro", 0.1, 0.5, 20, 100}

	fmt.Println("Hello, World!")
	// Add your code here

	client, ctx = registerAPIKey()
	fmt.Println("Hello, World!")

	model := generateAndConfigureModel(client, ModelConfig)
	fmt.Println("Hello, World!")

	resp, err := model.GenerateContent(ctx, genai.Text("Write 4 line poem about the moon"))
	fmt.Println("Hello, World!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, World!")

	if resp.Candidates != nil {
		for _, v := range resp.Candidates {
			for _, k := range v.Content.Parts {
				fmt.Println(k.(genai.Text))
			}
		}
	}

	defer client.Close()
}

/*
registerAPIKey initializes a new context and registers the API key for the generative AI client.

It retrieves the API key from the environment variables and uses it to create a new generative AI client.
If the client creation fails, the function logs the error and terminates the program.

Returns:

	context.Context: The context initialized for the generative AI client.
*/
func registerAPIKey() (*genai.Client, context.Context) {

	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	if err != nil {
		log.Fatal(err)
	}
	//defer client.Close()

	return client, ctx
}

func generateAndConfigureModel(client *genai.Client, modelConfig ModelConfig) *genai.GenerativeModel {
	// Create a new model
	model := client.GenerativeModel(modelConfig.name)
	model.SetTemperature(0.1)
	model.SetTopP(0.5)
	model.SetTopK(20)
	model.SetMaxOutputTokens(100)

	return model
}
