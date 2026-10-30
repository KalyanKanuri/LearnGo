package main

import (
	"encoding/hex"
	"testing"
)

func TestGenSecretLength(t *testing.T) {
	tests := []struct {
		name     string
		algo     string
		wantByte int
	}{
		{name: "HS256", algo: "HS256", wantByte: 32},
		{name: "HS384", algo: "HS384", wantByte: 48},
		{name: "HS512", algo: "HS512", wantByte: 64},
		{name: "legacy HS-256", algo: "HS-256", wantByte: 32},
		{name: "legacy HS-384", algo: "HS-384", wantByte: 48},
		{name: "legacy HS-512", algo: "HS-512", wantByte: 64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenSecret(tt.algo)
			if err != nil {
				t.Fatalf("GenSecret returned error: %v", err)
			}
			decoded, err := hex.DecodeString(got)
			if err != nil {
				t.Fatalf("output is not valid hex: %v", err)
			}
			if len(decoded) != tt.wantByte {
				t.Errorf("got %d bytes, want %d", len(decoded), tt.wantByte)
			}
		})
	}
}

func TestGenSecretUnsupportedAlgo(t *testing.T) {
	for _, algo := range []string{"", "HS-999", "RS256", "hs256"} {
		t.Run(algo, func(t *testing.T) {
			got, err := GenSecret(algo)
			if err == nil {
				t.Errorf("expected error for algo %q, got secret %q", algo, got)
			}
			if got != "" {
				t.Errorf("expected empty secret on error, got %q", got)
			}
		})
	}
}

func TestGenSecretIsRandom(t *testing.T) {
	a, err := GenSecret("HS-256")
	if err != nil {
		t.Fatalf("GenSecret returned error: %v", err)
	}
	b, err := GenSecret("HS-256")
	if err != nil {
		t.Fatalf("GenSecret returned error: %v", err)
	}
	if a == b {
		t.Errorf("two generated secrets are identical: %s", a)
	}
}

func TestGenSecretNotEmpty(t *testing.T) {
	got, err := GenSecret("HS-256")
	if err != nil {
		t.Fatalf("GenSecret returned error: %v", err)
	}
	if got == "" {
		t.Error("generated secret is empty")
	}
}

