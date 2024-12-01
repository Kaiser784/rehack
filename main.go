package main

import (
    "fmt"
    "os"
    "os/exec"
	"bufio"
    _"strings"
)

var attack = make([]string, 0)

func main() {
    
	read_files()
	install_attack()
	// remove()
}

func read_files() {
	file, _ := os.Open("./setup-strings/attack.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	
    for sc.Scan() {
		attack = append(attack, sc.Text())
	}
}

func install_attack() {

	// Print the list of tools.
    for _, tool := range attack {
        
        // command := fmt.Sprintf("sudo apt install %s", tool)
        fmt.Println("Installing", tool, "...")
        // os.system(command)
        
        // Use the exec.Command function to run the command.
        out, err := exec.Command("/bin/bash", "./setup-scripts/attack.sh").Output()
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println(string(out))
    }
}

func remove() {
	for _, tool := range attack {
        
        // command := fmt.Sprintf("sudo apt install %s", tool)
        fmt.Println("Purging", tool)
        // os.system(command)
        
        // Use the exec.Command function to run the command.
        out, err := exec.Command("sudo", "apt-get", "purge", tool, "-y").Output()
        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println(string(out))
    }
}