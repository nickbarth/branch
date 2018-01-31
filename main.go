package main

import (
	"bufio"
	"fmt"
	"github.com/nickbarth/chooser"
	"log"
	"os"
	"os/exec"
	"strings"
)

func (r repo) valid() bool {
	fileInfo, err := os.Stat(r.path + "/.git")
	if err != nil {
		return false
	}

	if !fileInfo.IsDir() {
		return false
	}

	return true
}

func (r repo) dirty() bool {
	cmdName := "git"
	cmdArgs := []string{"status", "--porcelain"}
	output, _ := exec.Command(cmdName, cmdArgs...).Output()
	return len(output) != 0
}

func (r repo) branches() []string {
	var branches []string

	cmdName := "git"
	cmdArgs := []string{"branch", "--list", "--format=%(refname:short)"}
	cmd := exec.Command(cmdName, cmdArgs...)
	reader, _ := cmd.StdoutPipe()
	cmd.Start()

	output := bufio.NewScanner(reader)

	for output.Scan() {
		branches = append(branches, output.Text())
	}

	return branches
}

func (r repo) setBranch(branch string) {
	cmdName := "git"
	cmdArgs := []string{"checkout", branch}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

type repo struct {
	path string
}

func main() {
	log.SetFlags(0)

	dir, err := os.Getwd()
	repo := repo{path: dir}

	if err != nil {
		log.Fatalf("Invalid directory.")
	}

	if !repo.valid() {
		log.Fatalf("`" + dir + "` is not a valid repository.")
	}

	if repo.dirty() {
		log.Fatal("`" + dir + "` has uncommitted changes.")
	}

	os.Chdir(repo.path)

	args := strings.Join(os.Args[1:], " ")
	chooser := chooser.NewChooser(5, 0)
	choice := chooser.Choose(args, repo.branches())

	if len(choice) > 0 {
		repo.setBranch(choice)
	} else {
		fmt.Println("Branch not changed.")
	}
}
