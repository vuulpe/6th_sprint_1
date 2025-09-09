package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Проверяем метод
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Ошибка парсинга формы: "+err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Ошибка получения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла: "+err.Error(), http.StatusInternalServerError)
		return
	}

	input := string(fileData)
	result, err := service.Convert(input)
	if err != nil {
		http.Error(w, "Ошибка конвертации: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := os.MkdirAll("uploads", 0755); err != nil {
		http.Error(w, "Ошибка создания папки: "+err.Error(), http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("20060102-150405")
	ext := filepath.Ext(header.Filename)
	newFilename := fmt.Sprintf("uploads/result_%s%s", timestamp, ext)

	outputFile, err := os.Create(newFilename)
	if err != nil {
		http.Error(w, "Ошибка создания файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(result); err != nil {
		http.Error(w, "Ошибка записи в файл: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(result))
}
