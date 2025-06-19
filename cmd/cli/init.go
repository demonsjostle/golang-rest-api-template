package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// รับ name จาก flag ก่อน
	newName := flag.String("name", "", "new module name")
	flag.Parse()

	// ถ้าไม่ใส่ flag ให้ prompt แบบ interactive
	if *newName == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter new module name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("failed to read input: %v", err)
		}
		*newName = strings.TrimSpace(input) // ตัด \n ออก  [oai_citation:0‡tutorialedge.net](https://tutorialedge.net/golang/reading-console-input-golang/?utm_source=chatgpt.com)
		if *newName == "" {
			log.Fatal("module name cannot be empty")
		}
	}

	// step 1: remove .github directory if exists
	if err := os.RemoveAll(".github"); err != nil && !os.IsNotExist(err) {
		log.Fatalf("failed to remove .github: %v", err)
	}

	// step 2: replace module imports
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// ข้ามโฟลเดอร์ .git และ vendor
		if d.IsDir() && (d.Name() == ".git" || d.Name() == "vendor") {
			return filepath.SkipDir
		}
		// สนใจแค่ .go และ go.mod
		if !d.IsDir() {
			ext := filepath.Ext(d.Name())
			if ext != ".go" && d.Name() != "go.mod" {
				return nil
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			if !bytes.Contains(content, []byte("goraft")) {
				return nil
			}
			newContent := strings.ReplaceAll(string(content), "goraft", *newName)
			return os.WriteFile(path, []byte(newContent), d.Type().Perm())
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error during init: %v", err)
	}

	fmt.Println("Initialization complete")
}
