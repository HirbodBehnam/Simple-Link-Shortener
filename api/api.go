package api

import (
	"UrlShortener/database"
	"UrlShortener/util"
	"database/sql"
	"log"
	"net/http"
)

// Endpoint is the main entry point of hour site
func Endpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getLink(w, r)
	case "DELETE":
		deleteLink(w, r)
	case "PUT":
		putLink(w, r)
	}
}

// getLink gets the link from database and redirects user
func getLink(w http.ResponseWriter, r *http.Request) {
	link, err := database.MainDatabase.GetLink(r.Context(), r.URL.Path[1:]) // skip the first /
	if err == sql.ErrNoRows {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		log.Printf("cannot get link from wallet: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, link, http.StatusPermanentRedirect)
}

// putLink puts a new link in database
func putLink(w http.ResponseWriter, r *http.Request) {
	// Check token
	if !authorizeHeader(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// Generate token
	token := r.URL.Query().Get("token")
	if token == "" {
		token = util.GenerateLinkToken()
	}
	link := r.URL.Query().Get("link")
	if !util.IsUrlValid(link) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`invalid link`))
		return
	}
	// Add link to database
	err := database.MainDatabase.AddLink(r.Context(), token, link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`server error: `))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write([]byte(token))
}

// deleteLink deletes a link from database by it's token
func deleteLink(w http.ResponseWriter, r *http.Request) {
	// Check token
	if !authorizeHeader(r) {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// Delete token
	token := r.URL.Query().Get("token")
	err := database.MainDatabase.DeleteLink(r.Context(), token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`server error: `))
		_, _ = w.Write([]byte(err.Error()))
	}
}
