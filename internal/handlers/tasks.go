package handlers

import (
	"fmt"
	"log"
	"time"
	"strings"
	"strconv"
	"encoding/json"
	"go18_api/internal/models"
	"go18_api/internal/storage"
	"net/http"
)

type Handler struct{ Store storage.Storage }

func New(s storage.Storage) *Handler { return &Handler{Store: s} }

func sendError(w *http.ResponseWriter, status int, message string) {
	err := models.Error{Message: message}

	(*w).WriteHeader(status)
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(err)
}

func writeLog(r *http.Request) {
	// t := time.Now().Format("2006-01-02 15:04:05")
	log.Printf("%s %s\n", r.Method, r.URL.Path)
}

func sendOK(w *http.ResponseWriter, status int, object any) {
	(*w).WriteHeader(status)
	(*w).Header().Set("Content-Type", "application/json")

	if (object != nil) {
		json.NewEncoder(*w).Encode(object)
	}
}

// /tasks (GET, POST)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	// TODO: реализуйте разбор метода, JSON, коды статусов, валидацию
	writeLog(r);

	if (r.Method != http.MethodGet) && (r.Method != http.MethodPost) {
		// http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		sendError(&w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	sendOK(&w, http.StatusOK, h.Store.List())
}

// /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	writeLog(r);

	// TODO: извлечение id, маршрутизация по методу, ошибки
	parts := strings.Split(r.URL.Path, "/")

	id, err := strconv.Atoi(parts[len(parts) - 1])

	if (err != nil) && ((r.Method == http.MethodGet) || (r.Method == http.MethodDelete)) {
		fmt.Println("Некорректный ID - " + parts[len(parts) - 1])
		sendError(&w, http.StatusBadRequest, "Некорректный ID - " + parts[len(parts) - 1])
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, is_exist := h.Store.Get(id)

		if is_exist {
			fmt.Println("GET - OK")
			sendOK(&w, http.StatusOK, task)
		} else {
			fmt.Println("GET - Not Found")
			sendError(&w, http.StatusNotFound, fmt.Sprintf("Задача с таким ID не найдена - %d", id))
		}

	case http.MethodDelete:
		err := h.Store.Delete(id)

		if err == nil {
			sendOK(&w, http.StatusOK, nil)
		} else {
			sendError(&w, http.StatusInternalServerError, "Ошибка при удалении: " + err.Error())
		}

	case http.MethodPut:
		var task models.Task
		var task2 models.Task

		if err2 := json.NewDecoder(r.Body).Decode(&task); err2 != nil {
			sendError(&w, http.StatusBadRequest, "Неверный формат данных")
			return
		} else {
			if err != nil {
				task.CreatedAt = time.Now().Format(time.RFC3339)
				task2, err = h.Store.Create(task)
			} else {
				task2, err = h.Store.Update(id, task)
			}

			if err == nil {
				sendOK(&w, http.StatusOK, task2)
			} else {
				sendError(&w, http.StatusInternalServerError, "Ошибка при обновлении: " + err.Error())
			}
		}

	default:
		sendError(&w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}
