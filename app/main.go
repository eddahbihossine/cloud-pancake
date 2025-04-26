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
	"bufio"
	"os"
	"github.com/fatih/color"
	// "app/archi"
	// "app/client"
	)	

func hamid(){
	fmt.Println(quote.Go());

}

func main() {
	stdin := bufio.NewReader(os.Stdin)
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
	color.Yellow("Hello Cloud Engineer Welcome to the Cloud Pancake ☁️\n")
	color.Yellow("This command line tool is designed to help you set up the infrastucture as code and also CI/CD pipelines using jenkins . you might find it silly i know but wait untill it gets pro")
	color.Cyan("to proceed \n (y) for yes || (n) for no\n\n")

	var res string 
	fmt.Fscan(stdin, &res)
	
	if(res != "y"){
		color.Red("Thanks for using the tool anyway")
		return ;
	}

		color.Magenta("Please Enter a number from 1 --- 4 to choose your provider:")
		color.Green("1 --> AWS")
		color.Cyan("2 --> GCP")
		color.HiMagenta("3 --> OCI")
		color.Red("4 --> Azure")
		color.White("--------------------------------------------------------")
	
		var input string
		fmt.Scanln(&input)
		// hamid := archi.hello();
		provider := cloud.ChooseProvider(input)
		color.Yellow(provider)
	
		cloud.CreateIaC(provider)
		hamid()
}