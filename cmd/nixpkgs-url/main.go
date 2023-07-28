package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	nixurl "github.com/kachick/nixpkgs-url"
)

var (
	// Used in goreleaser
	version = "dev"
	commit  = "none"
	date    = "unknown"

	revision = "rev"
)

func mustGetTargetPath() string {
	path, err := nixurl.GetTargetPath()
	if err != nil {
		log.Fatalf("Failed to get target files: %s", err.Error())
	}
	if path == "" {
		log.Fatalln("Any *.nix files are not found")
	}

	return path
}

func main() {
	const usage = `Usage: nixpkgs-url <subcommand> <flags>

$ nixpkgs-url detect -current
$ nixpkgs-url bump
$ nixpkgs-url -version`

	detectCmd := flag.NewFlagSet("detect", flag.ExitOnError)
	bumpCmd := flag.NewFlagSet("bump", flag.ExitOnError)
	versionFlag := flag.Bool("version", false, "print the version of this program")
	currentFlag := detectCmd.Bool("current", false, "print current nixpath without bumping")
	lastFlag := detectCmd.Bool("last", false, "print git head ref without bumping")
	targetFlag := detectCmd.Bool("target", false, "print which file will be bumped")

	// Do not call as xdg-open for WSL2, URL will be displayed as a clickable in newer terminals, it is enough
	// https://github.com/microsoft/WSL/issues/8892
	jumpFlag := detectCmd.Bool("jump", false, "print reasonable URL for the ref")

	flag.Usage = func() {
		// https://github.com/golang/go/issues/57059#issuecomment-1336036866
		fmt.Printf("%s", usage+"\n\n")
		fmt.Println("Usage of command:")
		flag.PrintDefaults()
		fmt.Println("")
		detectCmd.Usage()
		fmt.Println("")
		bumpCmd.Usage()
	}

	if len(commit) >= 7 {
		revision = commit[:7]
	}
	version := fmt.Sprintf("%s\n", "nixpkgs-url"+" "+version+" "+"("+revision+") # "+date)

	flag.Parse()
	if *versionFlag {
		fmt.Println(version)
		return
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "detect":
		err := detectCmd.Parse(os.Args[2:])
		if err != nil {
			flag.Usage()
		}

		if *lastFlag {
			last, err := nixurl.GetLastVersion()
			if err != nil {
				log.Fatalf("Getting the last version has been failed: %s", err.Error())
			}
			if *jumpFlag {
				fmt.Println("https://github.com/NixOS/nixpkgs/commit/" + last)
			} else {
				fmt.Println(last)
			}
			return
		}

		path := mustGetTargetPath()
		if *targetFlag {
			fmt.Println(path)
			return
		}

		if *currentFlag {
			current, err := nixurl.GetCurrentVersion(path)
			if err != nil {
				log.Fatalf("Getting the current version has been failed: %s", err.Error())
			}

			if *jumpFlag {
				fmt.Println("https://github.com/NixOS/nixpkgs/commit/" + current)
			} else {
				fmt.Println(current)
			}

			return
		}

		detectCmd.Usage()
	case "bump":
		err := bumpCmd.Parse(os.Args[2:])
		if err != nil {
			flag.Usage()
		}
		last, err := nixurl.GetLastVersion()
		if err != nil {
			bumpCmd.Usage()
			log.Fatalf("Getting the last version has been failed: %s", err.Error())
		}
		if err = nixurl.Bump(mustGetTargetPath(), last); err != nil {
			bumpCmd.Usage()
			log.Fatalf("Bumping the version has been failed: %s", err.Error())
		}
	default:
		flag.Usage()

		os.Exit(1)
	}
}
