package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const outputFile = "autoload_gen.go"

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	entries, err := os.ReadDir(workingDir)
	if err != nil {
		panic(err)
	}

	imports := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		name := entry.Name()
		if strings.HasPrefix(name, ".") || name == "registry" {
			continue
		}
		imports = append(imports, fmt.Sprintf("\t_ \"micro-mall.dev/app/user/internal/controller/%s\"", name))
	}
	sort.Strings(imports)

	var content bytes.Buffer
	content.WriteString("package controller\n\n")
	if len(imports) == 0 {
		content.WriteString("// No controller packages found.\n")
	} else {
		content.WriteString("import (\n")
		content.WriteString(strings.Join(imports, "\n"))
		content.WriteString("\n)\n")
	}

	outputPath := filepath.Join(workingDir, outputFile)
	if err = os.WriteFile(outputPath, content.Bytes(), 0o644); err != nil {
		panic(err)
	}
}
