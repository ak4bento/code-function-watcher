package similarity

import "github.com/xrash/smetrics"

// IsSimilar returns a similarity score between 0.0 and 1.0
func CalculateSimilarity(a, b string) float64 {
	return smetrics.JaroWinkler(a, b, 0.7, 4)
}

