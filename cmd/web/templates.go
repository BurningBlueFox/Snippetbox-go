package main

import "github.com/BurningBlueFox/letsgo/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []models.Snippet
}
