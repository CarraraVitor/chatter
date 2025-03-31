package middleware

import (
	"log"
	"net/http"

	"chatter/app/handlers"
	"chatter/app/services"
)

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/login" {
			authenticated := func() bool {
				cookie, err := r.Cookie("SessionToken")
				if err != nil {
					return false
				}

				token := cookie.Value
				if token == "" {
					return false
				}
				_, err = services.ValidateSessionToken(token)
				if err != nil {
					return false
				}
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return true
			}()
			if authenticated {
				return
			}
		}
		if !handlers.PublicRoutes[r.URL.Path] {
			cookie, err := r.Cookie("SessionToken")
			if err != nil {
				log.Printf("AuthMiddleware [%s] -> No Cookie Found: %s\n", r.URL.Path, err.Error())
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			token := cookie.Value
			if token == "" {
				log.Printf("AuthMiddleware [%s] -> No SessionToken Found: %s\n", r.URL.Path, token)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			validation, err := services.ValidateSessionToken(token)
			if err != nil {
				log.Printf("AuthMiddleware [%s] -> Session Validation Failed: %s\n", r.URL.Path, err.Error())
				services.DeleteSessionTokenCookie(w)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			session := validation.Session
			services.SetSessionTokenCookie(w, token, session)
            r.Header.Set("UserId", validation.User.Id.String())
		}
		next.ServeHTTP(w, r)
	})
}
