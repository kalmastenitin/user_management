package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"github.com/kalmastenitin/user_management/config"
	"github.com/kalmastenitin/user_management/routes"
)

func main() {
	config.Connect()
	r := mux.NewRouter()
	routes.Validate = validator.New()

	r.HandleFunc("/permission", routes.GetAllPermissions).Methods("GET")
	r.HandleFunc("/permission", routes.CreatePermission).Methods("POST")
	r.HandleFunc("/permission", routes.DeletePermission).Methods("DELETE")
	r.HandleFunc("/group", routes.GetAllGroups).Methods("GET")
	r.HandleFunc("/group", routes.CreateGroup).Methods("POST")
	r.HandleFunc("/group", routes.DeleteGroup).Methods("DELETE")
	r.HandleFunc("/assignpermission", routes.AssignPermissions).Methods("POST")
	r.HandleFunc("/assignpermissionmany", routes.AssignPermissionsMany).Methods("POST")
	r.HandleFunc("/getpermission", routes.GetPermission).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", r))
}
