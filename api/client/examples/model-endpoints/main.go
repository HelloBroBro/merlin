package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/caraml-dev/merlin/client"
	"github.com/caraml-dev/mlp/api/pkg/auth"
)

func main() {
	ctx := context.Background()

	basePath := "http://merlin.dev/api/merlin/v1"
	if os.Getenv("MERLIN_API_BASEPATH") != "" {
		basePath = os.Getenv("MERLIN_API_BASEPATH")
	}

	// Create an HTTP client with Google default credential
	httpClient := http.DefaultClient
	googleClient, err := auth.InitGoogleClient(context.Background())
	if err == nil {
		httpClient = googleClient
	} else {
		log.Println("Google default credential not found. Fallback to HTTP default client")
	}

	cfg := client.NewConfiguration()
	cfg.Servers = client.ServerConfigurations{
		{
			URL:         basePath,
			Description: "Merlin base path",
		},
	}
	cfg.HTTPClient = httpClient

	apiClient := client.NewAPIClient(cfg)

	// Get all projects
	projects, _, err := apiClient.ProjectAPI.ProjectsGet(ctx).Execute()
	if err != nil {
		panic(err)
	}
	log.Println("Projects:", projects)

	for _, project := range projects {
		log.Println()
		log.Println("---")
		log.Println()

		log.Println("Project:", project.Name)

		// Get all models in the given project

		models, _, err := apiClient.ModelsAPI.ProjectsProjectIdModelsGet(ctx, project.Id).Execute()
		if err != nil {
			panic(err)
		}
		log.Println("Models:", models)

		for _, model := range models {
			log.Println()
			log.Println("- Model:", model.Name)

			// Get all model's endpoints
			modelEndpoints, _, err := apiClient.ModelEndpointsAPI.ModelsModelIdEndpointsGet(ctx, *model.Id).Execute()
			if err != nil {
				panic(err)
			}

			if len(modelEndpoints) == 0 {
				log.Println("  Model endpoints: not available")
				continue
			}

			log.Println("  Model endpoints:")
			for _, modelEndpoint := range modelEndpoints {
				log.Printf("  - %v: %v\n", modelEndpoint.EnvironmentName, modelEndpoint.Url)
			}
		}
	}
}
