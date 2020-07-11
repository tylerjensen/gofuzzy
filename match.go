package gofuzzy

func FuzzyEquals(strA *string, strB *string, requiredProbabilityScore float64) bool {
	return FuzzyMatch(strA, strB) > requiredProbabilityScore
}

func FuzzyMatch(strA *string, strB *string) float64 {
	return 0.0
}
