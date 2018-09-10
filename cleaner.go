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
		return errors.New("Repository is needed")
	}

	if len(tags) == 0 {
		return errors.New("Tags are needed")
	}

	log.Println("call get tags")
	err := GetTags(repository, tags)

	if err != nil {
		log.Println("returned error: " + err.Error())
	}

	return nil
}

//GetTags var func
var GetTags = func(repository string, tags []string) error {
	log.Println("get tags")
	out, err := ExecuteCommandFunction(
		"az",
		"acr",
		"repository",
		"show-tags",
		"-n",
		"segurosfalabella",
		"--repository",
		repository)

	if err != nil {
		return errors.New("Error returned from execute command function")
	}

	var response []string
	errMarshall := UnmarshalFunction([]byte(out), &response)

	if errMarshall != nil {
		return errMarshall
	}

	log.Println("iterate")
	for _, tag := range response {
		if !inArray(tag, tags) && !strings.Contains(tag, "c-") {
			log.Println("going to delete " + tag + "\xE2\x9C\x94")
			DeleteUnusedTags(tag, repository)
		}
	}

	return nil
}

//DeleteUnusedTags var function
var DeleteUnusedTags = func(tag string, repository string) {
	if !strings.Contains(tag, "latest") {
		_, err := ExecuteCommandFunction(
			"az",
			"acr",
			"repository",
			"delete",
			"-n",
			"segurosfalabella",
			"--image",
			repository+":"+tag,
			"--yes")

		log.Println("deleting " + tag + "\xE2\x9C\x94 \xE2\x9C\x94")

		//err := cmd.Run()
		if err != nil {
			panic(err)
		}
	}
}

func inArray(val string, array []string) bool {
	for pos := 0; pos < len(array); pos++ {
		if array[pos] == val {
			return true
		}
	}
	return false
}

//ExecuteCommandFunction function var
var ExecuteCommandFunction = func(params ...string) ([]byte, error) {
	if len(params) == 0 {
		return nil, errors.New("Parameters are needed")
	}

	param := params[0]
	rest := append(params[:0], params[0+1:]...)
	log.Println("executing command")
	out, err := exec.Command(param, rest...).Output()
	log.Println("waiting for response")

	if err != nil {
		log.Fatal(err)
	}

	return out, err
}

//UnmarshalFunction var function
var UnmarshalFunction = func(out []byte, response interface{}) error {
	err := json.Unmarshal(out, response)

	log.Println("unmarshaling out")
	if err != nil {
		return errors.New("something goes wrong running json unmarshal")
	}

	return nil
}
