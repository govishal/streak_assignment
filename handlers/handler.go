package handlers

import (
	"encoding/json"
	"net/http"
	"streak_assignment/logic"
	"streak_assignment/models"
	"streak_assignment/utils"
)

// FindPairsHandler handles the POST /find-pairs request
func FindPairsHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody models.RequestBody

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil || !utils.ValidateInput(reqBody) {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	solutions := logic.FindPairs(reqBody.Numbers, reqBody.Target)

	respBody := models.ResponseBody{Solutions: solutions}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}
