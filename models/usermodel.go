package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var Rolecollection *mongo.Collection
var Permissioncollection *mongo.Collection
var Groupcollection *mongo.Collection
var Grouppermissioncollection *mongo.Collection
var Usercollection *mongo.Collection

func RoleCollection(c *mongo.Database) {
	Rolecollection = c.Collection("role")
}

func PermissionCollection(c *mongo.Database) {
	Permissioncollection = c.Collection("permission")
}

func GroupCollection(c *mongo.Database) {
	Groupcollection = c.Collection("group")
}

func GroupPermissionCollection(c *mongo.Database) {
	Grouppermissioncollection = c.Collection("group_permissions")
}

func UserCollection(c *mongo.Database) {
	Usercollection = c.Collection("user")
}

// type Role struct {
// 	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
// 	Role        string             `json:"role,omitempty" bson:"role,omitempty"`
// 	Permissions []string           `json:"permissions,omitempty" bson:"permissions,omitempty"`
// 	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
// 	CreatedBy   primitive.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
// 	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
// }

type Permission struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	CodeName  string             `json:"codename,omitempty" bson:"codename,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Group struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type GroupPermissions struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GroupID     *Group             `json:"group,omitempty" bson:"group,omitempty"`
	PermissonID *Permission        `json:"permission,omitempty" bson:"permission,omitempty"`
}

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname     string             `json:"fullname,omitempty" bson:"fullname,omitempty" validate:"required"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password     []byte             `json:"password,omitempty" bson:"password,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty" validate:"required,email"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty" validate:"required"`
	CreatedBy    primitive.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
	Group        primitive.ObjectID `json:"group,omitempty" bson:"group,omitempty"`
	Is_active    bool               `json:"is_active" bson:"is_active" validate:"required"`
	DateJoined   time.Time          `json:"date_joined,omitempty" bson:"date_joined,omitempty"`
	Is_superuser bool               `json:"is_superuser" bson:"is_superuser"`
}
