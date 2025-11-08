package main

type flashCard struct {
	Front string
	Back  string
}

var testCards = []flashCard{
	{Front: "What is the first letter of the alphabet", Back: "A"},
	{Front: "What is the last letter of the alphabet", Back: "Z"},
}
