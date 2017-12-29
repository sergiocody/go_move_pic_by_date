package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"io"
	"log"
)

type FileTime struct {
	//Filetime
	MTime time.Time
	CTime time.Time
	ATime time.Time
}

func main() {
	var dirPath  = "/Users/codonser/Documents/PERSONAL/fotos/test"
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
						fmt.Println("*FILE : ",path)
						fmt.Println("Year : ", year)
						fmt.Println("Month : ", int(month))
						fmt.Println("Day : ", day)

						// FOLDERS: If they doesn't exist : CREATE 
						//fmt.Println("DEST PATH:" + destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month)))
						if _, err := os.Stat(destPath + "/" + strconv.Itoa(year)); os.IsNotExist(err) {
							os.Mkdir(destPath+"/"+strconv.Itoa(year), 0700)							
						}
						if _, err := os.Stat(destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month))); os.IsNotExist(err) {
							os.Mkdir(destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month)), 0700)							
						}


						var currentFile= destPath + "/" + strconv.Itoa(year)+ "/" + strconv.Itoa(int(month))+ "/"+info.Name()
						// FILE : Does it exists????
						if _, err := os.Stat(currentFile); os.IsNotExist(err) {
							//It should be moved	
							fmt.Println("Will be MOVED ", path, currentFile)

							from, err := os.Open(path)
							if err != nil {
							  log.Fatal(err)
							  return nil
							}
							defer from.Close()
						  
							to, err := os.OpenFile(currentFile, os.O_RDWR|os.O_CREATE, 0666)
							if err != nil {
							  log.Fatal(err)
							  return nil
							}
							defer to.Close()
						  
							_, err = io.Copy(to, from)
							if err != nil {
							  log.Fatal(err)
							  return nil
							}
							
							//NOW WE DELETE THE ORIGINAL FILE
							// delete file
							err = os.Remove(path)
							if err != nil {
								return nil
							}
							fmt.Println("==> done deleting file")


						} else{
							//here the file exist... is the same??
							//TODO: Check size...
							fmt.Println("*FILE ALREADY EXISTS : ",path )
						}

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


