package config

import (
	"fmt"

	"github.com/Diaphteiros/kw/pluginlib/pkg/debug"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/yaml"
)

type KindConfig struct {
	// Binary is the path to the kind binary or just a name (has to be on the PATH in the latter case).
	// Defaults to 'kind'.
	Binary string `json:"binary"`
}

func (c *KindConfig) String() string {
	if c == nil {
		return ""
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Sprintf("error marshaling config: %v", err)
	}
	return string(data)
}

func (gc *KindConfig) Default() error {
	if gc.Binary == "" {
		gc.Binary = "kind"
	}
	return nil
}

func (c *KindConfig) Validate() error {
	errs := field.ErrorList{}
	if c.Binary == "" {
		errs = append(errs, field.Required(field.NewPath("binary"), "kind binary path is required"))
	}
	return errs.ToAggregate()
}

func LoadFromBytes(data []byte) (*KindConfig, error) {
	cfg := &KindConfig{}
	if len(data) > 0 {
		err := yaml.Unmarshal(data, cfg)
		if err != nil {
			return nil, fmt.Errorf("error unmarshaling kw_kind config: %w", err)
		}
	} else {
		debug.Debug("No kw_kind config provided, using default values")
	}
	if err := cfg.Default(); err != nil {
		return nil, fmt.Errorf("error setting default values for kw_kind config: %w", err)
	}
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("error validating kw_kind config: %w", err)
	}
	return cfg, nil
}
