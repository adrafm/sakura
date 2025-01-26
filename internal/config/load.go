package config

import "github.com/adrafm/sakura/internal/types"

func (s *Sakura) New(root string) error {
	config := types.Structure{
		Root: root,
	 	Folders: []string{
			"handlers",
			"middleware",
			"data",
			"migrations",
			"logs",
			"public",
			"views",
		},
	}
	
	err := s.Init(config)
	if err != nil {
		return err
	}
	
	return nil
}

func (s *Sakura) Init(structure types.Structure) error {
	root := structure.Root
	for _, path := range structure.Folders {
		err := s.CreateDirectory(root + "/" + path)
		if err != nil {
			return err
		}
	}
	
	return nil
}
