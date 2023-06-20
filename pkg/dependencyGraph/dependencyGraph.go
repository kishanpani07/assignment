package dependencyGraph

import (
	"encoding/json"
	"lineaje-assignment/pkg/formatString"
	"os"
	"strings"
)

var (
	isRoot        map[string]bool     = map[string]bool{}
	dependencyMap map[string][]string = map[string][]string{}
)

func CreateDependencyGraph(rawDependencies []formatString.Dependency) error {
	// Iterate through dependencies and fill the maps
	for _, rawDependency := range rawDependencies {
		isRoot[rawDependency.Parent] = true
		isRoot[rawDependency.Child] = true
		dependencyMap[rawDependency.Parent] = []string{}
		dependencyMap[rawDependency.Child] = []string{}
	}

	// Mark all child packages as false and fill dependencyMap
	for _, rawDependency := range rawDependencies {
		isRoot[rawDependency.Child] = false
		tempChildren := dependencyMap[rawDependency.Parent]
		tempChildren = append(tempChildren, rawDependency.Child)
		dependencyMap[rawDependency.Parent] = tempChildren
	}

	// Create JSON for every root package
	for parent, isPackageRoot := range isRoot {
		if isPackageRoot {
			createGraph(parent)
		}
	}

	return nil
}

func createGraph(root string) error {
	// Recursively populate the dependency graph
	var isTraversed map[string]bool = map[string]bool{} // To prevent dependency cycle
	artifactTree := populateGraph(root, isTraversed)
	createJSON(artifactTree)
	return nil
}

func populateGraph(rawArtifact string, isTraversed map[string]bool) Artifact {
	isTraversed[rawArtifact] = true
	var currentArtifact Artifact = Artifact{
		Name:         rawArtifact,
		Dependencies: []*Artifact{},
	}
	if len(dependencyMap[rawArtifact]) != 0 {
		for _, childArtifact := range dependencyMap[rawArtifact] {
			if !isTraversed[childArtifact] {
				nextArtifact := populateGraph(childArtifact, isTraversed)
				tempDependencies := currentArtifact.Dependencies
				tempDependencies = append(tempDependencies, &nextArtifact)
				currentArtifact.Dependencies = tempDependencies
			}
		}
	}
	return currentArtifact
}

func createJSON(artifact Artifact) error {
	file, err := os.OpenFile("./data/"+strings.ReplaceAll(artifact.Name, "/", "")+".json", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	artifactByte, err := json.Marshal(artifact)
	if err != nil {
		return err
	}
	_, err = file.Write(artifactByte)
	if err != nil {
		return err
	}
	return nil
}
