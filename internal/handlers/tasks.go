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

func sendError(w http.ResponseWriter, status int, message string) {
	err := models.Error{Message: message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func writeLog(r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.Path)
}

func sendOK(w http.ResponseWriter, status int, object any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if (object != nil) {
		json.NewEncoder(w).Encode(object)
	}
}

// /tasks (GET, POST)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	writeLog(r);

	switch r.Method {
	case http.MethodGet:
		sendOK(w, http.StatusOK, h.Store.List())

	case http.MethodPost:
		var task models.Task
		var task2 models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			sendError(w, http.StatusBadRequest, "Неверный формат данных")
		} else {
			task.CreatedAt = time.Now().Format(time.RFC3339)
			task2, err = h.Store.Create(task)

			if err == nil {
				sendOK(w, http.StatusCreated, task2)
			} else {
				sendError(w, http.StatusInternalServerError, "Ошибка при создании задачи: " + err.Error())
			}
		}

	default:
		sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	writeLog(r);

	parts := strings.Split(r.URL.Path, "/")
	if (len(parts) != 3) {
		log.Println("Некорректный путь - " + r.URL.Path)
		sendError(w, http.StatusBadRequest, "Некорректный путь - " + r.URL.Path)
		return
	}

	id, err := strconv.Atoi(parts[len(parts) - 1])

	if (err != nil) {
		log.Println("Некорректный ID - " + parts[len(parts) - 1])
		sendError(w, http.StatusBadRequest, "Некорректный ID - " + parts[len(parts) - 1])
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, is_exist := h.Store.Get(id)

		if is_exist {
			log.Println("GET - OK")
			sendOK(w, http.StatusOK, task)
		} else {
			log.Println("GET - Not Found")
			sendError(w, http.StatusNotFound, fmt.Sprintf("Задача с таким ID не найдена - %d", id))
		}

	case http.MethodDelete:
		err := h.Store.Delete(id)

		if err == nil {
			sendOK(w, http.StatusNoContent, nil)
		} else {
			sendError(w, http.StatusInternalServerError, "Ошибка при удалении: " + err.Error())
		}

	case http.MethodPut:
		var task models.Task
		var task2 models.Task

		if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
			sendError(w, http.StatusBadRequest, "Неверный формат данных")
			return
		} else {
			task2, err = h.Store.Update(id, task)

			if err == nil {
				sendOK(w, http.StatusOK, task2)
			} else {
				sendOK(w, http.StatusNotFound, nil)
			}
		}

	default:
		sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}
