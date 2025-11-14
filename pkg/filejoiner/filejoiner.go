package filejoiner

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

type FileData struct {
	Name    string
	Content string
}

type TemplateData struct {
	Files []FileData
}

type FileJoiner struct {
	DestinationBasePath string
	TemplatePath        string
}

func NewFileJoiner(basePath string, templatePath string) *FileJoiner {
	return &FileJoiner{
		DestinationBasePath: basePath,
		TemplatePath:        templatePath,
	}
}

func (fj *FileJoiner) JoinFiles(inputPaths []string, outputFileName string) error {
	var files []FileData

	for _, path := range inputPaths {
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}
		files = append(files, FileData{
			Name:    filepath.Base(path),
			Content: string(content),
		})
	}

	tmpl, err := template.ParseFiles(fj.TemplatePath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	outputPath := filepath.Join(fj.DestinationBasePath, outputFileName)

	// Ensure destination directory exists (useful for callers that provide a temp dir or a new folder)
	if err := os.MkdirAll(fj.DestinationBasePath, 0755); err != nil {
		return fmt.Errorf("failed to create destination dir %s: %w", fj.DestinationBasePath, err)
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	if err := tmpl.Execute(outputFile, TemplateData{Files: files}); err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	return nil
}
