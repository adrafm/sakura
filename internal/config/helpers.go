package config

import "github.com/adrafm/sakura/helpers"

func (s *Sakura) CreateDirectory(path string) error {
	return helpers.CreateDirectory(path)
}