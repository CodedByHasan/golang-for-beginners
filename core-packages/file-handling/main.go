package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path1 := filepath.Join("dir1", "dir2", "text.txt")
	
	fmt.Printf("filepath.Join: %s\n", path1)
	fmt.Printf("filepath.dir: %s\n", filepath.Dir(path1))

	// filepath.Join will normalise filepath
	path2 := filepath.Join("dir1", "dir2/../dir3", "text.txt")
	
	fmt.Printf("filepath.Join (normalised): %s\n", path2)
	fmt.Printf("filepath.Base: %s\n", filepath.Base(path2))

	fmt.Printf("filepath.IsAbs: %t\n", filepath.IsAbs(path2))
	fmt.Printf("filepath.IsAbs: %t\n", filepath.IsAbs("/dir/file"))
	
	fmt.Printf("filepath.Ext: %s\n", filepath.Ext(path2))
}
