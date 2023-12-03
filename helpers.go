package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetBrazilCurrentTimeHelper() (string, error) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		return "", fmt.Errorf("An error occured getting Brazilian time: %v", err.Error())
	}

	currentTime := time.Now().In(loc)
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	return formattedTime, nil
}

func WriteJsonHelper(w http.ResponseWriter, status int, v any) error {
	if w == nil {
		return fmt.Errorf("Error validanting data or data is null")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}
