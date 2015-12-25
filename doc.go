// Package jwt provides methods to create and check JSON Web Tokens. It only implements HMAC 256 encryption and has a very small footprint, ideal for simple usage when authorizing clients
/*

	package main

	import (
		auth "github.com/nilslice/jwt"
		"fmt"
		"net/http"
		"strings"
	)

	func main() {
		http.HandleFunc("/auth/new", func(res http.ResponseWriter, req *http.Request) {
			token, err := auth.New()
			if err != nil {
				http.Error(res, "Error", 500)
				return
			}
			res.Header().Add("Authorization", "Bearer "+token)

			res.WriteHeader(http.StatusOK)
		})

		http.HandleFunc("/auth", func(res http.ResponseWriter, req *http.Request) {
			userToken := strings.Split(req.Header.Get("Authorization"), " ")[1]

			if auth.Passes(userToken) {
				fmt.Println("ok")
			} else {
				fmt.Println("no")
			}
		})

		http.ListenAndServe(":8080", nil)
	}
*/
package jwt
