package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// conn, err := net.Dial("tcp", "10.10.1.226:8080")
	ClearTerminal()

	// Connect Server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error connect to server: ", err)
		return
	}
	fmt.Println("Conected to server")
	defer conn.Close()

	// Printing All massage
	go ReceiveResponse(conn)

	reader := bufio.NewReader(os.Stdin)

	// Asking name of the user
	fmt.Print("Enter your username: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error entering username")
		return
	}

	// Send username to server
	_, err = conn.Write([]byte(name))
	if err != nil {
		log.Println("Error sending username")
		return
	}

	ClearTerminal()
	fmt.Print("Enter massage: ")
	for {
		massage, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input: ", err)
			return
		}

		_, err = conn.Write([]byte(massage))
		if err != nil {
			log.Println("Error send message to server: ", err)
			return
		}

	}
}

func ReceiveResponse(conn net.Conn) {
	for {
		response, err := bufio.NewReader(conn).ReadString(byte(3))
		if err != nil {
			log.Println("Error receive response from server: ", err)
			return
		}
		ClearTerminal()

		fmt.Print(response)
	}
}

func ClearTerminal() {
	cmd := exec.Command("clear") // For Unix-like systems
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
