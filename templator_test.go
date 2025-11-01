package templator_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ssongin/templator"
	"gopkg.in/yaml.v3"
)

func TestProcessYAML(t *testing.T) {
	// create temp working dir for generated output and config
	tmpDir := t.TempDir()

	// prepare absolute paths to testdata files
	base := filepath.Join("testdata", "joiner")
	absTemplate, err := filepath.Abs(filepath.Join(base, "join.tmpl"))
	if err != nil {
		t.Fatalf("failed to resolve template path: %v", err)
	}
	absInput1, err := filepath.Abs(filepath.Join(base, "input1.js"))
	if err != nil {
		t.Fatalf("failed to resolve input1 path: %v", err)
	}
	absInput2, err := filepath.Abs(filepath.Join(base, "input2.js"))
	if err != nil {
		t.Fatalf("failed to resolve input2 path: %v", err)
	}

	// build a clean YAML config that matches the structs used by ProcessYAML
	cfg := map[string]interface{}{
		"joiners": []interface{}{
			map[string]interface{}{
				"template":    absTemplate,
				"destination": tmpDir,
				"join": []interface{}{
					map[string]interface{}{
						"generate": "output.js",
						"source":   []string{absInput1, absInput2},
					},
				},
			},
		},
	}

	b, err := yaml.Marshal(cfg)
	if err != nil {
		t.Fatalf("marshal yaml: %v", err)
	}

	cfgPath := filepath.Join(tmpDir, "joiners.yaml")
	if err := os.WriteFile(cfgPath, b, 0644); err != nil {
		t.Fatalf("failed to write temp config: %v", err)
	}

	// execute
	templator.ProcessYAML(cfgPath)

	// compare generated file with expected fixture
	gotPath := filepath.Join(tmpDir, "output.js")
	got, err := os.ReadFile(gotPath)
	if err != nil {
		t.Fatalf("failed to read generated file %s: %v", gotPath, err)
	}

	wantPath := filepath.Join("testdata", "joiner", "output.js")
	want, err := os.ReadFile(wantPath)
	if err != nil {
		t.Fatalf("failed to read expected file %s: %v", wantPath, err)
	}

	if strings.TrimSpace(string(got)) != strings.TrimSpace(string(want)) {
		t.Logf("EXPECTED:\n%s\n", string(want))
		t.Logf("GOT:\n%s\n", string(got))
		t.Fatalf("generated content does not match expected\n--- expected (%s)\n--- got (%s)\n", wantPath, gotPath)
	}
}
