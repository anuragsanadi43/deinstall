package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Defile struct {
	App          string
	Description  string
	Dependencies struct {
		Deb string
	}
}

func main() {
	// create an object of Defile
	dependenciesInstaller := Defile{}

	// read Defile.yml
	contents, err := ioutil.ReadFile("Defile.yml")

	if err != nil {
		fmt.Println("No Defile exists. Please create one....")
	}

	// place contents of the file into dependenciesInstaller object
	error := yaml.Unmarshal([]byte(contents), &dependenciesInstaller)
	if error != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Deinstaller installing dependencies for %s\n", dependenciesInstaller.App)

	//fmt.Println(runtime.GOOS)
	if runtime.GOOS == "linux" {
		var command string = fmt.Sprintf("sudo apt install %s\n", dependenciesInstaller.Dependencies.Deb)
		cmd := exec.Command(command)
		fmt.Println(cmd)
	} else {
		fmt.Println("This command only works on Debian based systems")
	}

}
