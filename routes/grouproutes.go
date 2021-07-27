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
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	groups := []models.Group{}

	groupname := r.URL.Query().Get("name")
	if groupname != "" {
		var group models.Group
		err := models.Groupcollection.FindOne(context.TODO(), bson.M{"name": groupname}).Decode(&group)
		if err != nil {
			helpers.GetError(err, w, http.StatusNotFound)
			return
		}
		groups = append(groups, group)
		json.NewEncoder(w).Encode(groups)
		return
	}

	cursor, err := models.Groupcollection.Find(context.TODO(), bson.M{})

	if err != nil {
		log.Printf("error while getting all groups, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var group models.Group
		err := cursor.Decode(&group)
		if err != nil {
			log.Fatal(err)
		}

		groups = append(groups, group)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(groups)
}

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var group models.Group

	_ = json.NewDecoder(r.Body).Decode(&group)

	newGroup := models.Group{
		Name:      group.Name,
		CreatedAt: time.Now().UTC(),
	}

	errs := Validate.Struct(newGroup)
	if errs != nil {
		helpers.GetError(errs, w, http.StatusBadRequest)
		return
	}

	err := models.Groupcollection.FindOne(context.TODO(), bson.M{"name": group.Name}).Decode(&group)
	if err == nil {
		helpers.GetErrorCustom("group is already exists in system", w, http.StatusConflict)
		return
	}

	result, err := models.Groupcollection.InsertOne(context.TODO(), newGroup)
	if err != nil {
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	groupname := r.URL.Query().Get("name")

	deleteresult, err := models.Groupcollection.DeleteOne(context.TODO(), bson.M{"name": groupname})
	if err != nil {
		helpers.GetError(err, w, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(deleteresult)
}

func AssignPermissions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var group models.Group
	var permission models.Permission
	var m map[string]string

	_ = json.NewDecoder(r.Body).Decode(&m)
	log.Println(m["PermissionID"])
	err := models.Groupcollection.FindOne(context.TODO(), bson.M{"name": m["GroupID"]}).Decode(&group)
	if err != nil {
		log.Printf("error while getting group, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	err = models.Permissioncollection.FindOne(context.TODO(), bson.M{"codename": m["PermissionID"]}).Decode(&permission)
	if err != nil {
		log.Printf("error while getting permission, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}
	// json.Unmarshal(r.Body, m)

	newGroupPermissions := models.GroupPermissions{
		GroupID:     &group,
		PermissonID: &permission,
	}

	result, err := models.Grouppermissioncollection.InsertOne(context.TODO(), newGroupPermissions)
	if err != nil {
		helpers.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)

}

type ManyPermission struct {
	GroupID      string
	PermissionID []string
}

func AssignPermissionsMany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group models.Group
	var permission models.Permission
	var permissionObjects []interface{}
	var permissions ManyPermission

	json.NewDecoder(r.Body).Decode(&permissions)

	err := models.Groupcollection.FindOne(context.TODO(), bson.M{"name": permissions.GroupID}).Decode(&group)
	if err != nil {
		log.Printf("error while getting group, Reason: %v\n", err)
		helpers.GetError(err, w, http.StatusNotFound)
		return
	}

	for _, value := range permissions.PermissionID {
		log.Println(value)
		err := models.Permissioncollection.FindOne(context.TODO(), bson.M{"codename": value}).Decode(&permission)
		if err != nil {
			log.Printf("error while getting permission, Reason: %v\n", err)
			helpers.GetError(err, w, http.StatusNotFound)
			return
		}

		log.Println(group.Name, value)
		var group_permission models.GroupPermissions
		errs := models.Grouppermissioncollection.FindOne(context.TODO(), bson.M{"group.name": group.Name, "permission.codename": value}).Decode(&group_permission)

		if errs != nil {
			newGroupPermissions := models.GroupPermissions{
				GroupID:     &group,
				PermissonID: &permission,
			}
			log.Println("error nil")
			permissionObjects = append(permissionObjects, newGroupPermissions)
		} else {
			log.Println("error")
		}

	}
	log.Println(permissionObjects...)
	if len(permissionObjects) != 0 {
		result, err := models.Grouppermissioncollection.InsertMany(context.TODO(), permissionObjects, options.InsertMany().SetOrdered(true))
		if err != nil {
			helpers.GetError(err, w, http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(result)
		return
	}
	helpers.ReturnSuccess("Success", w, http.StatusAccepted)
	return
}

func GetPermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var group_permission models.GroupPermissions

	err := models.Grouppermissioncollection.FindOne(context.TODO(), bson.M{"group.name": "Superadmin", "permission.codename": "view_permission"}).Decode(&group_permission)
	if err != nil {
		helpers.GetErrorCustom("Permission Not Found", w, http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(group_permission)
}
