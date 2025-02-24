package health

import (
	"io"
	"net/http"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Ok")
}
