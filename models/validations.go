package models

// import (
// 	"log"
// 	"net/http"

// 	"github.com/go-playground/validator"
// 	"github.com/kalmastenitin/user_management/helpers"
// )

// var validate *validator.Validate

// func ValidatePermissions(*permission struct, w http.ResponseWriter) {
// 	w.Header().Set("Content-Type", "application/json")
// 	err := validate.Struct(permission)
// 	if err != nil {
// 		log.Println(err)
// 		helpers.GetError(err, w, http.StatusBadRequest)
// 	}
// }
