package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		os.Exit(404)
	}

	// place contents of the file into dependenciesInstaller object
	error := yaml.Unmarshal([]byte(contents), &dependenciesInstaller)
	if error != nil {
		log.Fatalf("error: %v", err)
	}

	packages := ""
	helpPage := "usage: deinstall [options]\n    options:\n      --help display this help message\n      --dist=deb|arch|redhat|fedora|freebsd install debpendencies depending on your linux distribution"

	if len(os.Args) != 2 {
		// proper arguments not given by user
		fmt.Println(helpPage)
	} else {
		// len(os.Args) == 2 -- correct command
		if string(os.Args[1]) == "--help" {
			fmt.Println(helpPage)
			os.Exit(404)
		} else {
			if len(string(os.Args[1])) > 6 && string(os.Args[1])[:7] == "--dist=" {
				if (os.Args[1])[7:] == "deb" {
					packages = dependenciesInstaller.Dependencies.Deb
					fmt.Printf("Deinstaller installing dependencies for %s\n", dependenciesInstaller.App)
				} else {
					fmt.Printf("'%s' package manager not available\n", (os.Args[1])[7:])
					fmt.Println("Run 'deinstall --help' for help page")
				}
			} else {
				fmt.Println(helpPage)
			}
		}
	}

	if runtime.GOOS == "linux" {
		if packages != "" {
			command := fmt.Sprintf("sudo apt-get install %s -y", packages)
			c := exec.Command("/bin/bash", "-c", command)
			//fmt.Println(c)

			err := c.Run()
			if err != nil {
				fmt.Println("Error in installation: ", err)
			}
		}
	} else {
		fmt.Println("This command only works on Debian based systems")
	}

}
