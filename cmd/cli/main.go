package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/varshaprasad96/mcp-kserve/pkg/kserve"
)

func main() {
	namespace := flag.String("namespace", "default", "Kubernetes namespace")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Usage: mcp-kserve [command] [args]")
		fmt.Println("Commands:")
		fmt.Println("  deploy <name> <framework> <modelURI>")
		fmt.Println("  list")
		fmt.Println("  status <name>")
		fmt.Println("  delete <name>")
		os.Exit(1)
	}

	client, err := kserve.NewClient(*namespace)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()
	cmd := args[0]

	switch cmd {
	case "deploy":
		if len(args) != 4 {
			log.Fatal("Usage: deploy <name> <framework> <modelURI>")
		}
		err = client.DeployModel(ctx, args[1], args[3], args[2])
		if err != nil {
			log.Fatalf("Failed to deploy model: %v", err)
		}
		fmt.Printf("Model %s deployment initiated\n", args[1])

	case "list":
		services, err := client.ListModels(ctx)
		if err != nil {
			log.Fatalf("Failed to list models: %v", err)
		}
		for _, svc := range services.Items {
			fmt.Printf("%s: %s\n", svc.Name, svc.Status.URL)
		}

	case "status":
		if len(args) != 2 {
			log.Fatal("Usage: status <name>")
		}
		service, err := client.GetModelStatus(ctx, args[1])
		if err != nil {
			log.Fatalf("Failed to get model status: %v", err)
		}
		fmt.Printf("Status: %s\n", service.Status.URL)

	case "delete":
		if len(args) != 2 {
			log.Fatal("Usage: delete <name>")
		}
		err = client.DeleteModel(ctx, args[1])
		if err != nil {
			log.Fatalf("Failed to delete model: %v", err)
		}
		fmt.Printf("Model %s deleted\n", args[1])

	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
