package exception

import (
	"errors"
	"log"
	"net/http"

	"github.com/IwanPlamboyan/contact-manajemen-golang/helper"
)

func HandleError(writer http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		if appErr.Internal != nil {
			log.Printf("[ERROR] %+v\n", appErr.Internal)
		}

		helper.ResponseJSONError(writer, appErr.Code, appErr.Message)
		return
	}

	// fallback (unexpected error)
	log.Printf("[ERROR] %+v\n", err)
	helper.ResponseJSONError(writer, http.StatusInternalServerError, "Internal Server Error")
}