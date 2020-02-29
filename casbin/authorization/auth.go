package authorization

import (
	"errors"
	"log"
	"net/http"

	"github.com/alexedwards/scs/session"
	"github.com/casbin/casbin"

	"github.com/a8uhnf/hello-world/casbin/model"
)

func Authorizer(e *casbin.Enforcer, users model.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role, err := session.GetString(r, "role")
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}

			if role == "" {
				role = "anonymous"
			}

			// if it's a member, check if the user still exists
			if role == "member" {
				uid, err := session.GetInt(r, "userID")
				if err != nil {
					writeError(http.StatusInternalServerError, "ERROR", w, err)
					return
				}
				exists := users.Exists(uid)
				if !exists {
					writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}

			// casbin rule enforcing
			res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}
			if res {
				next.ServeHTTP(w, r)
			} else {
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
		}

		return http.HandlerFunc(fn)
	}
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
