package main

import (
	"strings"
	"testing"
)

func TestClient_createExecCmd(t *testing.T) {
	tests := []struct {
		name   string
		client *Client
		args   Args
		want   string
	}{
		{
			name: "all flags",
			client: &Client{
				BinPath: "../llama.cpp/main",
			},
			args: Args{
				modelPath:  "../llama.cpp/models/mistral/mistral_7b_v1.gguf",
				prompt:     "between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?",
				n:          400,
				e:          true,
				logDisable: true,
			},
			want: "-m ../llama.cpp/models/mistral/mistral_7b_v1.gguf -p \"between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?\" -n 400 -e --log-disable",
		},
		{
			name: "min flags",
			client: &Client{
				BinPath: "../ml/llama.cpp/main",
			},
			args: Args{
				modelPath:  "models/mistral/mistral_7b_v1.gguf",
				prompt:     "between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?",
				n:          100,
				e:          false,
				logDisable: false,
			},
			want: "-m models/mistral/mistral_7b_v1.gguf -p \"between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?\" -n 100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.client.createExecArgs(tt.args); strings.Join(tt.client.Args, " ") != tt.want {
				t.Errorf("Client.createExecArgs() = \n%v, want \n%v", strings.Join(tt.client.Args, " "), tt.want)
			}
		})
	}
}
