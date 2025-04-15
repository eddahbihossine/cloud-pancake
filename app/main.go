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
	"app/client"
)	

func hamid(){
	fmt.Println(quote.Go());

}

func main(){
	var input string

	fmt.Print("Hey Developer Welcome to the provider Please Enter a number from 1 --- 4:\n ");
	fmt.Print("1 --> aws \n 2--> Gcp \n 3 --> OCI \n 4 --> Azure\n\r " );
	fmt.Println("--------------------------------------------------------")
	fmt.Scanln(&input)
	samir := client.thisis(1);
	fmt.Println(samir);

	
	var provider string = cloud.ChooseProvider(input);
	fmt.Println(provider);
	hamid();
}