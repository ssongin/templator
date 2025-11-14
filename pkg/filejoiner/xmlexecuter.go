package filejoiner

import (
	"os"
	"path/filepath"
	"strings"

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

func expandSources(sources []string) ([]string, error) {
	var expanded []string
	for _, src := range sources {
		// If src is a directory, add all files (non-recursive)
		info, err := os.Stat(src)
		if err == nil && info.IsDir() {
			entries, err := os.ReadDir(src)
			if err != nil {
				return nil, err
			}
			for _, entry := range entries {
				if !entry.IsDir() {
					expanded = append(expanded, filepath.Join(src, entry.Name()))
				}
			}
			continue
		}
		// If src contains a glob pattern
		if strings.ContainsAny(src, "*?[") {
			matches, err := filepath.Glob(src)
			if err != nil {
				return nil, err
			}
			expanded = append(expanded, matches...)
			continue
		}
		// Otherwise, treat as a file
		expanded = append(expanded, src)
	}
	return expanded, nil
}

func ExecuteJoiners(Joiners []YAMLJoiner) {
	for _, j := range Joiners {
		fj := NewFileJoiner(j.Destination, j.Template)

		for _, join := range j.Joins {
			expandedSources, err := expandSources(join.Sources)
			core.CheckError("Source expansion failed", err, "sources", join.Sources)
			err = fj.JoinFiles(expandedSources, join.Generate)
			core.CheckError("File join failed", err, "output", join.Generate, "sources", expandedSources)
			core.GetLogger().Info("Generated temporary file by joining fragments", "target", filepath.Join(j.Destination, join.Generate))
		}
	}
}
