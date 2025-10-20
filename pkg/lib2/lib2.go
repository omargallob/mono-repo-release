package lib2

func Farewell() string {
	return "Good bye from lib2"
}

// add a simple function to be tested along the lines of farwell, maybe an ascii art
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

// a call that uses both functions
func FarewellWithArt() string {
	return Farewell() + "\n" + AsciiArt()
}
