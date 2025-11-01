package filejoiner

import (
	"path/filepath"

	"github.com/ssongin/core"
)

type Joiners struct {
	Joiner []YAMLJoiner `yaml:"joiners"`
}

type YAMLJoiner struct {
	Template    string     `yaml:"template"`
	Destination string     `yaml:"destination"`
	Joins       []YAMLJoin `yaml:"join"`
}

type YAMLJoin struct {
	Generate string   `yaml:"generate"`
	Sources  []string `yaml:"source"`
}

func ExecuteJoiners(Joiners []YAMLJoiner) {
	for _, j := range Joiners {
		fj := NewFileJoiner(j.Destination, j.Template)

		for _, join := range j.Joins {
			err := fj.JoinFiles(join.Sources, join.Generate)
			core.CheckError("File join failed", err, "output", join.Generate, "sources", join.Sources)
			core.GetLogger().Info("Generated temporary file by joining fragments", "target", filepath.Join(j.Destination, join.Generate))
		}
	}
}
