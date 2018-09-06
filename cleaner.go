package main

import (
	"encoding/json"
	"errors"
	"flag"
	"log"
	"os/exec"
	"strings"
)

func main() {
	var repository string
	flag.StringVar(&repository, "repository", "", "The Azure Cloud Registry repository")
	flag.Parse()
	args := flag.Args()
	CleanRegistry(repository, args)
}

//CleanRegistry function
func CleanRegistry(repository string, tags []string) error {
	if repository == "" {
		return errors.New("repository is needed!")
	}

	if len(tags) == 0 {
		return errors.New("tags are needed!")
	}

	//getTags(repo)
	return nil
}

func getTags(repo string) {
	var resp []string
	out, err := exec.Command(
		"az",
		"acr",
		"repository",
		"show-tags",
		"-n",
		"segurosfalabella",
		"--repository",
		repo).Output()

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		panic(err)
	}

	for _, tag := range resp {
		deleteUnusedTags(tag, repo)
	}
}

// deletes everything but latest tag
func deleteUnusedTags(tag string, repo string) {
	if !strings.Contains(tag, "latest") {
		cmd := exec.Command(
			"az",
			"acr",
			"repository",
			"delete",
			"-n",
			"segurosfalabella",
			"--image",
			repo+":"+tag,
			"--yes")

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// use the following to delete nil tags from manifest
func getManifests(repo string) {
	var resp []interface{}
	out, err := exec.Command(
		"az",
		"acr",
		"repository",
		"show-manifests",
		"-n",
		"segurosfalabella",
		"--repository",
		repo,
	).Output()

	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal([]byte(out), &resp); err != nil {
		panic(err)
	}

	for _, m := range resp {
		tag := m.(map[string]interface{})["tags"]
		digest := m.(map[string]interface{})["digest"].(string)
		if tag == nil {
			// fmt.Println(m.(map[string]interface{})["digest"])
			deleteNilTags(repo, digest)
		}
	}
}

func deleteNilTags(repo string, digest string) {
	image := (repo + "@" + digest)
	cmd := exec.Command(
		"az",
		"acr",
		"repository",
		"delete",
		"-n",
		"segurosfalabella",
		"--image",
		image,
		"-y",
	)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
