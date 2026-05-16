package installer

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveBasePath(t *testing.T) {
	t.Run("normalizes relative path", func(t *testing.T) {
		t.Setenv(basePathEnv, "")
		wd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}

		got, err := ResolveBasePath(filepath.Join("docs", "..", "ard"))
		if err != nil {
			t.Fatal(err)
		}

		want := filepath.Join(wd, "ard")
		if got != want {
			t.Fatalf("got %q, want %q", got, want)
		}
	})

	t.Run("empty path uses environment default", func(t *testing.T) {
		want := filepath.Join(t.TempDir(), "custom", "ard")
		t.Setenv(basePathEnv, want)

		got, err := ResolveBasePath("")
		if err != nil {
			t.Fatal(err)
		}

		if got != filepath.Clean(want) {
			t.Fatalf("got %q, want %q", got, filepath.Clean(want))
		}
	})
}
