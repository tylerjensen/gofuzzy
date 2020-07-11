package gofuzzy

import (
	"fmt"
	"testing"
)

type TestPair struct {
	A, B string
}

func TestDiceCoefficient(t *testing.T) {
	pairs := []TestPair{
		TestPair{A: "test", B: "w"},
		TestPair{A: "test", B: "W"},
		TestPair{A: "test", B: "w"},
		TestPair{A: "test", B: "W"},
		TestPair{A: "test", B: "w "},
		TestPair{A: "test", B: "W "},
		TestPair{A: "test", B: " w"},
		TestPair{A: "test", B: " W"},
		TestPair{A: "test", B: " w "},
		TestPair{A: "test", B: " W "},
		TestPair{A: "Jensn", B: "Adams"},
		TestPair{A: "Jensn", B: "Benson"},
		TestPair{A: "Jensn", B: "Geralds"},
		TestPair{A: "Jensn", B: "Johannson"},
		TestPair{A: "Jensn", B: "Johnson"},
		TestPair{A: "Jensn", B: "Jensen"},
		TestPair{A: "Jensn", B: "Jordon"},
		TestPair{A: "Jensn", B: "Madsen"},
		TestPair{A: "Jensn", B: "Stratford"},
		TestPair{A: "Jensn", B: "Wilkins"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "2689 East Milkin Ave."},
		TestPair{A: "2130 South Fort Union Blvd.", B: "85 Morrison"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "2350 North Main"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "567 West Center Street"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "2130 Fort Union Boulevard"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "2310 S. Ft. Union Blvd."},
		TestPair{A: "2130 South Fort Union Blvd.", B: "98 West Fort Union"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "Rural Route 2 Box 29"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "PO Box 3487"},
		TestPair{A: "2130 South Fort Union Blvd.", B: "3 Harvard Square"},
	}
	for _, pair := range pairs {
		d := DiceCoefficient(pair.A, pair.B)
		if d == 0.0 {
			t.Errorf("Dice failed for A=%s and B=%s with %f\n", pair.A, pair.B, d)
		} else {
			fmt.Printf("Dice for A=%s and B=%s with %f\n", pair.A, pair.B, d)
		}
	}
}

/*
[InlineData("test", "w")]
[InlineData("test", "W")]
[InlineData("test", "w ")]
[InlineData("test", "W ")]
[InlineData("test", " w")]
[InlineData("test", " W")]
[InlineData("test", " w ")]
[InlineData("test", " W ")]
[InlineData("Jensn", "Adams")]
[InlineData("Jensn", "Benson")]
[InlineData("Jensn", "Geralds")]
[InlineData("Jensn", "Johannson")]
[InlineData("Jensn", "Johnson")]
[InlineData("Jensn", "Jensen")]
[InlineData("Jensn", "Jordon")]
[InlineData("Jensn", "Madsen")]
[InlineData("Jensn", "Stratford")]
[InlineData("Jensn", "Wilkins")]
[InlineData("2130 South Fort Union Blvd.", "2689 East Milkin Ave.")]
[InlineData("2130 South Fort Union Blvd.", "85 Morrison")]
[InlineData("2130 South Fort Union Blvd.", "2350 North Main")]
[InlineData("2130 South Fort Union Blvd.", "567 West Center Street")]
[InlineData("2130 South Fort Union Blvd.", "2130 Fort Union Boulevard")]
[InlineData("2130 South Fort Union Blvd.", "2310 S. Ft. Union Blvd.")]
[InlineData("2130 South Fort Union Blvd.", "98 West Fort Union")]
[InlineData("2130 South Fort Union Blvd.", "Rural Route 2 Box 29")]
[InlineData("2130 South Fort Union Blvd.", "PO Box 3487")]
[InlineData("2130 South Fort Union Blvd.", "3 Harvard Square")]
public void DiceCoefficientTests(string input, string match)
{
	var result = input.DiceCoefficient(match);
	Assert.True(result >= 0.0);
	output.WriteLine($"DiceCoefficient of \"{match}\" against \"{input}\" was {result}.");
}
*/
