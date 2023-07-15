package core

import (
	"os"
	"testing"
)

func TestGetCurrentVersion(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	t.Cleanup(func() {
		err := os.Chdir(cwd)
		if err != nil {
			t.Fatalf("failed to rollback working directory: %v", err)
		}
	})
	err = os.Chdir("testdata")
	if err != nil {
		t.Fatalf("failed to walk through testdata directory: %v", err)
	}

	got, err := GetCurrentVersion("default.nix")
	if err != nil {
		t.Fatalf("Getting the version has been failed: %s", err.Error())
	}
	want := "e57b65abbbf7a2d5786acc86fdf56cde060ed026"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

// Calling actual GitHub API, May be necessary to stub or disabling in CI
func TestGetLastVersion(t *testing.T) {
	got, err := GetLastVersion()
	if err != nil {
		t.Fatalf("Getting the last version has been failed: %s", err.Error())
	}
	wantLength := 40

	if len(got) != wantLength {
		t.Errorf("got %q, wanted %q", got, "a string that have 40 length")
	}
}

func TestTargetPath(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	t.Cleanup(func() {
		err := os.Chdir(cwd)
		if err != nil {
			t.Fatalf("failed to rollback working directory: %v", err)
		}
	})
	err = os.Chdir("testdata")
	if err != nil {
		t.Fatalf("failed to walk through testdata directory: %v", err)
	}

	got, err := GetTargetPath()
	if err != nil {
		t.Fatalf("Failed to get target files: %s", err.Error())
	}
	want := "default.nix"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
