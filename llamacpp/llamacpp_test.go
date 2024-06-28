package llamacpp

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
				BinPath:   "../llama.cpp/main",
				ModelPath: "../llama.cpp/models/mistral/mistral_7b_v1.gguf",
			},
			args: Args{
				Prompt:     "between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?",
				N:          400,
				E:          true,
				LogDisable: true,
			},
			want: "-m ../llama.cpp/models/mistral/mistral_7b_v1.gguf -p \"between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?\" -n 400 -e --log-disable",
		},
		{
			name: "min flags",
			client: &Client{
				BinPath:   "../ml/llama.cpp/main",
				ModelPath: "models/mistral/mistral_7b_v1.gguf",
			},
			args: Args{
				Prompt:     "between little ceasers, mountain mikes, and papa murphy's where should I get my pizza tonight?",
				N:          100,
				E:          false,
				LogDisable: false,
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
