package menu

import "github.com/autumnleaf-ra/go-resto-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
