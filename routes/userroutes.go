package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/kalmastenitin/user_management/helpers"
	"github.com/kalmastenitin/user_management/models"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetUserGroupObject(name string, w http.ResponseWriter) models.Group {
	var group models.Group
	err := models.Groupcollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&group)
	if err != nil {
		log.Fatal(err)

	}
	return group
}

type UserInput struct {
	Fullname  string `json:"fullname,omitempty" bson:"fullname,omitempty" validate:"required"`
	Username  string `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	CreatedBy string `json:"created_by,omitempty" bson:"created_by,omitempty"`
	Group     string `json:"group,omitempty" bson:"group,omitempty"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user UserInput
	var userObject models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	err := models.Usercollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&userObject)
	if err == nil {
		helpers.GetErrorCustom("User already exists in system", w, http.StatusConflict)
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	newUserObject := models.User{
		Fullname:     user.Fullname,
		Username:     user.Username,
		Email:        user.Email,
		Phone:        user.Phone,
		Password:     hash,
		DateJoined:   time.Now().UTC(),
		Group:        GetUserGroupObject(user.Group, w).ID,
		Is_active:    true,
		Is_superuser: true,
	}

	result, err := models.Usercollection.InsertOne(context.TODO(), newUserObject)
	if err != nil {
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}
