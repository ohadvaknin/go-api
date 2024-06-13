package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/ohadvaknin/go-api/storage"
	"github.com/ohadvaknin/go-api/models"
)

func CreateTask(w http.ResponseWriter, r *http.Request){
	var task models.Task = json.Unmarshal()
}