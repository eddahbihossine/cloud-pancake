// Generated on: 2025-04-15 19:48:41

// ╔═╗┬ ┬ ┌─┐┌┬┐┌┬┐┌─┐┬ ┬┌┐ ┬          
// ║╣ ├─┤ ├┤  ││ ││├─┤├─┤├┴┐│          
// ╚═╝┴ ┴o└─┘─┴┘─┴┘┴ ┴┴ ┴└─┘┴  o       
// ╦ ╦╔═╗╔═╗╔═╗╦ ╦  ╔═╗╔═╗╔╦╗╦╔╗╔╔═╗  ┬
// ╠═╣╠═╣╠═╝╠═╝╚╦╝  ║  ║ ║ ║║║║║║║ ╦  │
// ╩ ╩╩ ╩╩  ╩   ╩   ╚═╝╚═╝═╩╝╩╝╚╝╚═╝  o

package cloud

import(
	"fmt"
	"os"
)

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


func CreateIaC(provider string) string {

	if (provider == "unknown"){
		fmt.Println("Could not be able to create a TF for you sorry !");
		return "www" ;
	}
	dir := fmt.Sprintf("../iac/%s", provider)

	// Create directory if it doesn't exist
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return "❌ Failed to create directory"
	}

	// Create Jenkinsfile content
	jenkins := fmt.Sprintf(`pipeline {
    agent any
    stages {
        stage('Deploy %s') {
            steps {
                echo 'Deploying to %s'
            }
        }
    }
}`, provider, provider)

	// Write Jenkinsfile
	err = os.WriteFile(fmt.Sprintf("%s/Jenkinsfile", dir), []byte(jenkins), 0644)
	if err != nil {
		return "❌ Failed to create Jenkinsfile"
	}

	// Create main.tf content for Terraform
	mainTF := fmt.Sprintf(`# Terraform configuration for %s

provider "%s" {
  # Example configuration for %s provider
  region = "us-east-1"
}

resource "%s_instance" "example" {
  ami           = "ami-123456"
  instance_type = "t2.micro"
}

output "%s_instance_ip" {
  value = "%s_instance.example.public_ip"
}`, provider, provider, provider, provider, provider, provider)

	// Write main.tf file
	err = os.WriteFile(fmt.Sprintf("%s/main.tf", dir), []byte(mainTF), 0644)
	if err != nil {
		return "❌ Failed to create main.tf"
	}

	// Return success message
	return fmt.Sprintf("✅ %s IaC folder, Jenkinsfile, and main.tf created!", provider)
}
