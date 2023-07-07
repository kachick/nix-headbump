package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	nhb "nix-headbump"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	usageError := errors.New("expected 'bump' or 'detect' subcommands")

	detectCmd := flag.NewFlagSet("detect", flag.ExitOnError)
	bumpCmd := flag.NewFlagSet("bump", flag.ExitOnError)
	versionFlag := flag.Bool("version", false, "print the version of this program")
	currentFlag := detectCmd.Bool("current", false, "print current nixpath without bumping")
	lastFlag := detectCmd.Bool("last", false, "print git head ref without bumping")
	target := detectCmd.Bool("target", false, "print which file will be bumped")
	flag.Parse()
	if *versionFlag {
		revision := commit[:7]
		fmt.Printf("%s\n", "nix-headbump"+" "+version+" "+"("+revision+") # "+date)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println(usageError.Error())
		os.Exit(1)
	}

	path, err := nhb.GetTargetPath()
	if err != nil {
		log.Fatalf("Failed to get target files: %s", err.Error())
	}
	if path == "" {
		log.Fatalln("Both default.nix and shell.nix are not found")
	}

	switch os.Args[1] {
	case "detect":
		err := detectCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(usageError.Error())
		}
		if *target {
			fmt.Println(path)
			return
		}
		if *currentFlag {
			current, err := nhb.GetCurrentVersion(path)
			if err != nil {
				log.Fatalf("Getting the current version has been failed: %s", err.Error())
			}
			fmt.Println(current)
			return
		}
		last, err := nhb.GetLastVersion()
		if err != nil {
			log.Fatalf("Getting the last version has been failed: %s", err.Error())
		}
		if *lastFlag {
			fmt.Println(last)
			return
		}

		detectCmd.PrintDefaults()
	case "bump":
		err := bumpCmd.Parse(os.Args[2:])
		if err != nil {
			fmt.Println(usageError.Error())
		}
		last, err := nhb.GetLastVersion()
		if err != nil {
			log.Fatalf("Getting the last version has been failed: %s", err.Error())
		}
		if err = nhb.Bump(path, last); err != nil {
			log.Fatalf("Bumping the version has been failed: %s", err.Error())
		}
	default:
		fmt.Println(usageError.Error())
		os.Exit(1)
	}
}
