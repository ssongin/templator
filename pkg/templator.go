package templator

import (
	"os"

	"github.com/ssongin/core"
	"github.com/ssongin/templator/pkg/filejoiner"
	"gopkg.in/yaml.v3"
)

type configuration struct {
	Joiner []filejoiner.YAMLJoiner `yaml:"joiners"`
}

func ProcessYAML(configPath string) {
	data, err := os.ReadFile(configPath)
	core.CheckError("Failed to read YAML file", err, "file", configPath)

	var root configuration
	err = yaml.Unmarshal(data, &root)
	core.CheckError("Failed to parse YAML", err, "file", configPath)

	filejoiner.ExecuteJoiners(root.Joiner)
}
