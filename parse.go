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
)

type TokenType int

const (
	DeckMetaDelim TokenType = iota

	CardMetaDelim

	QuestionStart

	AnswerStart
)

var (
	DeckMetaDelimToken = Token{Type: DeckMetaDelim, Literal: "==="}
	CardMetaDelimToken = Token{Type: CardMetaDelim, Literal: "---"}
	QuestionStartToken = Token{Type: QuestionStart, Literal: "Q:"}
	AnswerStartToken   = Token{Type: AnswerStart, Literal: "A:"}
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

func parse() {
	file, err := os.Open("./card.abc")
	if err != nil {
		// FIXME: what's the space of options for handling errs (log.Fatalf, panic, etc)
		log.Fatalf("impossible to open file: %s", err)
	}

	defer file.Close() // FIXME: how's defer work

	scanner := bufio.NewScanner(file)

	// FIXME: reminder on looping in golang... how is this looping??
	for scanner.Scan() {
		//FIXME: what's the space of APIs on scanners...
		// is bufio the canonical way to do this in golang
		line := scanner.Text()

		fmt.Print("+ " + line)
		//FIXME: reminder on how to do input (scanf in golang)
		var input string
		fmt.Scanln(&input)
	}

	// FIXME: what sort of errors can be encountered here??
	if err := scanner.Err(); err != nil {
		log.Fatalf("scanner encountered an error: %s", err)
	}

}
