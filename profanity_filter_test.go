package profanityfilter

import (
	"testing"
)

// TestNewProfanityFilterFromStrings checks that the profanity filter intialises properly
// and makes sure the replace function works on it.
func TestNewProfanityFilterFromStrings(t *testing.T) {
	p := NewProfanityFilterFromStrings([]string{"fuck", "fucking", "shit"})
	// TODO - test that the swear words are in the Trie

	result := p.ReplaceProfanities("Fucking 	 fuck     lol... Whatever haha")
	correctResult := "******* 	 ****     lol... Whatever haha"
	if result != correctResult {
		t.Fatalf("Replace profanities failed. Expected %q but got %q", correctResult, result)
	}
}

// TestNewProfanityFilterFromStrings checks that the profanity filter intialises properly
// and makes sure the replace function works on it.
func TestNewProfanityFilerFromCsvFile(t *testing.T) {
	p := NewProfanityFilerFromCsvFile("./test_profanities.csv")
	// TODO - test that the swear words are in the Trie

	result := p.ReplaceProfanities("Fucking 	 fuck     lol... Whatever haha")
	correctResult := "******* 	 ****     lol... Whatever haha"
	if result != correctResult {
		t.Fatalf("Replace profanities failed. Expected %q but got %q", correctResult, result)
	}
}
