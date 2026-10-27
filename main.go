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

	basicPorts := map[string]int{
		"SMTP":     25,
		"POP3":     110,
		"IMAP":     143,
		"DHCP":     67,
		"TFTP":     69,
		"NTP":      123,
		"SNMP":     161,
		"HTTP-Alt": 8080,
	}

	mediumPorts := map[string]int{
		"MySQL":           3306,
		"PostgreSQL":      5432,
		"RDP":             3389,
		"SMB":             445,
		"LDAP":            389,
		"VNC":             5900,
		"SMTPS":           465,
		"SMTP Submission": 587,
		"IMAPS":           993,
		"POP3S":           995,
	}

	welcome()

	for {
		// Level 1
		fmt.Println("Everyone should know these ports")
		if !Game(easyPorts, &lifes, reader) {
			if gameOver(reader) {
				return
			}
		}
		// Level 2
		fmt.Println("Typical ports for networking and mail")
		if !Game(basicPorts, &lifes, reader) {
			if gameOver(reader) {
				return
			}
		}
		// Level 3
		fmt.Println("Databases, remote access and secure mail")
		if !Game(mediumPorts, &lifes, reader) {
			if gameOver(reader) {
				return
			}
		}
		// Level 4

	}

}

func welcome() {
	cleanScrean() // Función para limpiar pantalla
	fmt.Printf("Welcome to Guess The Port!!!\nA game to learn the defualt port by its service!\n")
	fmt.Printf("Press any key to continue\n")
	getch()
	cleanScrean()
}

func Game(Ports map[string]int, lifes *int, reader *bufio.Reader) bool {
	// fmt.Println("Every one should know these ports")
	for service, port := range Ports {

		fmt.Printf("Lifes: %d\n", *lifes)
		fmt.Printf("What is the defualt port of %s? ", service)

		// Read the anwser
		input, _ := reader.ReadString('\n')
		guessedPort, err := strconv.Atoi(strings.TrimSpace(input))

		if err == nil && guessedPort == port {
			fmt.Println("You got it right! ...now press a key")
		} else {
			*lifes--
			fmt.Printf("Wrong port It was %d, press a key\n", port)
		}
		if *lifes <= 0 {
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
	if answer == "y" {
		return false
	}
	return true
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
