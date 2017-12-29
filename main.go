package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
)

type FileTime struct {
	MTime time.Time
	CTime time.Time
	ATime time.Time
}

func main() {
	var dirPath string = "/Users/codonser/Documents/PERSONAL/fotos/test"
	var destPath = "/Users/codonser/Documents/PERSONAL/fotos/test/2"

	// walk all files in directory
	extensions := []string{".mp4", ".jpeg", ".avi"}
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			for _, extension := range extensions {
				if extension == filepath.Ext(path) {
					file, err := FTime(path)
					if err == nil {
						fmt.Println(path, "Mo Time", file.MTime)
						fi, err := os.Stat(path)
						if err != nil {
							return nil
						}
						var mtime time.Time
						mtime = fi.ModTime()
						year, month, day := mtime.Date()
						fmt.Println("Year : ", year)
						fmt.Println("Month : ", int(month))
						fmt.Println("Day : ", day)

						if _, err := os.Stat(destPath + "/" + strconv.Itoa(year)); os.IsNotExist(err) {
							os.Mkdir(destPath+"/"+strconv.Itoa(year), 0700)
							if _, err := os.Stat(destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month))); os.IsNotExist(err) {
								os.Mkdir(destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month)), 0700)
								fmt.Println("DEST PATH:" + destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month)))
							}
						}
						//fmt.Println( "Cr Time", file.CTime)
						//fmt.Println("Ac Time", file.ATime)

						//NOW MOVING TO THE NEW PATH

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

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
