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

	levels := []map[string]int{
		{ //Easy ports
			"HTTP":   80,
			"HTTPS":  443,
			"SSH":    22,
			"FTP":    21,
			"Telnet": 23,
			"DNS":    53,
		}, { //Basic ports
			"SMTP":     25,
			"POP3":     110,
			"IMAP":     143,
			"DHCP":     67,
			"TFTP":     69,
			"NTP":      123,
			"SNMP":     161,
			"HTTP-Alt": 8080,
		}, { //Medium ports
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
		}, { //Hard ports
			"MSSQL":          1433,
			"Oracle DB":      1521,
			"MongoDB":        27017,
			"Redis":          6379,
			"Memcached":      11211,
			"Elasticsearch":  9200,
			"Kibana":         5601,
			"RabbitMQ":       5672,
			"Docker API":     2375,
			"Kubernetes API": 6443,
			"etcd":           2379,
			"Cassandra CQL":  9042,
		}, { //Expert port
			"TACACS+":         49,
			"Kerberos":        88,
			"RPC":             135,
			"NetBIOS Session": 139,
			"BGP":             179,
			"Syslog":          514,
			"LDAPS":           636,
			"Rsync":           873,
			"RADIUS":          1812,
			"MQTT":            1883,
			"NFS":             2049,
			"Zookeeper":       2181,
			"SIP":             5060,
			"WinRM":           5985,
			"Kafka":           9092,
		},
	}

	welcome()

	// Actual game
	for i := 0; i < len(levels); {
		fmt.Println("Level: ", i+1)
		if !Game(levels[i], &lifes, reader) {
			if gameOver(reader, &lifes, &i) {
				return
			}
		} else {
			i++
		}
	}
}

func welcome() {
	cleanScreen() // Función para limpiar pantalla
	fmt.Printf("Welcome to Guess The Port!!!\nA game to learn the defualt port by its service!\n")
	fmt.Printf("Press any key to continue\n")
	getch()
	cleanScreen()
}

// Function that ask the port given the service
// returns false if lifes <= 0, otherways returns true
func Game(Ports map[string]int, lifes *int, reader *bufio.Reader) bool {
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

		getch()
		cleanScreen()

		if *lifes <= 0 {
			return false
		}
	}
	return true
}

// Function that's triggered when lifes >= 0
func gameOver(reader *bufio.Reader, lifes *int, i *int) bool {
	fmt.Println("Wanna play again from the start? [y/n]")
	input, _ := reader.ReadString('\n')
	answer := strings.TrimSpace(input)
	if answer == "y" {
		*lifes = 5
		*i = 0
		return false
	}
	return true
}

// Function to wait for user's input
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

// Function to clean the screen
func cleanScreen() {
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" { // Si el sistema operativo es Windows
		cmd = exec.Command("cmd", "/c", "cls")
	} else { // Si es tipo *nix
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
