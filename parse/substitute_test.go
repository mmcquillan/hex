package parse

import (
	"testing"
)

func setupTokens(input string) (tokens map[string]string) {
	tokens = make(map[string]string)
	tokens["hex.id"] = "12345abcde"
	tokens["hex.timestamp"] = "123456789"
	tokens["hex.input"] = input
	return tokens
}

func TestSubstituteInput(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "never eat soggy waffles"
	a := Substitute("${hex.input}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteEscapedToken(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "$${hex.input}"
	a := Substitute("$${hex.input}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputSingle(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "eat"
	a := Substitute("${hex.input.1}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputRange(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "eat soggy"
	a := Substitute("${hex.input.1:2}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputRemainder(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "eat soggy waffles"
	a := Substitute("${hex.input.1:*}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputSingleOut(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := ""
	a := Substitute("${hex.input.6}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputRangeOut(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := "eat soggy waffles"
	a := Substitute("${hex.input.1:6}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputRangeOut2(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := ""
	a := Substitute("${hex.input.6:8}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteInputRangeMixed(t *testing.T) {
	tokens := setupTokens("never eat soggy waffles")
	e := ""
	a := Substitute("${hex.input.6:1}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteJson(t *testing.T) {
	tokens := setupTokens("{\"Name\": \"Hex\", \"Birth\": {\"Location\": \"Indiana\", \"Date\": \"Oct 2015\"}}")
	e := "Hex"
	a := Substitute("${hex.input.json:Name}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}

func TestSubstituteJson2(t *testing.T) {
	tokens := setupTokens("{\"Name\": \"Hex\", \"Birth\": {\"Location\": \"Indiana\", \"Date\": \"Oct 2015\"}}")
	e := "Indiana"
	a := Substitute("${hex.input.json:Birth.Location}", tokens)
	if e != a {
		t.Fatalf("Expected %s, but got %s", e, a)
	}
}
