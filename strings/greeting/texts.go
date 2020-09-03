//Package greeting Nested Package
package greeting

// Exported
const (
	WelcomeText = "Hola, este es el mundo Golang"
	MorningText = "Buenos dias"
	EveningText = "Buenas tardes"
)

// Not exported (only visible inside the `greeting` package)
var loremIpsumText = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad 
minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea 
commodo consequat.`
