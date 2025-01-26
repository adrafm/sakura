package sakura

import (
	"github.com/adrafm/sakura/internal/config"
)

// create interface for the user to create the sakura module
type Sakura config.Sakura
func (s *Sakura) New(root string) error {
	return (&config.Sakura{}).New(root)
}
