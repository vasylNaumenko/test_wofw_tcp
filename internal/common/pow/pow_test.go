package pow

import "testing"

func TestGenerateChallenge(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"[success] Non-empty challenge", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateChallenge()
			if (got != "") != tt.want {
				t.Errorf("GenerateChallenge() = %v, want non-empty string", got)
			}
			t.Log("Generated challenge:", got)
		})
	}
}

func TestVerifyChallengeResponse(t *testing.T) {
	type args struct {
		challenge string
		response  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"[success] Valid challenge and response", args{"123123", SolveChallenge("123123")}, true},
		{"[success] Valid challenge and response", args{"221", SolveChallenge("221")}, true},
		{"[fail] Invalid challenge and response", args{"1234", "123123"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyChallengeResponse(tt.args.challenge, tt.args.response); got != tt.want {
				t.Errorf("VerifyChallengeResponse() = %v, want %v", got, tt.want)
			}
			t.Log("Challenge:", tt.args.challenge)
			t.Log("Response:", tt.args.response)
		})
	}
}
