package lib2

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
