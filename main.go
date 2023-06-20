package main

import (
	"fmt"
	"lineaje-assignment/pkg/dependencyGraph"
	"lineaje-assignment/pkg/formatString"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

const (
	repoURL          = "https://github.com/etcd-io/etcd"
	branchName       = "main"
	dataPath         = "data/etcd"
	relativeDataPath = "./" + dataPath
	gomodBinPath     = "../../bin/gomod"
	goBackPath       = "../.."
)

func main() {
	// arguments := os.Args
	// repoURL, branchName := arguments[0], arguments[1]

	_, err := git.PlainClone(relativeDataPath, false, &git.CloneOptions{
		URL:           repoURL,
		ReferenceName: plumbing.ReferenceName(branchName),
	})
	if err != nil {
		fmt.Println("Download failed", err)
		return
	} else {
		fmt.Println("Download successful")
	}

	err = os.Chdir(dataPath)
	if err != nil {
		fmt.Println("Change directory failed", err)
		return
	}

	terminalCmd := exec.Command(gomodBinPath, "graph")
	graphDataByte, err := terminalCmd.Output()
	if err != nil {
		fmt.Println("Graph query failed", err)
		return
	}

	err = os.Chdir(goBackPath)
	if err != nil {
		fmt.Println("Change directory failed", err)
		return
	}

	dependencies := formatString.FormatString(string(graphDataByte))

	err = dependencyGraph.CreateDependencyGraph(dependencies)
	if err != nil {
		fmt.Println("Graph creation failed", err)
		return
	}
	fmt.Println("JSON file/s for graph created successfully")
}
