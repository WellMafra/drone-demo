package config

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	envPort           string = "8282"
	envAllowedOrigins string = "allowed_origins, allowed_origins2"
)

func TestMain(m *testing.M) {
	before()
	exitResult := m.Run()
	after()
	os.Exit(exitResult)
}

func before() {
	os.Setenv("PORT", envPort)
	os.Setenv("ALLOWED_ORIGINS", envAllowedOrigins)
}

func after() {
	os.Unsetenv("PORT")
	os.Unsetenv("ALLOWED_ORIGINS")
}

func TestReadEnv(t *testing.T) {
	cfg := MustGet()
	assert.Equal(t, envPort, cfg.Port)
	assert.Equal(t, envAllowedOrigins, strings.Join(cfg.AllowedOrigins, ","))
}
