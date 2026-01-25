package envdbar

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	err := Load()

	if err != nil {
		t.Fatal(err)
	}

	if got := os.Getenv("PORT"); got != "3000" {
		t.Errorf("PORT: expected %q, got %q", "3000", got)
	}

	if got := os.Getenv("SERVER"); got != "0.0.0.0" {
		t.Errorf("SERVER: expected %q, got %q", "0.0.0.0", got)
	}

	if got := os.Getenv("DATABASE_NAME"); got != "database_name" {
		t.Errorf("DATABASE_NAME: expected %q, got %q", "database_name", got)
	}

	if got := os.Getenv("DATABASE_USERNAME"); got != "username" {
		t.Errorf("DATABASE_USERNAME: expected %q, got %q", "username", got)
	}

	if got := os.Getenv("TOKEN"); got != "mk23m123=sdd" {
		t.Errorf("TOKEN: expected %q, got %q", "mk23m123=sdd", got)
	}

	if got := os.Getenv("INDENTED_VAR"); got != "indented_value" {
		t.Errorf("INDENTED_VAR: expected %q, got %q", "indented_value", got)
	}

	if got := os.Getenv("QUOTED_VALUE"); got != "hello world" {
		t.Errorf("QUOTED_VALUE: expected %q, got %q", "hello world", got)
	}

	if got := os.Getenv("SINGLE_QUOTED"); got != "single quoted" {
		t.Errorf("SINGLE_QUOTED: expected %q, got %q", "single quoted", got)
	}

	if got := os.Getenv("INLINE_COMMENT"); got != "value" {
		t.Errorf("INLINE_COMMENT: expected %q, got %q", "value", got)
	}

	if got := os.Getenv("QUOTED_WITH_HASH"); got != "value#with#hash" {
		t.Errorf("QUOTED_WITH_HASH: expected %q, got %q", "value#with#hash", got)
	}
}

func TestGet(t *testing.T) {
	_ = Load()

	if got := Get("PORT"); got != "3000" {
		t.Errorf("PORT: expected %q, got %q", "3000", got)
	}

	if got := Get("DATABASE_PASSWORD"); got != "" {
		t.Errorf("DATABASE_PASSWORD: expected empty string for undefined variable, got %q", got)
	}

	if got := Get("DATABASE_PASSWORD", "default_pass"); got != "default_pass" {
		t.Errorf("DATABASE_PASSWORD with default: expected %q, got %q", "default_pass", got)
	}

	if got := Get("PORT", "8080"); got != "3000" {
		t.Errorf("PORT with default: expected %q (existing value), got %q", "3000", got)
	}
}

func TestLoadCustomFile(t *testing.T) {
	err := Load(".env.test")
	if err != nil {
		t.Fatal(err)
	}

	if got := Get("TEST_API_KEY"); got != "abc123xyz" {
		t.Errorf("TEST_API_KEY: expected %q, got %q", "abc123xyz", got)
	}

	if got := Get("TEST_DEBUG"); got != "true" {
		t.Errorf("TEST_DEBUG: expected %q, got %q", "true", got)
	}

	if got := Get("TEST_MAX_CONNECTIONS"); got != "100" {
		t.Errorf("TEST_MAX_CONNECTIONS: expected %q, got %q", "100", got)
	}
}

func TestParseValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple value", "hello", "hello"},
		{"value with spaces", "  hello  ", "hello"},
		{"double quoted", "\"hello world\"", "hello world"},
		{"single quoted", "'hello world'", "hello world"},
		{"inline comment", "hello # this is a comment", "hello"},
		{"quoted with hash", "\"hello#world\"", "hello#world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseValue(tt.input)
			if got != tt.expected {
				t.Errorf("parseValue(%q): expected %q, got %q", tt.input, tt.expected, got)
			}
		})
	}
}
