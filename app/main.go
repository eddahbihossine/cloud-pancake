// Generated on: 2025-04-15 19:48:28

// ╔═╗┬ ┬ ┌─┐┌┬┐┌┬┐┌─┐┬ ┬┌┐ ┬          
// ║╣ ├─┤ ├┤  ││ ││├─┤├─┤├┴┐│          
// ╚═╝┴ ┴o└─┘─┴┘─┴┘┴ ┴┴ ┴└─┘┴  o       
// ╦ ╦╔═╗╔═╗╔═╗╦ ╦  ╔═╗╔═╗╔╦╗╦╔╗╔╔═╗  ┬
// ╠═╣╠═╣╠═╝╠═╝╚╦╝  ║  ║ ║ ║║║║║║║ ╦  │
// ╩ ╩╩ ╩╩  ╩   ╩   ╚═╝╚═╝═╩╝╩╝╚╝╚═╝  o

package main

import (
	"fmt"
	"rsc.io/quote"
	"app/cloud"
	"github.com/fatih/color"
	// "app/client"
)	

func hamid(){
	fmt.Println(quote.Go());

}

func main() {
	// Gopher ASCII Art
	gopher := `
           ,_---~~~~~----._
  _,,_,*^____      _____''*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \'/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           |
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
`

	color.Blue(gopher)
	color.Yellow("Hey Developer 👋 Welcome to the Cloud Pancake ☁️\n")
	color.Magenta("Please Enter a number from 1 --- 4 to choose your provider:")
	color.Green("1 --> AWS")
	color.Cyan("2 --> GCP")
	color.HiMagenta("3 --> OCI")
	color.Red("4 --> Azure")
	color.White("--------------------------------------------------------")

	var input string
	fmt.Scanln(&input)

	provider := cloud.ChooseProvider(input)
	color.Yellow(provider)

	cloud.CreateIaC(provider)
	hamid()
}