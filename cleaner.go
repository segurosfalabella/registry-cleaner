package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	//flag.StringVar(&repo, "repo", "", "The Azure Cloud Registry repository")
	flag.Parse()
	args := flag.Args()
	fmt.Println("repository:", repo)
	fmt.Println("arguments: ", args)

	for index := 0; index < len(args); index++ {
		fmt.Println(args[index])
	}
	///CleanRegistry()
}

//CleanRegistry function
func CleanRegistry(repo string, tags []string) error {
	if repo == "" {
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
