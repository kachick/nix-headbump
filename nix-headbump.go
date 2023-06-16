package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

func main() {
	path := "default.nix"
	isNixFileExist := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		path = "shell.nix"
		if _, err := os.Stat(path); os.IsNotExist(err) {
			isNixFileExist = false
		}
	}

	if isNixFileExist {
		err := bump(path)
		if err != nil {
			log.Fatalf("Bumping the version has been failed: %s", err.Error())
		}
	} else {
		log.Fatalln("Both default.nix and shell.nix are not found")
	}
}

type Commit struct {
	Sha string `json:"sha"`
}

type Response struct {
	Commit Commit `json:"commit"`
}

func bump(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`(import \(fetchTarball "https://github.com/NixOS/nixpkgs/archive/)(?:[^"]+?)(\.tar\.gz"\))`)
	req, _ := http.NewRequest("GET", "https://api.github.com/repos/NixOS/nixpkgs/branches/master", nil)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	jsonRes := &Response{}
	if json.Unmarshal(body, jsonRes) != nil {
		return err
	}
	replaced := re.ReplaceAll(bytes, []byte("${1}"+jsonRes.Commit.Sha+"${2}"))
	return os.WriteFile(path, replaced, os.ModePerm)
}
