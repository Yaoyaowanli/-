package objects01

import "net/http"

func Handler01(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	if m == http.MethodPut {
		put01(w, r)
		return
	}
	if m == http.MethodGet {
		get01(w, r)
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
