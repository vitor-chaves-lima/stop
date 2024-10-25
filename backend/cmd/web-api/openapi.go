package main

import (
	"fmt"
	"os"
)

func loadOpenAPIVersionDocument(version string) (*string, error) {
	filePath := fmt.Sprintf("./docs/openapi-%s.yaml", version)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	openapiContent := string(content)

	return &openapiContent, nil
}

func loadOpenAPIDocuments() (map[string]*string, error) {
	openAPISpecifications := make(map[string]*string)

	for _, version := range APIVersions {
		openAPIContent, err := loadOpenAPIVersionDocument(version)
		if err != nil {
			return nil, fmt.Errorf("failed to load OpenAPI document for version %s: %w", version, err)
		}

		openAPISpecifications[version] = openAPIContent
	}

	return openAPISpecifications, nil
}
