package main

import (
	"github.com/joho/godotenv"
	"github.com/nenitf/devtome/pkg/devtome"
	"log"
	"os"
)

func main() {
	err := godotenv.Load("config.txt")
	if err != nil {
		log.Fatalf("Erro: %s", err)
	}

	token := os.Getenv("DEVTOME_TOKEN")
	c := devtome.NewClient("https://dev.to", token)

	if token == "" {
		log.Fatalf("DEVTOME_TOKEN não especificado")
	}

	path := os.Getenv("DEVTOME_PATH")
	if path == "" {
		log.Fatalf("DEVTOME_PATH não especificado")
	}

	arts, err := c.GetAll()
	if err != nil {
		log.Fatalf("Erro: %s", err)
	}

	err = devtome.ArticlesPersistence(path, arts)
	if err != nil {
		log.Fatalf("Erro: %s", err)
	}
}
