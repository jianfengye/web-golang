package util

import(
	"net/http"
)

func ShowAlert(w http.ResponseWriter, r *http.Request, reason string) {
    w.Write([]byte("<script>alert('" + reason + "')</script>"))
}