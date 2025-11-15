package main

import (
	"fmt"
	"image/color"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

var (
	dialogBoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	blends = gamut.Blends(lipgloss.Color("#F25D94"), lipgloss.Color("#EDFF82"), 50)
)

const (
	// In real life situations we'd adjust the document to fit the width we've
	// detected. In the case of this example we're hardcoding the width, and
	// later using the detected width only to truncate in order to avoid jaggy
	// wrapping.
	width = 96

	columnWidth = 30
)

func rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Foreground(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}

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
	s := strings.Builder{}

	// Dialog Scope
	{
		question := lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(rainbow(lipgloss.NewStyle(), testCards[m.CardIndex].Front, blends))

		answer := ""
		if m.Revealed == true {
			answer = lipgloss.NewStyle().Width(50).Align(lipgloss.Center).Render(testCards[m.CardIndex].Back)
		}

		ui := lipgloss.JoinVertical(lipgloss.Center, question, answer)

		dialog := lipgloss.Place(width, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		s.WriteString(dialog + "\n\n")
	}

	return s.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
