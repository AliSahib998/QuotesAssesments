package errhandler

import (
	"github.com/AliSahib998/QuotesAssesments/util"
	"net/http"
)

func ErrorHandler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			switch e := err.(type) {
			case *BadRequestError:
				w.WriteHeader(http.StatusBadRequest)
				util.Encode(w, e)
			case *NotFoundError:
				w.WriteHeader(http.StatusNotFound)
				util.Encode(w, e)
			case *AuthenticationError:
				w.WriteHeader(http.StatusForbidden)
				util.Encode(w, e)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}
	}
}

type handlerFunc func(w http.ResponseWriter, r *http.Request) error
