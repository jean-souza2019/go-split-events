package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	inputEvent    string
	inputDir      string
	maxGoroutines int
	counter       int
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Deve-se passar o diretório dos arquivos e o evento")
	}

	inputDir = os.Args[1]
	inputEvent = os.Args[2]
	maxGoroutines = 16

	if isMatchEvent := matchEvent(inputEvent); isMatchEvent == false {
		log.Fatalf("Evento não existente. Evento: %v", inputEvent)
	}

	files, err := os.ReadDir(inputDir)
	if err != nil {
		go log.Printf("Erro ao processar arquivos: %v", err)
	}

	var wg sync.WaitGroup
	guard := make(chan struct{}, maxGoroutines)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		wg.Add(1)
		guard <- struct{}{}

		go func(file os.DirEntry) {
			defer wg.Done()
			defer func() { <-guard }()

			if err := splitAndMoveFile(file, inputDir, inputEvent); err != nil {
				go log.Printf("Erro ao mover arquivo %s: %v", file.Name(), err)
			}
		}(file)
	}

	wg.Wait()

}

func splitAndMoveFile(file os.DirEntry, inputDir string, inputEvent string) error {
	inputEventPrepared := strings.ToUpper(inputEvent)

	if strings.Contains(strings.ToUpper(file.Name()), inputEventPrepared) {
		outputPath := filepath.Join(inputDir, "EXCLUIDOS", inputEventPrepared)
		inputFilePath := filepath.Join(inputDir, file.Name())
		outputFilePath := filepath.Join(outputPath, file.Name())

		if err := createDir(outputPath); err != nil {
			return err
		}

		if err := moveFile(inputFilePath, outputFilePath); err != nil {
			return err
		}

		counter++
		log.Printf("Moveu até agora %v arquivos", counter)
	}

	return nil
}

func moveFile(inputPath, outputPath string) error {
	if err := os.Rename(inputPath, outputPath); err != nil {
		go log.Printf("Erro ao mover arquivo: %s para %s, erro: %s", inputPath, outputPath, err)
		return err
	}
	return nil
}

func createDir(dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Printf("Erro ao criar diretório: %v", err)
		return err
	}
	return nil
}

func matchEvent(inputEvent string) bool {
	events := []string{
		"S-2200",
		"S-2190",
		"S-2205",
		"S-2206",
		"S-2230",
		"S-2299",
		"S-3000",
		"S-2210",
		"S-2220",
		"S-2240",
		"S-1200",
		"S-2306",
		"S-2399",
		"S-2300",
		"S-1210",
		"S-1030",
		"S-1050",
		"S-1010",
		"S-5001",
		"S-5002",
		"S-5003",
	}

	for _, event := range events {
		if event == strings.ToUpper(inputEvent) {
			return true
		}
	}
	return false
}
