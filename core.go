package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

var (
	re = regexp.MustCompile(`(?s)(import\s+\(fetchTarball\s+"https://github.com/NixOS/nixpkgs/archive/)([^"]+?)(\.tar\.gz"\))`)
)

func Bump(path string, last string) error {
	origin, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	replaced := re.ReplaceAll(origin, []byte("${1}"+last+"${3}"))
	if bytes.Equal(origin, replaced) {
		return nil
	}

	return os.WriteFile(path, replaced, os.ModePerm)
}

func GetTargetPath() (string, error) {
	paths := [3]string{"flake.nix", "default.nix", "shell.nix"}
	for _, path := range paths {
		_, err := os.Stat(path)
		if err == nil {
			return path, nil
		}

		if !os.IsNotExist(err) {
			return "", fmt.Errorf("can not open %s: %w", path, err)
		}
	}
	return "", fmt.Errorf("%v are not found", paths)
}

func GetCurrentVersion(path string) (string, error) {
	origin, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	matches := re.FindStringSubmatch(string(origin))
	if err != nil || len(matches) < 2 {
		return "", err
	}

	return matches[2], nil
}

func GetLastVersion() (string, error) {
	type commit struct {
		Sha string `json:"sha"`
	}

	type response struct {
		Commit commit `json:"commit"`
	}

	// https://docs.github.com/en/rest/branches/branches?apiVersion=2022-11-28#get-a-branch
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/NixOS/nixpkgs/branches/master", nil)
	// May be necessary to set "Authorization" header if frequent requests are needed.
	// -H "Authorization: Bearer <YOUR-TOKEN>"\
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	jsonRes := &response{}
	if json.Unmarshal(body, jsonRes) != nil {
		return "", err
	}
	return jsonRes.Commit.Sha, nil
}
