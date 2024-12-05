package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var loremWords = []string{
	"lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit",
	"sed", "do", "eiusmod", "tempor", "incididunt", "ut", "labore", "et", "dolore",
	"magna", "aliqua", "ut", "enim", "ad", "minim", "veniam", "quis", "nostrud",
	"exercitation", "ullamco", "laboris", "nisi", "ut", "aliquip", "ex", "ea", "commodo",
	"consequat",
}

func generateLoremIpsum(sentences int) string {
	var loremText []string

	for i := 0; i < sentences; i++ {
		// ランダムな単語数の文を作成
		wordCount := rand.Intn(15) + 5 // 5～20単語の文
		var sentence []string

		for j := 0; j < wordCount; j++ {
			word := loremWords[rand.Intn(len(loremWords))]
			sentence = append(sentence, word)
		}

		// 文の最初を大文字にし、末尾にピリオドを追加
		sentence[0] = strings.Title(sentence[0])
		loremText = append(loremText, strings.Join(sentence, " ")+".")
	}

	return strings.Join(loremText, " ")
}

func main() {
	// 3つの文で構成されたLorem Ipsumを生成
	lorem := generateLoremIpsum(3)
	fmt.Println("---")
	fmt.Println(lorem)
}
