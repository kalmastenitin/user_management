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
	r.HandleFunc("/assignpermissionmany", routes.AssignPermissionsMany).Methods("POST")
	r.HandleFunc("/getpermission", routes.GetPermission).Methods("GET")
	r.HandleFunc("/removegrouppermission", routes.DeleteGroupPermission).Methods("DELETE")

	r.HandleFunc("/user/create", routes.CreateUser).Methods("POST")
	r.HandleFunc("/user/update", routes.UpdateUser).Methods("PUT")
	r.HandleFunc("/user", routes.GetUser).Methods("GET")
	r.HandleFunc("/user", routes.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8001", r))
}
