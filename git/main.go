package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

	"golang.org/x/crypto/ssh"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport"
	go_git_ssh "gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"
)

func getSSHKeyAuth(privateSSHKeyFile string) transport.AuthMethod {
	sshKey, _ := ioutil.ReadFile(privateSSHKeyFile)
	signer, _ := ssh.ParsePrivateKey([]byte(sshKey))
	return &go_git_ssh.PublicKeys{User: "git", Signer: signer}
}

func runCMD(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	if dir != "" {
		cmd.Dir = dir
	}
	return cmd.Run()
}

func printOnError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

var (
	rules = make(map[string]string)
)

// writeToFile interface
func writeToFile(name, content string) error {
	rules[name] = content
	return nil
}

func run(name, rule string) error {
	return nil
}

func parse() error {
	return nil
}

var (
	src = `
package main

func main() { fmt.Print("ok") }
`
)

func main() {
	url := "ssh://git@git.chotot.org:2222/robocop/papa-plugin-template.git"
	directory, _ := os.Getwd()
	// printOnError(runCMD("mkdir", "-p", "papa-plugin-template"))
	tmpDir, err := ioutil.TempDir(directory, "papa-plugin-template")
	if err != nil {
		panic(err)
	}

	fmt.Println(tmpDir)
	// defer os.RemoveAll(tmpDir)

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Println(usr)

	auth := getSSHKeyAuth(usr.HomeDir + "/.ssh/id_rsa")
	fmt.Println(auth)

	_, err = git.PlainClone(tmpDir, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Auth:              auth,
	})
	if err != nil {
		panic(err)
	}

	// ref, err := r.Head()
	// if err != nil {
	// 	panic(err)
	// }

	// commit, err := r.CommitObject(ref.Hash())
	// fmt.Println(commit, err)
	in := filepath.Join(tmpDir, "rules/rule.go")
	if err := ioutil.WriteFile(in, []byte(src), 0400); err != nil {
		panic(err)
	}

	printOnError(runCMD(tmpDir, "goimports", "-w", "-v", "./rules"))
	printOnError(runCMD(tmpDir, "gofmt", "-w", "-v", "./rules"))

	printOnError(runCMD(tmpDir, "dep", "init", "-v"))
	printOnError(runCMD(tmpDir, "dep", "ensure", "-v", "-vendor-only"))
	printOnError(runCMD(tmpDir, "docker", "build", "-t", "docker.chotot.org/rule:0.0.1", "."))
}
