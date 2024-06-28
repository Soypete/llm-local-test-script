package main

import "fmt"

func main() {
	// connect to supabase for per
	// TODO: add supabase connection secrets
	db, err := database.NewPostgres()
	if err != nil {
		log.Fatalln(err)
	}

	llm, err := langchain.Setup(db)
	if err != nil {
		log.Fatalln(err)
	}
	if llm == nil {
		log.Fatalln("llm is nil")
	}

	client := &Client{
		BinPath:   "../llama.cpp/llama-server",
		modelPath: "../llama.cpp/models/pedro/mistral_7b_v1.gguf",
	}

	args := Args{
		logDisable:      true,
	}
	// TODO: add waitgroup and start server
	wg := sync.WaitGroup{}

	// TODO: look over models for setup? how do we do that?
	wg.Add(1)
	// TODO: instead of iterating over models, what if we iterated over functions... :thinking:
	go func() {
		// run the server


	}

	// call loop to generate prompts and send them to the server

	// persist the generated to the database
}
