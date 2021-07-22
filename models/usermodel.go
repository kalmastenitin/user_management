package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Role        string             `json:"role,omitempty" bson:"role,omitempty"`
	Permissions []string           `json:"permissions,omitempty" bson:"permissions,omitempty"`
	CreatedAt   time.Now           `json:"created_at,omitempty" bson:"created_at,omitempty"`
	CreatedBy   primitive.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Permission struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	CodeName  string             `json:"codename,omitempty" bson:"codename,omitempty"`
	CreatedAt time.Now           `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type Group struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	CreatedAt time.Now           `json:"created_at,omitempty" bson:"created_at,omitempty"`
}

type GroupPermissions struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	GroupID     primitive.ObjectID `json:"group_id,omitempty" bson:"group_id,omitempty"`
	PermissonID primitive.ObjectID `json:"permission_id,omitempty" bson:"permission_id,omitempty"`
}

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Fullname     string             `json:"fullname,omitempty" bson:"fullname,omitempty"`
	Username     string             `json:"username,omitempty" bson:"username,omitempty"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	Role         primitive.ObjectID `json:"role,omitempty" bson:"role,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
	CreatedBy    primitive.ObjectID `json:"created_by,omitempty" bson:"created_by,omitempty"`
	Group        primitive.ObjectID `json:"group,omitempty" bson:"group,omitempty"`
	Is_active    bool               `json:"is_active,omitempty" bson:"is_active,omitempty"`
	DateJoined   time.Now           `json:"date_joined,omitempty" bson:"date_joined,omitempty"`
	Is_superuser bool               `json:"is_superuser,omitempty" bson:"is_superuser,omitempty"`
}
