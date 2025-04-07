package exporter

import (
	"encoding/json"
	"os"

	"github.com/ak4bento/code-function-watcher/pkg/scanner"
)

// ExportToJSON menyimpan hasil scan ke file JSON
func ExportToJSON(functions []scanner.FunctionInfo, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(functions)
}

func LoadFromJSON(path string) ([]scanner.FunctionInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var functions []scanner.FunctionInfo
	err = json.NewDecoder(file).Decode(&functions)
	if err != nil {
		return nil, err
	}

	return functions, nil
}
