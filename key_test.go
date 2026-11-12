package main

import (
	"encoding/hex"
	"strings"
	"testing"
)

func TestGenerateAPIKey(t *testing.T) {
	key, err := generateAPIKey()
	if err != nil {
		t.Fatalf("generateAPIKey() returned an error: %v", err)
	}

	const prefix = "tm_key_"
	if !strings.HasPrefix(key, prefix) {
		t.Fatalf("generateAPIKey() = %q, want prefix %q", key, prefix)
	}

	encoded := strings.TrimPrefix(key, prefix)
	if len(encoded) != 64 {
		t.Fatalf("generated key has %d hex characters, want 64", len(encoded))
	}
	if _, err := hex.DecodeString(encoded); err != nil {
		t.Fatalf("generated key suffix is not valid hexadecimal: %v", err)
	}
}

func TestHashAPIKey(t *testing.T) {
	const key = "tm_key_test"

	first := hashAPIKey(key)
	if len(first) != 64 {
		t.Fatalf("hashAPIKey() returned %d characters, want 64", len(first))
	}
	if first != hashAPIKey(key) {
		t.Fatal("hashAPIKey() returned different hashes for the same key")
	}
	if first == key {
		t.Fatal("hashAPIKey() returned the original key")
	}
}
