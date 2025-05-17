package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runGit(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	err := cmd.Run()
	return out.String(), err
}

func main() {
	delete := flag.Bool("d", false, "Delete stale local branches")
	remote := flag.String("r", "origin", "Remote to compare against")
	flag.Parse()

	localBranchesRaw, err := runGit("branch")
	if err != nil {
		fmt.Println("Error listing local branches:", err)
		os.Exit(1)
	}

	remoteBranchesRaw, err := runGit("ls-remote", "--heads", *remote)
	if err != nil {
		fmt.Println("Error listing remote branches:", err)
		os.Exit(1)
	}

	remoteBranches := map[string]bool{}
	scanner := bufio.NewScanner(strings.NewReader(remoteBranchesRaw))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")
		if len(parts) == 2 {
			ref := strings.TrimPrefix(parts[1], "refs/heads/")
			remoteBranches[ref] = true
		}
	}

	scanner = bufio.NewScanner(strings.NewReader(localBranchesRaw))
	var stale []string
	for scanner.Scan() {
		branch := strings.TrimSpace(scanner.Text())
		branch = strings.TrimPrefix(branch, "* ")
		if _, ok := remoteBranches[branch]; !ok && branch != "main" && branch != "master" {
			stale = append(stale, branch)
		}
	}

	if len(stale) == 0 {
		fmt.Println("✅ No stale local branches found.")
		return
	}

	fmt.Println("🧹 Stale local branches:")
	for _, b := range stale {
		fmt.Println("  ", b)
	}

	if *delete {
		for _, b := range stale {
			_, err := runGit("branch", "-D", b)
			if err != nil {
				fmt.Printf("⚠️  Failed to delete branch %s\n", b)
			} else {
				fmt.Printf("🗑️  Deleted branch: %s\n", b)
			}
		}
	}
}
