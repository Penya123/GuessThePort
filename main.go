package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/eiannone/keyboard"
)

func main() {
	welcome()

}

func welcome() {
	cleanScrean() // Función para limpiar pantalla
	fmt.Printf("Welcome to Guess The Port!!!\nA game to learn the defualt port by its service!\n")
	fmt.Printf("Press any key to continue\n")

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
