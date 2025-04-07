package compare

import (
	"encoding/json"
	"os"

	"github.com/ak4bento/code-function-watcher/pkg/scanner"
	"github.com/ak4bento/code-function-watcher/pkg/similarity"
)

type SimilarityResult struct {
	FuncA      scanner.FunctionInfo
	FuncB      scanner.FunctionInfo
	Similarity float64
}

// LoadFromJSON loads previously scanned functions
func LoadFromJSON(path string) ([]scanner.FunctionInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var functions []scanner.FunctionInfo
	err = json.NewDecoder(file).Decode(&functions)
	return functions, err
}

// Compare checks similarity between new and old functions
func Compare(oldFuncs, newFuncs []scanner.FunctionInfo, threshold int) []SimilarityResult {
	var results []SimilarityResult

	for _, oldFn := range oldFuncs {
		for _, newFn := range newFuncs {
			sim := similarity.CalculateSimilarity(oldFn.Name, newFn.Name)
			if sim*100 >= float64(threshold) {
				results = append(results, SimilarityResult{
					FuncA:      oldFn,
					FuncB:      newFn,
					Similarity: sim,
				})
			}
		}
	}

	return results
}

func LoadFromFile(path string) ([]scanner.FunctionInfo, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var functions []scanner.FunctionInfo
	err = json.NewDecoder(file).Decode(&functions)
	return functions, err
}
