package getaoe2civ

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

func GetCiv(configFile string) {
	var confStruct config
	confStruct.readInYaml(configFile)

	rand.Seed(time.Now().Unix())

	for {
		randomIndex := rand.Intn(len(confStruct.Civs))
		civ := confStruct.Civs[randomIndex]
		fmt.Printf("Your random civ is: %s\n", civ)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("\nWould you like a new civ? (Y/n)")
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if !strings.EqualFold(text, "y") {
			fmt.Println("Good luck out there!")
			break
		}
	}

}

type config struct {
	Civs []string `yaml:"civs"`
}

func (confStruct *config) readInYaml(configFile string) *config {
	yamlFile, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatalln("No yaml file passed")
	}
	err = yaml.Unmarshal(yamlFile, &confStruct)
	if err != nil {
		log.Fatalf("Error reading in Yaml file")
	}
	return confStruct
}
