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
	ClearTerminal()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println("Error setting up listener: ", err)
		return
	}
	fmt.Println("Server is listening on port 8080")
	clients := &map[string]net.Conn{}
	history := "Enter massage: " + string(byte(3))
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error acception connetion: ", err)
			continue
		}
		go handleConnetion(conn, clients, &history)
	}
}

func handleConnetion(conn net.Conn, clients *map[string]net.Conn, history *string) {
	defer conn.Close()
	fmt.Println("Conneted client: ", conn.RemoteAddr().String())

	var name string
	defer DeleteUser(clients, name)

	// Taking username
	reader := bufio.NewReader(conn)
	name, err := reader.ReadString('\n')
	name = name[:len(name)-1] + ": "
	if err != nil {
		log.Println("Error Receiving username: ", err)
		return
	}
	(*clients)[name] = conn

	for {
		massage, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading massage: ", err)
			return
		}
		if len(massage) == 1 {
			continue
		}

		*history = (*history)[:len(*history)-16]
		*history += name + massage + "Enter massage: " + string(byte(3))
		for _, conn1 := range *clients {
			fmt.Println("Received massage: ", massage)
			_, err = conn1.Write([]byte(*history))
			if err != nil {
				log.Println("Error write message to connetion: ", err)
				return
			}
		}
	}
}

func DeleteUser(clients *map[string]net.Conn, name string) {
	delete(*clients, name)
}

func ClearTerminal() {
	cmd := exec.Command("clear") // For Unix-like systems
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
