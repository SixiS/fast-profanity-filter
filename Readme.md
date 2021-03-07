# Fast Profanity Filter

[![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Fast Profanity Filter Using a Trie

## Usage

```
go get github.com/sixis/fast-profanity-filter
```

```golang
import "github.com/sixis/fast-profanity-filter"

p := NewProfanityFilterFromStrings([]string{"dog", "dogging", "words"})
result := p.ReplaceProfanities("Dogging 	 dog     lol... Whatever haha")
// => "******* 	 ****     lol... Whatever haha"

p := NewProfanityFilerFromCsvFile("./profanities.csv")
result := p.ReplaceProfanities("Dogging 	 dog     lol... Whatever haha")
// => "******* 	 ****     lol... Whatever haha"
```