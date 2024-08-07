package definition

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func parseYamlFile(p string, v interface{}) (err error) {

	body, err := os.ReadFile(p)
	if err != nil {
		log.Println(err)
		return
	}

	err = yaml.Unmarshal(body, v)
	if err != nil {
		log.Println(err)
		return
	}

	return
}

func parseYaml(in []byte, v interface{}) (err error) {
	err = yaml.Unmarshal(in, v)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
