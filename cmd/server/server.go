package server

import (
	"encoding/json"
	"fmt"
	. "github.com/eradszewski/dos-template-minon/cmd/domain"
	loggerI "github.com/eradszewski/dos-template-minon/cmd/logger"
	"net/http"
)

var Articles []Article
var logger = loggerI.NewLogger()

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() error {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	return http.ListenAndServe(":10000", nil)

}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	logger.PrettyHttpRequestServer(r)

	err := json.NewEncoder(w).Encode(Articles)
	if err != nil {
		return
	}
}

func Run(config *Config) error {
	Articles = []Article{
		{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	err := handleRequests()
	return err

}
