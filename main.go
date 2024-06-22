// Description: This file is a simple example of how to run a local llama.cpp client.
// TODO:
// 1. flags for models path, and flags
// 2. interate over all models
// 3. pull in twitch chat and run the model on it
// 4. standard evalutation of the model

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

// Client contains the necessary information to run a local llama.cpp client.
type Client struct {
	BinPath string
	Args    []string
}

type Args struct {
	modelPath       string
	prompt          string
	file            string
	n               int  // number of tokens
	e               bool // process escape sequences
	logDisable      bool
	temperature     float64
	presencePenalty float64
}

func (c *Client) createExecArgs(args Args) {
	cmdArgs := []string{"-m", args.modelPath, "-p", fmt.Sprintf("\"%s\"", args.prompt), "-f", "pedro.txt", "-n", strconv.Itoa(args.n)}
	if args.e {
		cmdArgs = append(cmdArgs, "-e")
	}
	if args.logDisable {
		cmdArgs = append(cmdArgs, "--log-disable")
	}
	if args.presencePenalty != 0 {
		cmdArgs = append(cmdArgs, "--presence-penalty", strconv.FormatFloat(args.presencePenalty, 'f', -1, 64))
	}
	if args.temperature != 0 {
		cmdArgs = append(cmdArgs, "--temp", strconv.FormatFloat(args.temperature, 'f', -1, 64))
	}

	c.Args = cmdArgs
}

func main() {
	client := &Client{
		BinPath: "../llama.cpp/llama-cli",
	}

	args := Args{
		modelPath:       "../llama.cpp/models/pedro/microsoft_phi_3_mini_4k_instruct.gguf",
		prompt:          "what is a golang interface?",
		n:               75,
		e:               true,
		logDisable:      true,
		temperature:     0.7,
		presencePenalty: 1.0,
	}
	client.createExecArgs(args)

	cmd := exec.Command(client.BinPath, client.Args...)

	fmt.Println("Running command: ", cmd.String())
	stdOut, err := cmd.Output()
	if err != nil {
		log.Fatal(fmt.Errorf("did not run command | %w", err))
	}
	println(string(stdOut))
}
