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
	case "1":
		return "aws"
	case "2":
		return "Gcp"
	case "3":
		return "OCI"
	case "4" :
		return "azure"
	default:
		return "unknown"
	}
}


func CreateIaC(provider string) string {
	infra := "None"
	if (provider == "unknown"){
		fmt.Println("Could not be able to create a TF for you sorry !");
		return "www" ;
	}
	if(provider == "Gcp"){
		infra = `
		[Global Infrastructure]  
		└── Regions (e.g., us-central1, europe-west4)  
		  ├── Zones (e.g., us-central1-a, europe-west4-b)  
		  │  ├── Compute  
		  │  │  ├── Compute Engine VMs  
		  │  │  │  ├── General-Purpose  
		  │  │  │  ├── Memory-Optimized  
		  │  │  │  ├── GPU Instances (A2, T4)  
		  │  │  │  └── Containers (GKE, Cloud Run)  
		  │  │  ├── Storage  
		  │  │  │  ├── Block Storage (Persistent Disk)  
		  │  │  │  ├── Object Storage (Cloud Storage)  
		  │  │  │  ├── File Storage (Filestore)  
		  │  │  │  └── Archive Storage (Coldline)  
		  │  │  ├── Networking  
		  │  │  │  ├── Virtual Private Cloud (VPC)  
		  │  │  │  │  ├── Subnets  
		  │  │  │  │  ├── Firewall Rules  
		  │  │  │  │  └── Routes  
		  │  │  │  ├── Load Balancing (Global, Regional)  
		  │  │  │  ├── Cloud Interconnect (Hybrid Networking)  
		  │  │  │  └── Cloud CDN  
		  │  │  └── Databases  
		  │  │    ├── Cloud SQL (MySQL, PostgreSQL)  
		  │  │    ├── Firestore (NoSQL)  
		  │  │    ├── Bigtable  
		  │  │    └── Spanner (Global Scale)  
		  │  └── Zone Redundancy (Cross-Zone Replication)  
		  ├── Global Services  
		  │  ├── Cloud Storage (Multi-Regional)  
		  │  ├── Cloud CDN  
		  │  └── Global Load Balancer  
		  └── Hybrid Connectivity  
		    ├── Cloud VPN (Site-to-Site/Client)  
		    └── Dedicated Interconnect  
		
		[Organization & Tenancy]  
		├── Organizations  
		│  ├── Folders (Resource Hierarchy)  
		│  └── Projects (Isolated Resources)  
		├── Identity & Access Management (IAM)  
		│  ├── Service Accounts  
		│  ├── Roles (Predefined/Custom)  
		│  └── Resource Manager Policies  
		└── Labels (Resource Tagging)  
		
		[Security & Compliance]  
		├── Encryption  
		│  ├── Cloud KMS  
		│  └── Cloud HSM  
		├── Threat Detection  
		│  ├── Security Command Center  
		│  ├── Web Risk API  
		│  └── Chronicle (SIEM)  
		├── Network Security  
		│  ├── Cloud Armor (DDoS/WAF)  
		│  ├── VPC Service Controls  
		│  └── Identity-Aware Proxy (IAP)  
		└── Secrets Management  
		   ├── Secret Manager  
		   └── Binary Authorization  
		
		[Management & Observability]  
		├── Monitoring  
		│  ├── Cloud Monitoring (Metrics, Dashboards)  
		│  ├── Cloud Logging  
		│  └── Error Reporting  
		├── Infrastructure as Code (IaC)  
		│  ├── Deployment Manager  
		│  └── Terraform  
		├── Cost Management  
		│  ├── Billing Reports  
		│  └── Recommender (Optimization)  
		└── Analytics  
		   ├── BigQuery (Data Warehouse)  
		   └── Looker (BI)  
		
		[Developer Tools]  
		├── Serverless  
		│  ├── Cloud Functions  
		│  ├── Cloud Run (Containers)  
		│  └── API Gateway  
		├── CI/CD  
		│  ├── Cloud Build  
		│  ├── Artifact Registry  
		│  └── Cloud Deploy  
		└── SDKs & CLI  
		   ├── gcloud CLI  
		   └── Cloud Shell  
		
		[AI/ML & Advanced Services]  
		├── Machine Learning  
		│  ├── Vertex AI (End-to-End ML)  
		│  ├── Vision AI  
		│  └── Natural Language API  
		├── High-Performance Compute  
		│  ├── TPUs (TensorFlow)  
		│  └── Batch  
		└── IoT & Edge  
		   ├── IoT Core  
		   └── Anthos (Hybrid/Multi-Cloud)  
		
		[Specialized Workloads]  
		├── Healthcare (Healthcare API)  
		├── Quantum Computing (Cirq)  
		└── Gaming (Open Match)  
		
		`
	}
	if(provider == "azure"){
		infra = `
		[Global Infrastructure]  
		└── Regions (e.g., eastus, westeurope)  
		  ├── Availability Zones (AZ1, AZ2, AZ3)  
		  │  ├── Compute  
		  │  │  ├── Virtual Machines (VMs)  
		  │  │  │  ├── General-Purpose (B-Series)  
		  │  │  │  ├── GPU Instances (NCv3, NDv2)  
		  │  │  │  └── Containers (AKS, Container Instances)  
		  │  │  ├── Storage  
		  │  │  │  ├── Block Storage (Managed Disks)  
		  │  │  │  ├── Object Storage (Blob Storage)  
		  │  │  │  ├── File Storage (Azure Files)  
		  │  │  │  └── Archive Storage (Archive Blob)  
		  │  │  ├── Networking  
		  │  │  │  ├── Virtual Network (VNet)  
		  │  │  │  │  ├── Subnets  
		  │  │  │  │  ├── Network Security Groups (NSGs)  
		  │  │  │  │  └── Route Tables  
		  │  │  │  ├── Load Balancers (Basic, Standard)  
		  │  │  │  ├── Application Gateway (WAF)  
		  │  │  │  └── ExpressRoute (Private Connectivity)  
		  │  │  └── Databases  
		  │  │    ├── Azure SQL Database  
		  │  │    ├── Cosmos DB (NoSQL)  
		  │  │    └── Database for PostgreSQL/MySQL  
		  │  └── Zone-Redundant Services (ZRS)  
		  ├── Global Services  
		  │  ├── Azure Traffic Manager  
		  │  ├── Azure Front Door (CDN/WAF)  
		  │  └── Cross-Region Replication  
		  └── Hybrid Connectivity  
		    ├── ExpressRoute (Dedicated Links)  
		    └── VPN Gateway (Site-to-Site/Point-to-Site)  

		[Organization & Tenancy]  
		├── Management Groups  
		│  ├── Subscriptions  
		│  └── Resource Groups (Logical Grouping)  
		├── Azure Active Directory (AAD)  
		│  ├── Users, Groups, Roles  
		│  ├── Enterprise Applications  
		│  └── Conditional Access  
		└── Tags (Resource Governance)  

		[Security & Compliance]  
		├── Encryption  
		│  ├── Azure Key Vault  
		│  └── Azure Disk Encryption  
		├── Threat Detection  
		│  ├── Microsoft Defender for Cloud  
		│  ├── Sentinel (SIEM)  
		│  └── Azure Firewall  
		├── Network Security  
		│  ├── DDoS Protection  
		│  ├── Network Security Groups (NSGs)  
		│  └── Web Application Firewall (WAF)  
		└── Access Control  
		   ├── Managed Identities  
		   └── Azure Bastion (SSH/RDP Gateway)  

		[Management & Observability]  
		├── Monitoring  
		│  ├── Azure Monitor (Metrics, Logs)  
		│  ├── Log Analytics  
		│  └── Application Insights  
		├── Infrastructure as Code (IaC)  
		│  ├── ARM Templates  
		│  └── Terraform  
		├── Cost Management  
		│  ├── Cost Analysis  
		│  ├── Budgets  
		│  └── Azure Advisor  
		└── Analytics  
		   ├── Synapse Analytics  
		   └── Power BI  

		[Developer Tools]  
		├── Serverless  
		│  ├── Azure Functions  
		│  ├── Logic Apps  
		│  └── API Management  
		├── CI/CD  
		│  ├── Azure DevOps  
		│  ├── GitHub Actions  
		│  └── Azure Pipelines  
		└── SDKs & CLI  
		   ├── Azure CLI  
		   └── Cloud Shell  

		[AI/ML & Advanced Services]  
		├── Machine Learning  
		│  ├── Azure Machine Learning Studio  
		│  ├── Cognitive Services (Vision, Speech)  
		│  └── Bot Service  
		├── High-Performance Compute  
		│  ├── HPC Virtual Machines  
		│  └── CycleCloud  
		└── IoT & Edge  
		   ├── IoT Hub  
		   └── Azure Arc (Hybrid Cloud)  

		[Specialized Workloads]  
		├── Gaming (PlayFab)  
		├── Healthcare (Azure Health Data Services)  
		└── Quantum Computing (Azure Quantum)  
		`

	}
	if(provider == "aws"){
		infra = `
		[AWS Global Infrastructure]  
				└── Regions (e.g., us-east-1, eu-central-1)  
				  ├── Availability Zones (AZs)  
				  │  ├── Data Centers  
				  │  │  ├── Compute  
				  │  │  │  ├── EC2 Instances  
				  │  │  │  │  ├── Bare Metal (EC2 Metal)  
				  │  │  │  │  ├── Virtual Machines (General Purpose, Memory Optimized, etc.)  
				  │  │  │  │  ├── GPU Instances (P3, P4, G5)  
				  │  │  │  │  └── Containers (ECS, EKS, Fargate)  
				  │  │  │  ├── Storage  
				  │  │  │  │  ├── Block Storage (EBS: gp3, io2)  
				  │  │  │  │  ├── Object Storage (S3)  
				  │  │  │  │  ├── File Storage (EFS, FSx)  
				  │  │  │  │  └── Archive Storage (Glacier)  
				  │  │  │  ├── Networking  
				  │  │  │  │  ├── Virtual Private Cloud (VPC)  
				  │  │  │  │  │  ├── Subnets (Public/Private)  
				  │  │  │  │  │  ├── Security Groups & NACLs  
				  │  │  │  │  │  └── Route Tables  
				  │  │  │  │  ├── Load Balancers (ALB, NLB, GWLB)  
				  │  │  │  │  ├── Elastic Network Adapters (ENA)  
				  │  │  │  │  └── Direct Connect (Private Network Links)  
				  │  │  │  └── Databases  
				  │  │  │    ├── Managed RDS (Aurora, PostgreSQL, MySQL)  
				  │  │  │    ├── DynamoDB (NoSQL)  
				  │  │  │    └── Specialty DBs (Redshift, DocumentDB)  
				  │  │  └── High-Speed AZ Interconnect (AWS Backbone)  
				  ├── Global Services  
				  │  ├── S3 Cross-Region Replication  
				  │  ├── CloudFront (CDN)  
				  │  └── Global Accelerator (Low-Latency Routing)  
				  └── Hybrid Connectivity  
				    ├── Direct Connect (Dedicated On-Premises Links)  
				    └── VPN (Site-to-Site/Client VPN)  

		[AWS Organization & Tenancy]  
				├── AWS Accounts  
				│  ├── Organizational Units (OUs)  
				│  ├── Resource Groups  
				│  └── IAM Policies (Resource Permissions)  
				├── Identity & Access Management (IAM)  
				│  ├── Users, Groups, Roles  
				│  ├── SSO & Federation (SAML, Cognito)  
				│  └── Permission Boundaries  
				└── Tagging (Cost Allocation, Governance)  

				[Security & Compliance]  
				├── Encryption  
				│  ├── KMS (Key Management Service)  
				│  ├── CloudHSM (Hardware Security Modules)  
				│  └── TLS/SSL Certificates (ACM)  
				├── Threat Detection  
				│  ├── GuardDuty (Malware/Threat Scanning)  
				│  ├── Inspector (Vulnerability Scanning)  
				│  └── Macie (Data Privacy Monitoring)  
				├── Network Security  
				│  ├── Web Application Firewall (WAF)  
				│  ├── Shield (DDoS Protection)  
				│  └── Network Firewall  
				└── Access Control  
				   ├── Secrets Manager (Credential Storage)  
				   └── Session Manager (SSH/RDP Alternative)  

		[Management & Observability]  
				├── Observability  
				│  ├── CloudWatch (Metrics, Logs, Alarms)  
				│  ├── CloudTrail (Audit Logs)  
				│  └── X-Ray (Distributed Tracing)  
				├── Infrastructure as Code (IaC)  
				│  ├── CloudFormation (YAML/JSON Templates)  
				│  └── CDK (Cloud Development Kit)  
				├── Cost Management  
				│  ├── Cost Explorer  
				│  ├── Budgets  
				│  └── Trusted Advisor (Optimization Tips)  
				└── Analytics  
				   ├── QuickSight (Business Intelligence)  
				   └── Athena (SQL on S3)  

		[Developer Tools]  
				├── Serverless  
				│  ├── Lambda (Event-Driven Functions)  
				│  ├── API Gateway (REST/WebSocket APIs)  
				│  └── Step Functions (Workflow Orchestration)  
				├── CI/CD  
				│  ├── CodePipeline (Deployment Pipelines)  
				│  ├── CodeBuild (Build Automation)  
				│  └── CodeDeploy (App Deployment)  
				└── SDKs & CLI  
				   ├── AWS CLI (Command-Line Interface)  
				   └── CloudShell (Browser-Based Shell)  

		[AI/ML & Advanced Services]  
				├── Machine Learning  
				│  ├── SageMaker (End-to-End ML Platform)  
				│  ├── Rekognition (Computer Vision)  
				│  └── Comprehend (Natural Language Processing)  
				├── High-Performance Compute  
				│  ├── EC2 HPC Instances (ParallelCluster)  
				│  └── Batch (Scalable Batch Processing)  
				└── IoT & Edge  
				   ├── IoT Core (Device Management)  
				   └── Outposts (Hybrid Cloud Hardware)  

		[Specialized Workloads]  
				├── Quantum Computing (Braket)  
				├── Blockchain (Managed Blockchain)  
				└── Satellite (Ground Station)  
		`
	}

	if(provider == "OCI"){
		infra = `
		[Global Infrastructure]
				└── Regions (e.g., US West, EU Frankfurt)
					├── Availability Domains (AD1, AD2, AD3)
					│     ├── Fault Domains (FD1, FD2, FD3)
					│     │     ├── Compute Instances
					│     │     │     ├── Bare Metal
					│     │     │     ├── Virtual Machines (VMs)
					│     │     │     ├── GPU Instances
					│     │     │     └── Containers (OKE)
					│     │     ├── Storage
					│     │     │     ├── Block Storage (NVMe SSD)
					│     │     │     ├── Object Storage
					│     │     │     ├── File Storage (NFS)
					│     │     │     └── Archive Storage
					│     │     ├── Networking
					│     │     │     ├── Virtual Cloud Network (VCN)
					│     │     │     │     ├── Subnets
					│     │     │     │     ├── Security Lists
					│     │     │     │     └── Route Tables
					│     │     │     ├── Load Balancers
					│     │     │     ├── SmartNICs (Isolated Network Virtualization)
					│     │     │     └── RDMA Interconnects
					│     │     └── Databases
					│     │           ├── Autonomous Database
					│     │           ├── Exadata Cloud Service
					│     │           └── Other DBs (MySQL, NoSQL)
					│     └── High-Bandwidth Interconnect (AD-to-AD)
					├── Cross-Region Replication
					└── FastConnect / VPN (On-Premises Connectivity)

			[Tenancy]
				├── Compartments
				│     ├── Resources (Compute, Storage, Networking)
				│     └── Policies (IAM Access Control)
				├── Identity and Access Management (IAM)
				│     ├── Users, Groups, Policies
				│     └── Federation (SSO)
				└── Tagging (Governance)

			[Security Layer]
				├── Encryption (At Rest, In Transit)
				├── Cloud Guard (Threat Detection)
				├── Web Application Firewall (WAF)
				├── Key Management Service (KMS)
				└── Bastion Service

			[Management & Observability]
				├── Monitoring & Logging
				├── Resource Manager (Terraform)
				├── Cost Management
				└── Analytics

			[Developer Tools]
				├── API Gateway
				├── Functions (Serverless)
				└── DevOps Services (CI/CD)

			[AI/ML Services]
				├── Data Science Platform
				├── AI Services (Vision, Language)
				└── GPU Compute
		`
	}
	dir := fmt.Sprintf("../iac/%s", provider)
	
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
		
		err = os.WriteFile(fmt.Sprintf("%s/Jenkinsfile", dir), []byte(jenkins), 0644)
		if err != nil {
			return "❌ Failed to create Jenkinsfile"
		}
		err = os.WriteFile(fmt.Sprintf("%s/architecture-guide", dir), []byte(infra), 0644)
			if err != nil {
				return "❌ Failed to write architecture file"
			}

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
