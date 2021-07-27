package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/kalmastenitin/user_management/helpers"
	"github.com/kalmastenitin/user_management/models"
	"go.mongodb.org/mongo-driver/bson"
)

var Validate *validator.Validate

func GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	// get all permission or get selected permission
	w.Header().Set("Content-Type", "application/json")
	permissions := []models.Permission{}

	// if user provided query
	codename := r.URL.Query().Get("codename")
	if codename != "" {
		var permission models.Permission
		err := models.Permissioncollection.FindOne(context.TODO(), bson.M{"codename": codename}).Decode(&permission)
		if err != nil {
			helpers.GetError(err, w, http.StatusNotFound)
			return
		}
		permissions = append(permissions, permission)
		json.NewEncoder(w).Encode(permissions)
		return
	}

	// if user not provided query
	cursor, err := models.Permissioncollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("Error while getting all permissions, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var permission models.Permission
		err := cursor.Decode(&permission)
		if err != nil {
			log.Fatal(err)
		}

		permissions = append(permissions, permission)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(permissions)
}

func CreatePermission(w http.ResponseWriter, r *http.Request) {
	// add new permission
	w.Header().Set("Content-Type", "application/json")
	var permission models.Permission

	_ = json.NewDecoder(r.Body).Decode(&permission)

	newPermission := models.Permission{
		Name:      permission.Name,
		CodeName:  permission.CodeName,
		CreatedAt: time.Now().UTC(),
	}

	errs := Validate.Struct(newPermission)
	if errs != nil {
		helpers.GetError(errs, w, http.StatusBadRequest)
		return
	}

	err := models.Permissioncollection.FindOne(context.TODO(), bson.M{"codename": permission.CodeName}).Decode(&permission)
	if err == nil {
		helpers.GetErrorCustom("permission already exists in system", w, http.StatusConflict)
		return
	}

	result, err := models.Permissioncollection.InsertOne(context.TODO(), newPermission)

	if err != nil {
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeletePermission(w http.ResponseWriter, r *http.Request) {
	// delete permission from database
	w.Header().Set("Content-Type", "application/json")
	codename := r.URL.Query().Get("codename")

	deleteresult, err := models.Permissioncollection.DeleteOne(context.TODO(), bson.M{"codename": codename})
	log.Println(deleteresult)
	if err != nil {
		helpers.GetError(err, w, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(deleteresult)
}
