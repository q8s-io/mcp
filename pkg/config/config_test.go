package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		name       string
		configFile string
		expectErr  string
	}{
		{
			"config_notexists",
			"configs/config-notexists.yaml",
			"open configs/config-notexists.yaml: no such file or directory",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := LoadConfig(test.configFile)
			assert.EqualError(t, err, test.expectErr)
		})
	}
}
