package main

import (
	"errors"
	"os"
)

var (
	ErrEnv        = errors.New("not found RABBIT_PATH  environment variable")
	ErrEnvChannel = errors.New("not found RABBIT_CHANNEL  environment variable")
)

func main() {
	rabbit_url := os.Getenv("RABBIT_PATH")
	rabbit_ch := os.Getenv("RABBIT_CHANNEL")
	if len(rabbit_url) == 0 {
		return ErrEnv
	}
	if len(rabbit_ch) == 0 {
		return ErrEnvChannel
	}

}
