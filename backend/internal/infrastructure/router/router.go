package router

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/hasura/go-graphql-client"
	"github.com/marcustut/fyp/backend/config"
)

// New creates route endpoint
func New(srv *handler.Server) *mux.Router {
	r := mux.NewRouter()

	r.Handle("/graphql", srv).Methods("POST")
	r.Handle("/playground", playground.Handler("GraphQL", "/graphql")).Methods("GET")

	r.HandleFunc("/oauth/github", func(w http.ResponseWriter, r *http.Request) {
		// get the value of the `code` query param
		err := r.ParseForm()
		if err != nil {
			log.Printf("could not parse query: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		code := r.FormValue("code")

		// create a request to call GitHub's OAuth2 endpoint to gain access token
		reqURL := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", config.C.Services.Auth.GithubClientID, config.C.Services.Auth.GithubClientSecret, code)
		req, err := http.NewRequest(http.MethodPost, reqURL, nil)
		if err != nil {
			log.Printf("unable to create HTTP request: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
		req.Header.Set("accept", "application/json")

		// execute the request
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("unable to send HTTP request: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer res.Body.Close()

		// parse response into JSON
		var t struct {
			AccessToken string `json:"access_token"`
		}
		if err = json.NewDecoder(res.Body).Decode(&t); err != nil {
			log.Printf("unable to parse JSON response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}

		// new graphql client
		client := graphql.NewClient("http://localhost:8081/graphql", nil)

		// send a graphql query to login with GitHub access_token
		var m struct {
			SignInWithGithub struct {
				AccessToken graphql.String `graphql:"access_token"`
			} `graphql:"SignInWithGithub(token: $token)"`
		}
		err = client.Mutate(context.Background(), &m, map[string]interface{}{
			"token": graphql.String(t.AccessToken),
		})
		if err != nil {
			log.Printf("unable to make GraphQL request to auth-service: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		log.Println(m)

		// finally, send response to redirect user to welcome page with access token
		w.Header().Set("Location", "https://www.google.com?access_token="+string(m.SignInWithGithub.AccessToken))
		w.WriteHeader(http.StatusFound)
	})

	return r
}
