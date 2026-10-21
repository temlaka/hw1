package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		os.Exit(1)
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка чтения:", err)
		os.Exit(1)
	}

	if response == "OK\n" {
		fmt.Println("Успех: получен корректный ответ OK")
	} else {
		fmt.Printf("Ошибка: ожидался 'OK\\n', получен %q\n", response)
		os.Exit(1)
	}
}
