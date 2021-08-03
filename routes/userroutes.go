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
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetGroupName(id primitive.ObjectID) string {
	var group models.Group
	err := models.Groupcollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&group)
	if err != nil {
		log.Fatal(err)

	}
	return group.Name
}

type UserInput struct {
	Fullname     string `json:"fullname,omitempty" bson:"fullname,omitempty" validate:"required"`
	Username     string `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password     string `json:"password,omitempty" bson:"password,omitempty"`
	Email        string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Phone        string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	CreatedBy    string `json:"created_by,omitempty" bson:"created_by,omitempty"`
	Group        string `json:"group,omitempty" bson:"group,omitempty"`
	Is_superuser bool   `json:"is_superuser,omitempty" bson:"is_superuser,omitempty"`
	Is_active    bool   `json:"is_active,omitempty" bson:"is_active,omitempty"`
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
		Is_active:    user.Is_active,
		Is_superuser: user.Is_superuser,
	}
	log.Println(newUserObject.Is_superuser)
	result, err := models.Usercollection.InsertOne(context.TODO(), newUserObject)
	if err != nil {
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

type UserOutPut struct {
	Fullname     string `json:"fullname,omitempty" bson:"fullname,omitempty" validate:"required"`
	Username     string `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Email        string `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Phone        string `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	CreatedBy    string `json:"created_by,omitempty" bson:"created_by,omitempty"`
	Group        string `json:"group,omitempty" bson:"group,omitempty"`
	Is_superuser bool   `json:"is_superuser,omitempty" bson:"is_superuser,omitempty"`
	Is_active    bool   `json:"is_active,omitempty" bson:"is_active,omitempty"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	var users []UserOutPut
	username := r.URL.Query().Get("username")

	// username is provided
	if username != "" {
		log.Println(username)
		err := models.Usercollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)

		if err != nil {
			helpers.GetErrorCustom("User Not Exist", w, http.StatusNotFound)
			return
		}

		newUserObject := UserOutPut{
			Fullname: user.Fullname,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Group:    GetGroupName(user.Group),
			// DateJoined:   user.DateJoined.String(),
			Is_active:    user.Is_active,
			Is_superuser: user.Is_superuser,
		}

		users = append(users, newUserObject)
		json.NewEncoder(w).Encode(users)
		return
	}

	cursor, err := models.Usercollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Printf("error while getting all users, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var user models.User
		err := cursor.Decode(&user)
		newUserObject := UserOutPut{
			Fullname: user.Fullname,
			Username: user.Username,
			Email:    user.Email,
			Phone:    user.Phone,
			Group:    GetGroupName(user.Group),
			// DateJoined:   user.DateJoined.String(),
			Is_active:    user.Is_active,
			Is_superuser: user.Is_superuser,
		}

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, newUserObject)
	}
	json.NewEncoder(w).Encode(users)

}
