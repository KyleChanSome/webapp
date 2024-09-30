package main

import "github.com/KyleChanSome/webapp/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
