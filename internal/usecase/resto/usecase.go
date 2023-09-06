package resto

import "github.com/autumnleaf-ra/go-resto-app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
