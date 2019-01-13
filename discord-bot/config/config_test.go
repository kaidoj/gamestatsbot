package config

import (
	"os"
	"path"
	"testing"
)

func TestInit(t *testing.T) {
	gp := os.Getenv("GOPATH")
	ap := path.Join(gp, "src/github.com/kaidoj/gamestatsbot/discord-bot")
	config := Init(ap)
	if config == nil {
		t.Errorf("Config not initialized")
	}
}
