package assets

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestMirroredAgentAssetsMatchSources(t *testing.T) {
	sourceRoot := filepath.Join("..", "..", "agents")
	mirrorRoot := "agents"

	sourceFiles := map[string]struct{}{}

	err := filepath.WalkDir(sourceRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(sourceRoot, path)
		if err != nil {
			return err
		}
		sourceFiles[filepath.ToSlash(rel)] = struct{}{}

		sourceContent, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		mirrorPath := filepath.Join(mirrorRoot, filepath.FromSlash(filepath.ToSlash(rel)))
		mirrorContent, err := os.ReadFile(mirrorPath)
		if err != nil {
			return err
		}

		if string(sourceContent) != string(mirrorContent) {
			t.Fatalf("mirrored asset drift for %s", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	err = filepath.WalkDir(mirrorRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		rel, err := filepath.Rel(mirrorRoot, path)
		if err != nil {
			return err
		}
		if _, ok := sourceFiles[filepath.ToSlash(rel)]; !ok {
			t.Fatalf("unexpected mirrored asset without source counterpart: %s", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
