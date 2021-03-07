package profanityfilter

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/dghubble/trie"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// ProfanityFilter contains the Trie used for profianity filtering
type ProfanityFilter struct {
	profanityTrie *trie.RuneTrie
}

// NewProfanityFilterFromStrings takes in an array of strings and returns a
// ProfanityFilter struct with a fully constructed Trie containing the profanities.
func NewProfanityFilterFromStrings(profanities []string) *ProfanityFilter {
	profanityTrie := trie.NewRuneTrie()

	for _, word := range profanities {
		word, _ = normalizeText(word)
		profanityTrie.Put(word, 1)
	}

	return &ProfanityFilter{profanityTrie: profanityTrie}
}

// NewProfanityFilerFromCsvFile takes in a path to a csv file and returns a
// ProfanityFilter struct with a fully constructed Trie containing the profanities.
func NewProfanityFilerFromCsvFile(filePath string) *ProfanityFilter {
	profanityTrie := trie.NewRuneTrie()

	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)
	r.FieldsPerRecord = -1

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		for _, word := range record {
			word, _ = normalizeText(word)
			profanityTrie.Put(word, 1)
		}
	}

	return &ProfanityFilter{profanityTrie: profanityTrie}
}

// ReplaceProfanities replaces any profanities detected in the input string with * runes
func (p *ProfanityFilter) ReplaceProfanities(input string) string {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			for j := i + 1; j < len(runes); j++ {
				if !unicode.IsLetter(runes[j]) {
					runes = p.replaceProfanity(runes, i, j)
					i = j
					break
				} else if j == len(runes)-1 {
					runes = p.replaceProfanity(runes, i, j+1)
					i = j
					break
				}
			}
		}
	}
	return string(runes)
}

func (p *ProfanityFilter) replaceProfanity(runes []rune, a int, b int) []rune {
	word, _ := normalizeText(string(runes[a:b]))
	if p.profanityTrie.Get(word) != nil {
		for i := a; i <= b-1; i++ {
			runes[i] = '*'
		}
	}
	return runes
}

func normalizeText(input string) (string, error) {
	// Lower case
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	// Normalize unicode characters
	bytes := make([]byte, len(input))
	normalize := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}), norm.NFC)
	_, _, err := normalize.Transform(bytes, []byte(input), true)
	if err != nil {
		return "", err
	}
	input = string(bytes)

	//Get rid of zero-width spaces
	input = strings.Replace(input, "\u200b", "", -1)

	return input, nil
}
