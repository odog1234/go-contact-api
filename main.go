package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Contact struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Number string `json:"number"`
}

var contacts = []Contact{}

func main() {
	http.HandleFunc("/contact", handleContacts)
	http.HandleFunc("/contact/", updateContact)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func createContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newConatct Contact
	err := json.NewDecoder(r.Body).Decode(&newConatct)
	if err != nil {
		http.Error(w, "JSON not found", http.StatusBadRequest)
		return
	}

	newConatct.ID = len(contacts) + 1
	contacts = append(contacts, newConatct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func handleContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getContacts(w, r)
	} else if r.Method == http.MethodPost {
		createContacts(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/contact/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			// Remove contact by slicing out the matching index
			contacts = append(contacts[:i], contacts[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Contact with ID %d deleted successfully\n", id)
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		// If it's not PUT, try deleteContact for DELETE method
		if r.Method == http.MethodDelete {
			deleteContact(w, r)
			return
		}
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/contact/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid contact ID", http.StatusBadRequest)
		return
	}

	var updatedData Contact
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, contact := range contacts {
		if contact.ID == id {
			contacts[i].Name = updatedData.Name
			contacts[i].Email = updatedData.Email
			contacts[i].Number = updatedData.Number

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(contacts[i])
			return
		}
	}

	http.Error(w, "Contact not found", http.StatusNotFound)
}
