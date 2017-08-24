// Here several spinners are demonstrated. Try adding your own []string patterns.,,
package main

import (
	"fmt"
	"time"

	"github.com/mdwhatcott/spin"
)

func main() {
	ShowDefault()
	ShowStyles()
	fmt.Println("\nDone.")
}

func ShowDefault() {
	fmt.Println("Default spinner: ")
	spin.GoStart()
	time.Sleep(time.Second * 1)
	spin.Stop()
}

func ShowStyles() {
	styles := map[string]string{
		"spin.StylePops        ": spin.StylePops,
		"spin.StyleBrackets    ": spin.StyleBrackets,
		"spin.StyleLine        ": spin.StyleLine,
		"spin.StyleSteps       ": spin.StyleSteps,
		"spin.StyleShutter     ": spin.StyleShutter,
		"spin.StyleNumbers     ": spin.StyleNumbers,
		"spin.StyleAphabet     ": spin.StyleAlphabet,
	}
	for title, style := range styles {
		Show(title, style)
	}
}

func Show(title string, style string) {
	fmt.Println("\n")
	spinner := spin.NewWithPadding(style, time.Millisecond*100, title, "    "+fmt.Sprint(style))
	go spinner.Start()
	time.Sleep(time.Second * 1)
	spinner.Stop()
}
