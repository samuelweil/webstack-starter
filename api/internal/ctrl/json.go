package ctrl

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, output interface{}) {
	resultBytes, err := json.Marshal(output)
	if err != nil {
		handleError(w, err)
		return
	}

	_, err = w.Write(resultBytes)

	if err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, e error) {
	// There's nothing we can do if there's an error trying to write the buffer so
	// we log and return
	log.Println(e)
	http.Error(w, "Unknown Server Error", http.StatusInternalServerError)
}
