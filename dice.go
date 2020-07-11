/* * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * *
 * Derived from http://www.codeguru.com/vb/gen/vb_misc/algorithms/article.php/c13137__1/Fuzzy-Matching-Demo-in-Access.htm
 * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * * */

package gofuzzy

const singlePercent = "%"
const singlePound = "#"
const doublePercent = "&&"
const doublePound = "##"

// Dice Coefficient based on bigrams.
// A good value would be 0.33 or above, a value under 0.2 is not a good match, from 0.2 to 0.33 is iffy.
func DiceCoefficient(input string, comparedTo string) float64 {
	var ngrams = toBiGrams(input)
	var compareToNgrams = toBiGrams(comparedTo)
	return diceCoefficientWithNGrams(ngrams, compareToNgrams)
}

// Dice Coefficient used to compare nGrams arrays produced in advance.
func diceCoefficientWithNGrams(nGrams []string, compareToNGrams []string) float64 {
	matches := len(intersect(nGrams, compareToNGrams))
	if matches == 0 {
		return 0.0
	}
	totalBigrams := float64(len(nGrams) + len(compareToNGrams))
	return (2.0 * float64(matches)) / totalBigrams
}

func toBiGrams(input string) []string {
	// nLength == 2
	//   from Jackson, return %j ja ac ck ks so on n#
	//   from Main, return #m ma ai in n#
	input = singlePercent + input + singlePound
	return toNGrams(input, 2)
}

func toTriGrams(input string) []string {
	// nLength == 3
	//   from Jackson, return %%j %ja jac ack cks kso son on# n##
	//   from Main, return ##m #ma mai ain in# n##
	input = doublePercent + input + doublePound
	return toNGrams(input, 3)
}

func toNGrams(input string, nLength int) []string {
	runes := []rune(input)
	itemsCount := len(runes) - 1
	ngrams := make([]string, itemsCount)
	for i := 0; i < itemsCount; i++ {
		gram := string(runes[i : i+nLength])
		ngrams = append(ngrams, gram)
	}
	return ngrams
}

func intersect(a, b []string) (c []string) {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}
	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
