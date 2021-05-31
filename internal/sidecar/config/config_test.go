package config

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig(`C:\Users\guyan\go\src\dtstack.com\dtstack\easyagent\prod\sidecar\example-config.yml`)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("cfg: %#v", cfg)

	b, err := yaml.Marshal(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("[%s]", b)
}

func TestParseAgents(t *testing.T) {
	agents, err := ParseAgents(`C:\Users\guyan\go\src\dtstack.com\dtstack\easyagent\prod\sidecar\agents-file.yml`)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cfg: %#v", agents)
}

func TestWriteAgents(t *testing.T) {
	agents, err := ParseAgents(`C:\Users\guyan\go\src\dtstack.com\dtstack\easyagent\prod\sidecar\agents-file.yml`)
	if err != nil {
		t.Fatal(err)
	}
	err = WriteAgents(agents, `C:\Users\guyan\go\src\dtstack.com\dtstack\easyagent\prod\sidecar\agents-file.yml`)
	if err != nil {
		t.Fatal(err)
	}
}
