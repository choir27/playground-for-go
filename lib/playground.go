package main

import (
	"github.com/appwrite/sdk-for-go/client"
)

func main() {

	appwrite_project := "YOUR_PROJECT_ID"
	appwrite_api_key := "YOUR_API_KEY"

	appwriteClient := client.NewClient()
	appwriteClient.SetProject(appwrite_project) // Your project ID
	appwriteClient.SetKey(appwrite_api_key)     // Your API key
}
