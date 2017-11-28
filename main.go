package main
import (
	"os"
	"path/filepath"
	"fmt"
	"syscall"
	"time"
)

type FileTime struct {
	MTime time.Time
	CTime time.Time
	ATime time.Time
}


func main() {
	var dirPath string = "/Users/codonser/Documents/PERSONAL/fotos/mapi"
	// walk all files in directory
	extensions := []string{".mp4", ".jpeg",".avi"}
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, extension := range extensions {
				if extension == filepath.Ext(path) {
					file, err := FTime(path)
					if err == nil {
						fmt.Println(path, "Mo Time", file.MTime)
						//fmt.Println( "Cr Time", file.CTime)
						//fmt.Println("Ac Time", file.ATime)
					}
				}
			}
		}
		return nil
	})
}


// Gets the Modified, Create and Access time of a file
func FTime(file string) (t *FileTime, err error) {
	fileinfo, err := os.Stat(file)
	if err != nil {
		return
	}
	t = new(FileTime)
	var stat = fileinfo.Sys().(*syscall.Stat_t)
	t.ATime = time.Unix(stat.Atimespec.Sec, stat.Atimespec.Nsec)
	t.CTime = time.Unix(stat.Ctimespec.Sec, stat.Ctimespec.Nsec)
	t.MTime = time.Unix(stat.Mtimespec.Sec, stat.Mtimespec.Nsec)
	return
}