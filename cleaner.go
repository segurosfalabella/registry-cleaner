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
	log.Println("clean registry")

	if repository == "" {
		return errors.New("repository is needed!")
	}

	if len(tags) == 0 {
		return errors.New("tags are needed!")
	}

	log.Println("call get tags")
	getTags(repository, tags)
	return nil
}

func getTags(repository string, tags []string) {
	log.Println("get tags")
	var resp []string
	out, err := exec.Command(
		"az",
		"acr",
		"repository",
		"show-tags",
		"-n",
		"segurosfalabella",
		"--repository",
		repository).Output()

	log.Println("waiting for response")
	if err != nil {
		log.Fatal(err)
	}

	errMarshall := json.Unmarshal([]byte(out), &resp)

	if errMarshall != nil {
		panic(errMarshall)
	}

	log.Println("iterate")
	for _, tag := range resp {
		if !in_array(tag, tags) && !strings.Contains(tag, "c-") {
			log.Println("going to delete " + tag)
			deleteUnusedTags(tag, repository)
		}
	}
}

func deleteUnusedTags(tag string, repository string) {
	if !strings.Contains(tag, "latest") {
		cmd := exec.Command(
			"az",
			"acr",
			"repository",
			"delete",
			"-n",
			"segurosfalabella",
			"--image",
			repository+":"+tag,
			"--yes")

		log.Println("deleting " + tag)

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func in_array(val string, array []string) bool {
	for pos := 0; pos < len(array); pos++ {
		if array[pos] == val {
			return true
		}
	}
	return false
}

// func getManifests(repository string) {
// 	var resp []interface{}
// 	out, err := exec.Command(
// 		"az",
// 		"acr",
// 		"repository",
// 		"show-manifests",
// 		"-n",
// 		"segurosfalabella",
// 		"--repository",
// 		repository,
// 	).Output()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := json.Unmarshal([]byte(out), &resp); err != nil {
// 		panic(err)
// 	}

// 	for _, m := range resp {
// 		tag := m.(map[string]interface{})["tags"]
// 		digest := m.(map[string]interface{})["digest"].(string)
// 		if tag == nil {
// 			// fmt.Println(m.(map[string]interface{})["digest"])
// 			deleteNilTags(repository, digest)
// 		}
// 	}
// }

// func deleteNilTags(repository string, digest string) {
// 	image := (repository + "@" + digest)
// 	cmd := exec.Command(
// 		"az",
// 		"acr",
// 		"repository",
// 		"delete",
// 		"-n",
// 		"segurosfalabella",
// 		"--image",
// 		image,
// 		"-y",
// 	)
// 	err := cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
