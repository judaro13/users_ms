package main

import (
	"os"
	"testing"
)

func TestValidEnvVars(t *testing.T) {
	os.Setenv("RABBIT_PATH", "path")
	os.Setenv("RABBIT_CHANNEL", "ch")
	os.Setenv("DB_HOST", "loclahost")
	os.Setenv("DB_USER", "test")
	os.Setenv("DB_PASSWORD", "test")
	os.Setenv("DB_NAME", "test")

	ValidEnvVars()
}
