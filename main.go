package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	lifes := 5

	easyPorts := map[string]int{
		"HTTP":   80,
		"HTTPS":  443,
		"SSH":    22,
		"FTP":    21,
		"Telnet": 23,
		"DNS":    53,
	}

	welcome()

	for {
		if !easyPortsLevel(easyPorts, lifes, reader) {
			if gameOver(reader) {
				return
			}
		} else {

		}
	}

}

func welcome() {
	cleanScrean() // Función para limpiar pantalla
	fmt.Printf("Welcome to Guess The Port!!!\nA game to learn the defualt port by its service!\n")
	fmt.Printf("Press any key to continue\n")
	getch()
}

func easyPortsLevel(easyPorts map[string]int, lifes int, reader *bufio.Reader) bool {
	cleanScrean()
	fmt.Println("Every one should know these ports")
	for service, port := range easyPorts {

		fmt.Printf("Lifes: %d\n", lifes)
		fmt.Printf("What is the defualt port of %s? ", service)

		// Read the anwser
		input, _ := reader.ReadString('\n')
		guessedPort, err := strconv.Atoi(strings.TrimSpace(input))

		if err == nil && guessedPort == port {
			fmt.Println("You got it right! ...now press a key")
		} else {
			lifes--
			fmt.Printf("Wrong port It was %d, press a key", port)
		}
		if lifes <= 0 {
			return false
		}
		getch()
		cleanScrean()
	}
	return true
}

func gameOver(reader *bufio.Reader) bool {
	fmt.Println("Wanna play again from the start? [y/n]")
	input, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(input)
	if answer == "n" {
		return true
	}
	return false
}

func getch() {
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	_, _, err = keyboard.GetKey()
	if err != nil {
		panic(err)
	}
}

func cleanScrean() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" { // Si el sistema operativo es Windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Si es tipo *nix
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
