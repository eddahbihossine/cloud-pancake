// Generated on: 2025-04-15 19:48:41

// ╔═╗┬ ┬ ┌─┐┌┬┐┌┬┐┌─┐┬ ┬┌┐ ┬          
// ║╣ ├─┤ ├┤  ││ ││├─┤├─┤├┴┐│          
// ╚═╝┴ ┴o└─┘─┴┘─┴┘┴ ┴┴ ┴└─┘┴  o       
// ╦ ╦╔═╗╔═╗╔═╗╦ ╦  ╔═╗╔═╗╔╦╗╦╔╗╔╔═╗  ┬
// ╠═╣╠═╣╠═╝╠═╝╚╦╝  ║  ║ ║ ║║║║║║║ ╦  │
// ╩ ╩╩ ╩╩  ╩   ╩   ╚═╝╚═╝═╩╝╩╝╚╝╚═╝  o

package cloud

func ChooseProvider(option string) string {
	switch option {
	case "1", "4":
		return "aws"
	case "2":
		return "Gcp"
	case "3":
		return "OCI"
	default:
		return "unknown"
	}
}
