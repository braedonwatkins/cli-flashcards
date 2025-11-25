package main

import (
	"image/color"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

// RENDER - basically treating this like JSX

// // global state

type model struct {
	tea.Model
	CardIndex int
	Revealed  bool
}

func initialModel() model {
	return model{
		CardIndex: 0,
		Revealed:  false,
	}
}

// // render functions

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
			m.Revealed = !m.Revealed // OPINION: Reveal should be a toggle cus it sucks to accidentally press it
		case "enter":
			m.Revealed = false

			if m.CardIndex < len(testCards)-1 {
				m.CardIndex++
			} else {
				return m, tea.Quit
			}
		case "backspace":
			m.Revealed = false

			if m.CardIndex > 0 {
				m.CardIndex--
			}
		}

	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	physicalWidth, _, _ := term.GetSize(uintptr((os.Stdout.Fd())))
	docStyle := lipgloss.NewStyle().Padding(1, 2, 1, 2)
	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}
	cardWidth := physicalWidth / 3

	// Dialog Scope
	{
		question := lipgloss.NewStyle().Width(cardWidth).Align(lipgloss.Center).Render(rainbow(lipgloss.NewStyle(), testCards[m.CardIndex].Front, blends))

		answer := ""
		if m.Revealed == true {
			answer = lipgloss.NewStyle().Width(cardWidth).Align(lipgloss.Center).Render(testCards[m.CardIndex].Back)
		}

		ui := lipgloss.JoinVertical(lipgloss.Center, question, answer)

		// OPINION: IDC about styling... may be nice to make it configurable
		dialog := lipgloss.Place(physicalWidth, 9,
			lipgloss.Center, lipgloss.Center,
			dialogBoxStyle.Render(ui),
			lipgloss.WithWhitespaceChars("猫咪"),
			lipgloss.WithWhitespaceForeground(subtle),
		)

		s.WriteString(dialog + "\n\n")
	}

	return s.String()
}

// STYLE - basically treating this like CSS

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

func rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		color, _ := colorful.MakeColor(colors[i%len(colors)])
		str = str + base.Foreground(lipgloss.Color(color.Hex())).Render(string(ss))
	}
	return str
}
