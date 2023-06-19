package main

import "testing"

func TestGetCurrentVersion(t *testing.T) {
	got, err := getCurrentVersion("default.nix")
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
	got, err := getLastVersion()
	if err != nil {
		t.Fatalf("Getting the last version has been failed: %s", err.Error())
	}
	wantLength := 40

	if len(got) != wantLength {
		t.Errorf("got %q, wanted %q", got, "a string that have 40 length")
	}
}
