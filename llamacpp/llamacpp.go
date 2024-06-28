package llamacpp

import (
	"fmt"
	"os/exec"
	"strconv"
)

// Client that allows for running a local llama.cpp client from your go code.
// This client can run llama.cpp in cli or server mode. Must use this tag or newer b3138 to
// run the correct binary.
type Client struct {
	BinPath    string
	ModelPath  string
	Args       []string
	serverMode bool
}

type Args struct {
	Prompt          string
	N               int  // number of tokens
	E               bool // process escape sequences
	LogDisable      bool
	Temperature     float64
	PresencePenalty float64
}

// createExecArts only gets run in cli mode. And cli mode only accepts one promt and model path at a time.
func (c *Client) createExecArgs(args Args) []string {
	cmdargs := []string{"-m", c.ModelPath, "-p", fmt.Sprintf("\"%s\"", args.Prompt), "-n", strconv.Itoa(args.N)}
	if args.E {
		cmdargs = append(cmdargs, "-e")
	}
	if args.LogDisable {
		cmdargs = append(cmdargs, "--log-disable")
	}
	if args.PresencePenalty != 0 {
		cmdargs = append(cmdargs, "--presence-penalty", strconv.FormatFloat(args.PresencePenalty, 'f', -1, 64))
	}
	if args.Temperature != 0 {
		cmdargs = append(cmdargs, "--temp", strconv.FormatFloat(args.Temperature, 'f', -1, 64))
	}

	c.Args = cmdargs
	return cmdargs
}

// Setup creates a new client with the given binary path, mode, and model paths.
// If the mode is "server" then the client will run in server mode. If the mode is "cli" then the client will run in cli mode.
// args are of type Args and are used to pass in the necessary arguments to llama.cpp to run the model. They are only used in cli mode.
// TODO: should we change the name of Args to options
func SetupCLI(binPath, mode, modelPath string, args Args) *Client {
	client := &Client{
		BinPath:    binPath,
		ModelPath:  modelPath,
		serverMode: true,
	}
	client.createExecArgs(args)
	return client
}

func SetupServer(binPath, mode, modelPath string) *Client {
	client := &Client{
		BinPath:   binPath,
		ModelPath: modelPath,
	}
	client.serverMode = true
	return client
}

// RunCli runs the llama.cpp client in cli mode with the given arguments.
func (c *Client) RunCli() (string, error) {
	cmd := exec.Command(c.BinPath, c.Args...)

	fmt.Println("")
	fmt.Println("Running command: ", cmd.String())
	stdOut, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w", err)
	}
	return string(stdOut), nil
}
