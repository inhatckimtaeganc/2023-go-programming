package main

import (
	"fmt"
	"strings"
)

func main() {
	brokenWords := "I l?ve y?u"
	replaceWords := strings.NewReplacer("?", "o")
	fixedWords := replaceWords.Replace(brokenWords)
	fmt.Println(fixedWords)
}
