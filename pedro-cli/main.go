// Description: This file is a simple example of how to run a local llama.cpp client.
// TODO:
// 1. flags for models path, and flags
// 2. interate over all models
// 3. run the model with different prompts
// 4. standard evalutation of the model

package main

import (
	"log"

	"github.com/Soypete/llm-local-test-script/llamacpp"
)

var prompts = []string{
	"Pedro, what is a golang interface?",
	"Pedro, what are we doing today?",
	"Pedro, what is the meaning of life?",
	"Pedro, what is the best programming language?",
	"Pedro, tell me a dad joke.",
}

func main() {
	binPath := "../llama.cpp/llama-cli"
	modelPath := "../llama.cpp/models/pedro/mistral_7b_v1.gguf"

	args := llamacpp.Args{
		N:               75,
		E:               true,
		LogDisable:      true,
		Temperature:     0.7,
		PresencePenalty: 1.0,
	}
	for _, prompt := range prompts {
		args.Prompt = prompt
		client := llamacpp.SetupCLI(binPath, "cli", modelPath, args)
		output, err := client.RunCli()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(output)
	}
}
