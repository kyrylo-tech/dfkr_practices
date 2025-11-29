package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hashText(input, algo string) (string, error) {
	data := []byte(input)

	switch algo {
	case "md5":
		sum := md5.Sum(data)
		return hex.EncodeToString(sum[:]), nil
	case "sha1":
		sum := sha1.Sum(data)
		return hex.EncodeToString(sum[:]), nil
	case "sha256":
		sum := sha256.Sum256(data)
		return hex.EncodeToString(sum[:]), nil
	case "sha512":
		sum := sha512.Sum512(data)
		return hex.EncodeToString(sum[:]), nil
	default:
		return "", fmt.Errorf("невідомий алгоритм: %s", algo)
	}
}

func main() {
	var input string
	var algorithm string

	if len(os.Args) >= 3 {
		input = os.Args[1]
		algorithm = os.Args[2]
	} else {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Введіть текст: ")
		input, _ = reader.ReadString('\n')

		fmt.Print("Введіть алгоритм (md5, sha1, sha256, sha512): ")
		algorithm, _ = reader.ReadString('\n')
	}

	input = strings.TrimSpace(input)
	algorithm = strings.TrimSpace(algorithm)

	hashValue, err := hashText(input, algorithm)

	if err != nil {
		fmt.Println("Помилка:", err)
		return
	}

	fmt.Println("Алгоритм:", algorithm)
	fmt.Println("Вхідні дані:", input)
	fmt.Println("Хеш:", hashValue)
}
