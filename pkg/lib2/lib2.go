package lib2

// add a nice header function using tui library
import (
	"github.com/charmbracelet/lipgloss"
)

func Farewell() string {
	return "Good bye from lib2"
}

// descriptor for an ASCII art function.
func AsciiArt() string {
	return `
   _____
  /     \
 | () () |
  \  ^  /
   |||||
   |||||
`
}

// a call that uses both functions.
func FarewellWithArt() string {
	return Farewell() + "\n" + AsciiArt()
}

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("205")).
	Bold(true)

func Header() string {
	return headerStyle.Render("Welcome to lib2!")
}

// add Header to FarewellWithArt
func FarewellWithHeaderAndArt() string {
	return Header() + "\n" + Farewell() + "\n" + AsciiArt()
}
