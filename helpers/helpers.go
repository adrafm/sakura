package helpers

import "os"

func CreateDirectory(path string) error {
	const permission = 0755
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, permission)
		if err != nil {
			return err
		}
	}
	
	return nil
}
