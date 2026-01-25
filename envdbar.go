// Package envdbar provides a simple way to load environment variables
// from .env files into the application.
package envdbar

import (
	"bufio"
	"os"
	"strings"
)

// Load reads environment variables from a file and sets them in the process.
// If no filename is provided, it defaults to ".env".
// It supports comments (#), quoted values, inline comments, and values containing "=".
func Load(filename ...string) error {
	file := ".env"
	if len(filename) > 0 {
		file = filename[0]
	}

	envfile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer envfile.Close()

	scanner := bufio.NewScanner(envfile)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if len(text) == 0 {
			continue
		}

		if strings.HasPrefix(text, "#") {
			continue
		}

		if !strings.Contains(text, "=") {
			continue
		}

		keyValue := strings.SplitN(text, "=", 2)
		key := strings.TrimSpace(keyValue[0])
		value := parseValue(keyValue[1])

		if err := os.Setenv(key, value); err != nil {
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseValue(raw string) string {
	value := strings.TrimSpace(raw)

	// Remove inline comments (only if not inside quotes)
	if !strings.HasPrefix(value, "\"") && !strings.HasPrefix(value, "'") {
		if idx := strings.Index(value, " #"); idx != -1 {
			value = strings.TrimSpace(value[:idx])
		}
	}

	// Remove surrounding quotes
	if len(value) >= 2 {
		if (strings.HasPrefix(value, "\"") && strings.HasSuffix(value, "\"")) ||
			(strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'")) {
			value = value[1 : len(value)-1]
		}
	}

	return value
}

// Get retrieves an environment variable by name.
// If the variable is not set or is empty, it returns the optional defaultValue.
func Get(variable string, defaultValue ...string) string {
	value := os.Getenv(variable)
	if value == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return value
}
