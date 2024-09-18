package http

import (
    "net/http"

    "encoding/json"
    "github.com/kperilla/habitapi/habitapi"
)

func (h *Handler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {
    var user habitapi.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    _, id, err := h.UserService.CreateUser(user.Name)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    WriteJSON(w, http.StatusCreated, id)
}

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")
    user, err := h.UserService.User(id)
    if err != nil {
        WriteJSON(w, http.StatusBadRequest, err)
    }
    WriteJSON(w, http.StatusOK, user)
}
