package url

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/urlshortner/cors"
)

const SHORTEN_PATH = "shorten"
const HTTPS = "https://"

func SetupRoutes(API_BASE_PATH string) {
	handleShortner := http.HandlerFunc(shortHandler)
	handleRedirection := http.HandlerFunc(redirectHandler)
	http.Handle(fmt.Sprintf("%s/%s", API_BASE_PATH, SHORTEN_PATH), cors.Middleware(handleShortner))
	http.Handle("/", cors.Middleware(handleRedirection))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			shortenUrl := strings.TrimPrefix(r.URL.Path, "/")
			originalUrl, err := fetchOriginalUrl(shortenUrl)
			newUrl := HTTPS + originalUrl
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, newUrl, http.StatusSeeOther)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
	}
} 

func shortHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			var url URL
			bodyBytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			err = json.Unmarshal(bodyBytes, &url)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return 
			}
			// way to fetch details from form
			// r.ParseForm()
			// url.OriginalUrl = r.FormValue("url")
			// url.ShortenUrl = r.FormValue("shortform")
			if(url.ShortenUrl == "") {
				url.ShortenUrl = getShortUrl(url.OriginalUrl)
			}
			err = insertOrUpdateUrl(url)
			if err != nil {
				// fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			url.ShortenUrl = r.Host + "/" + url.ShortenUrl
			urlJson, err := json.Marshal(url)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(urlJson)
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
	}
}

func getShortUrl(longUrl string) string {
	h := sha256.New()
	h.Write([]byte(longUrl))
    hash := hex.EncodeToString(h.Sum(nil))
	encodedUrl := base64.StdEncoding.EncodeToString([]byte(hash[:8]))
	return encodedUrl
}
