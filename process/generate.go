package process

import (
	"fmt"
	"os"
	"path/filepath"

	"kego.io/kerr"
	"kego.io/process/generate"
	"kego.io/process/scan"
	"kego.io/process/settings"
)

type sourceType string

const (
	S_STRUCTS sourceType = "structs"
	S_TYPES              = "types"
	S_EDITOR             = "editor"
)

// Generate generates the source code from templates and writes the files
// to the correct folders.
//
// file == F_STRUCTS: generated-structs.go in the root of the package.
//
// file == F_TYPES: generated-types.go containing advanced type information
// in the "types" sub package. Note that to generate this file, we need to
// have the main generated-structs.go compiled in, so we generate a temporary
// command and run it with "go run".
//
// file == F_EDITOR: generated-editor.go in the "editor" sub-package. This
// will be compiled to JS when the editor is launched.
//
func Generate(file sourceType, set *settings.Settings) error {

	if set.Verbose {
		fmt.Print("Generating ", file, "... ")
	}

	if file == S_STRUCTS {
		hasFiles, err := scan.HasSourceFiles(set)
		if err != nil {
			return kerr.New("GXGGDQVHHP", err, "ScanForKegoFiles")
		}
		if !hasFiles {
			return kerr.New("YSLREDFDLJ", nil, "No kego files found")
		}
	}

	if file == S_STRUCTS || file == S_TYPES {

		// We only tolerate unknown types when we're initially building the
		// struct files. At all other times, the generated structs should
		// provide all types.
		ignoreUnknownTypes := file == S_STRUCTS

		// When generating structs or types, we need to scan for types. All other runs will have
		// them compiled in the types sub-package.
		if err := scan.ScanForTypes(ignoreUnknownTypes, set); err != nil {
			return kerr.New("XYIUHERDHE", err, "scan.ScanForTypes")
		}
	}

	var outputDir string
	var filename string
	var source []byte
	var err error

	switch file {
	case S_STRUCTS:
		outputDir = set.Dir
		filename = "generated-structs.go"
		source, err = generate.Structs(set)
	case S_TYPES:
		outputDir = filepath.Join(set.Dir, "types")
		filename = "generated-types.go"
		source, err = generate.Types(set)
	case S_EDITOR:
		outputDir = filepath.Join(set.Dir, "editor")
		filename = "generated-editor.go"
		source, err = generate.Editor(set)
	}
	if err != nil {
		return kerr.New("XFNESBLBTQ", err, "generate: %s", file)
	}

	// We only backup in the system structs and types files because they are the only
	// generated files we ever need to roll back
	backup := set.Path == "kego.io/system" && (file == S_STRUCTS || file == S_TYPES)

	if err = save(outputDir, source, filename, backup); err != nil {
		return kerr.New("UONJTTSTWW", err, "save")
	} else {
		if set.Verbose {
			fmt.Println("OK.")
		}
	}

	return nil
}

func save(dir string, contents []byte, name string, backup bool) error {

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return kerr.New("BPGOUIYPXO", err, "os.MkdirAll")
		}
	}

	file := filepath.Join(dir, name)

	if backup {
		backupPath := filepath.Join(dir, fmt.Sprintf("%s.backup", name))
		if _, err := os.Stat(backupPath); err == nil {
			os.Remove(backupPath)
		}
		if _, err := os.Stat(file); err == nil {
			os.Rename(file, backupPath)
		}
	} else {
		os.Remove(file)
	}

	if len(contents) == 0 {
		return nil
	}

	output, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return kerr.New("NWLWHSGJWP", err, "os.OpenFile (could not open output file)")
	}
	defer output.Close()

	if _, err := output.Write(contents); err != nil {
		return kerr.New("FBMGPRWQBL", err, "output.Write")
	}

	if err := output.Sync(); err != nil {
		return kerr.New("EGFNTMNKFX", err, "output.Sync")
	}

	return nil
}
