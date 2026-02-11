/* Holy fuck I don't remember how to write a parser

- Since we don't need to tokenize we can just eat lines right?
- So, parse a line, look for deck metadata, look for card metadata, look for a question, then look for an answer
	- If you hit EOF then include last parsed card, anything after is a warning message, right?
	- If you hit unexpected token then include last fully parsed card, anything after is a warning message
	- If you hit either EOF/unexpected token with no cards parsed then don't render and send error message
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type TokenType int

const (
	DeckMetaDelim TokenType = iota

	CardMetaDelim

	QuestionStart

	AnswerStart
)

var tokenTable = []Token{
	{Type: DeckMetaDelim, Literal: "==="},
	{Type: CardMetaDelim, Literal: "---"},
	{Type: QuestionStart, Literal: "Q:"},
	{Type: AnswerStart, Literal: "A:"},
}

var tokenByLiteral = func() map[string]Token {
	m := make(map[string]Token, len(tokenTable))

	for _, t := range tokenTable {
		m[t.Literal] = t
	}

	return m
}()

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func parse() {
	file, err := os.Open("./card.abc")
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Print("+ " + line)
		var input string
		fmt.Scanln(&input)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error: %s", err)
	}

}
