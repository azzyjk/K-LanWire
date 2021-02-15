package src

import (
	"database/sql"
	"io/ioutil"
	"net/http"
)

func readFromRequest(request *http.Request) (string, error) {

	result, err := ioutil.ReadAll(request.Body)
	return string(result), err
}

func nullStringToString(inputString sql.NullString) string {

	if inputString.Valid {
		return inputString.String
	} else {
		return ""
	}
}
