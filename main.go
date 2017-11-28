package main
import (
	"log"
	"os"
	"path/filepath"
)



func main() {
	var dirPath string = "/Users/codonser/Documents/PERSONAL/fotos/sergio"
	// walk all files in directory
	extensions := []string{".mp4", ".jpeg",".png"}
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, extension := range extensions {
				log.Println(filepath.Ext(path))
				if extension == filepath.Ext(path) {
					log.Println(dirPath, " matches ", extension)
					return nil
				}	else {

					log.Println("NO")
				}
			}
		}
		return nil
	})
}