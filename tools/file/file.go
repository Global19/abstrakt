package file

import "os"

// Exists - basic utility function to check the provided filename can be opened and is not a folder/directory
func Exists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
