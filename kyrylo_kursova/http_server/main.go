package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const logDir = "logs"

func ensureLogDir() {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.Mkdir(logDir, os.ModePerm)
	}
}

func getLogFileName(t time.Time) string {
	return filepath.Join(logDir, t.Format("2006-01-02")+".log")
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	var targetDate time.Time
	var err error

	if date == "" {
		targetDate = time.Now()
	} else {
		targetDate, err = time.Parse("2006-01-02", date)
		if err != nil {
			w.Write([]byte("Невірний формат дати. Використовуйте ?date=YYYY-MM-DD\n"))
			return
		}
	}

	filePath := getLogFileName(targetDate)

	data, err := os.ReadFile(filePath)
	if err != nil {
		w.Write([]byte("Лог-файл за цю дату не знайдено або порожній\n"))
		return
	}

	w.Write(data)
}

type LogEntry struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var entry LogEntry

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Очікується JSON у тілі запиту з Content-Type: application/json\n"))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Помилка читання запиту\n"))
		return
	}

	err = json.Unmarshal(body, &entry)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Невірний формат JSON\n"))
		return
	}

	if entry.Username == "" || entry.Message == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Поля 'username' і 'message' є обов’язковими\n"))
		return
	}

	timeNow := time.Now()
	logPath := getLogFileName(timeNow)
	logLine := fmt.Sprintf("[%s] %s: %s\n", timeNow.Format("15:04:05"), entry.Username, entry.Message)

	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		w.Write([]byte("Помилка запису в лог\n"))
		return
	}
	defer file.Close()

	file.WriteString(logLine)
	w.Write([]byte("Лог збережено у файл " + filepath.Base(logPath) + "\n"))
}

func handleDelete(w http.ResponseWriter) {
	files, err := os.ReadDir(logDir)
	if err != nil {
		w.Write([]byte("Помилка доступу до логів\n"))
		return
	}
	for _, f := range files {
		os.Remove(filepath.Join(logDir, f.Name()))
	}
	w.Write([]byte("Усі логи очищено\n"))
}

func main() {
	ensureLogDir()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleGet(w, r)
		case "POST":
			handlePost(w, r)
		case "DELETE":
			handleDelete(w)
		default:
			w.Write([]byte("Метод не підтримується\n"))
		}
	})

	fmt.Println("Сервер запущено на http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
