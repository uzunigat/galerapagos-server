package env

import (
	"testing"
)

type NestedConfig struct {
	Bar string `env:"BAR"`
	Baz string `env:"BAZ" default:"b4z"`
}

type Config struct {
	Opt       string `env:"OPT" optional:"true"`
	Foo       string `env:"FOO"`
	NestedCfg NestedConfig
}

// the test with the incomplete env file needs to be executed first
func TestErrorOnRequiredFieldNotFound(t *testing.T) {
	cfg := Config{}
	Load("./fixtures/.env.incomplete")
	err := Marshal(&cfg)

	if err == nil {
		t.Fatal("should have thrown error when required env var is not present")
	}
}

func TestLoadEnv(t *testing.T) {
	cfg := Config{}
	Load("./fixtures/.env.complete")
	Marshal(&cfg)

	if cfg.Foo != "foo" {
		t.Fatal("variable does not have the default value", cfg.Foo)
	}
}

func TestLoadNestedEnv(t *testing.T) {
	cfg := Config{}
	Load("./fixtures/.env.complete")
	Marshal(&cfg)

	if cfg.NestedCfg.Bar != "bar" {
		t.Fatal("nested object was not properly populated", cfg.NestedCfg.Bar)
	}
}

func TestLoadDefaultValue(t *testing.T) {
	cfg := Config{}
	Load("./fixtures/.env.complete")
	Marshal(&cfg)

	if cfg.NestedCfg.Baz != "b4z" {
		t.Fatal("default value was not loaded", cfg.NestedCfg.Baz)
	}
}

func TestOptionalSetToEmpty(t *testing.T) {
	cfg := Config{}
	Load("./fixtures/.env.complete")
	Marshal(&cfg)

	if cfg.Opt != "" {
		t.Fatal("default value was not loaded", cfg.Opt)
	}
}
