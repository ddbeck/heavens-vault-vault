package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const programName string = "heavens-vault-vault"
const backupDir string = "hvvBackupSaves"

func isSave(fileInfoName string) bool {
	regular := "heavensVaultSave.json"
	return fileInfoName == regular
}

func makeName(fileInfoName string, date string) string {
	base := strings.Split(fileInfoName, ".json")[0]
	return fmt.Sprintf("%s.%s.json", base, date)
}

func makeDest(fileInfoName string, date string) string {
	newName := makeName(fileInfoName, date)
	return filepath.Join(backupDir, newName)
}

func makeBackupDirectory(dirPath string) {
	dest := filepath.Join(dirPath, backupDir)
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		fmt.Printf("Creating %s\n", dest)
		os.Mkdir(dest, 0700)
	}
}

func copy(src string, dest string) error {
	srcFile, srcErr := os.Open(src)

	if srcErr != nil {
		return srcErr
	}
	defer srcFile.Close()

	destFile, destErr := os.Create(dest)
	if destErr != nil {
		return destErr
	}
	defer destFile.Close()

	_, copyErr := io.Copy(destFile, srcFile)
	if copyErr != nil {
		return copyErr
	}

	syncErr := destFile.Sync()
	if syncErr != nil {
		return syncErr
	}

	return nil
}

func backup(dirPath string) {
	makeBackupDirectory(dirPath)

	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info.Name() != "Heaven's Vault" && info.IsDir() {
			return filepath.SkipDir
		}

		if name := info.Name(); isSave(name) {
			formatted := info.ModTime().Format("2006-01-02-15_04_05")
			dest := filepath.Join(dirPath, makeDest(name, formatted))

			// don't bother copying the file if it's already copied
			_, existenceErr := os.Stat(dest)
			if existenceErr != nil {
				msg := fmt.Sprintf("Copying %s ➡ %s", path, dest)
				fmt.Println(msg)

				copyErr := copy(path, dest)
				if copyErr != nil {
					fmt.Println(fmt.Sprintf("WARNING: %s", copyErr))
				}
			} else {
				msg := fmt.Sprintf("Skipping %s", path)
				fmt.Println(msg)
			}
		}

		return nil
	})
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "%s requires one argument only, the path to your save games", programName)
	}

	directory, _ := filepath.Abs(args[0])

	fmt.Println(fmt.Sprintf("Starting to watch: %s", directory))
	for {
		backup(directory)
		wait := time.Duration(10) * time.Minute
		fmt.Print(fmt.Sprintf("Waiting %s…", wait))
		time.Sleep(wait)
		fmt.Println(" done.")
	}
}
