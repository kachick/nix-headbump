package core

import (
	"os"
	"testing"
)

func TestGetCurrentVersionForFlakeNix(t *testing.T) {
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
	err = os.Chdir("testdata/flake")
	if err != nil {
		t.Fatalf("failed to walk through testdata directory: %v", err)
	}

	got, err := GetCurrentVersion("flake.nix")
	if err != nil {
		t.Fatalf("Getting the version has been failed: %s", err.Error())
	}
	want := "e57b65abbbf7a2d5786acc86fdf56cde060ed026"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGetCurrentVersionForDefaultNix(t *testing.T) {
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
	err = os.Chdir("testdata/classic")
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

func TestGetCurrentVersionInEmptyDir(t *testing.T) {
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
	err = os.Chdir("testdata/nothing")
	if err != nil {
		t.Fatalf("failed to walk through testdata directory: %v", err)
	}

	_, err = GetCurrentVersion("default.nix")
	if !os.IsNotExist(err) {
		t.Errorf("returned unexpected error: %v", err)
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

func TestClassicTargetPath(t *testing.T) {
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
	err = os.Chdir("testdata/classic")
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

func TestTargetPathMultipleCandidates(t *testing.T) {
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
	err = os.Chdir("testdata/candidates")
	if err != nil {
		t.Fatalf("failed to walk through testdata directory: %v", err)
	}

	got, err := GetTargetPath()
	if err != nil {
		t.Fatalf("Failed to get target files: %s", err.Error())
	}
	want := "flake.nix"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
