package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	nixurl "github.com/kachick/nixpkgs-url"
)

var (
	// Used in goreleaser
	version = "dev"
	commit  = "none"
	date    = "unknown"

	revision = "rev"
)

// https://stackoverflow.com/posts/39324149/revisions
// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
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
	jumpFlag := detectCmd.Bool("jump", false, "open github URL with current ref")

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

	path, err := nixurl.GetTargetPath()
	if err != nil {
		log.Fatalf("Failed to get target files: %s", err.Error())
	}
	if path == "" {
		log.Fatalln("Any *.nix files are not found")
	}

	switch os.Args[1] {
	case "detect":
		err := detectCmd.Parse(os.Args[2:])
		if err != nil {
			flag.Usage()
		}
		if *targetFlag {
			fmt.Println(path)
			return
		}
		if *currentFlag {
			current, err := nixurl.GetCurrentVersion(path)
			if err != nil {
				log.Fatalf("Getting the current version has been failed: %s", err.Error())
			}
			fmt.Println(current)

			if *jumpFlag {
				err = open("https://github.com/NixOS/nixpkgs/commit/" + current)
				if err != nil {
					log.Fatalf(err.Error())
				}
			}

			return
		}
		last, err := nixurl.GetLastVersion()
		if err != nil {
			log.Fatalf("Getting the last version has been failed: %s", err.Error())
		}
		if *lastFlag {
			fmt.Println(last)
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
		if err = nixurl.Bump(path, last); err != nil {
			bumpCmd.Usage()
			log.Fatalf("Bumping the version has been failed: %s", err.Error())
		}
	default:
		flag.Usage()

		os.Exit(1)
	}
}
