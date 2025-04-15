pipeline {
    agent any

    environment {
        TF_IN_AUTOMATION = 'true'
        APP_NAME = 'my-go-app'
    }

    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/eddahbihossine/cloud-pancake', branch: 'main'
            }
        }

        stage('Go Build') {
            steps {
                dir('app') {
                    sh 'go mod tidy'
                    sh 'go build -o bin/${APP_NAME}'
                }
            }
        }

        stage('Go Test') {
            steps {
                dir('app') {
                    sh 'go test ./...'
                }
            }
        }

        stage('Terraform Init') {
            steps {
                dir('infra') {
                    sh 'terraform init'
                }
            }
        }

        stage('Terraform Plan') {
            steps {
                dir('infra') {
                    sh 'terraform plan -out=tfplan'
                }
            }
        }

        stage('Terraform Apply') {
            steps {
                input "Approve apply?"
                dir('infra') {
                    sh 'terraform apply tfplan'
                }
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
    }
}
