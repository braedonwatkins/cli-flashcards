package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	tea.Model // FIXME: why dont the docs mention this?
	CardIndex int
	Revealed  bool
}

func initialModel() model {
	return model{
		CardIndex: 0,
		Revealed:  false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "z":
			m.Revealed = true
		case "enter":
			m.Revealed = false

			if m.CardIndex < len(testCards)-1 {
				m.CardIndex++
			} else {
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "This is the header\n\n"

	s += fmt.Sprintf("%s\n", testCards[m.CardIndex].Front)

	if m.Revealed == true {
		s += fmt.Sprintf("%s\n", testCards[m.CardIndex].Back)
	}

	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
