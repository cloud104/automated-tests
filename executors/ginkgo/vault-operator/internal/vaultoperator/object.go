package vaultoperator

import (
	"encoding/json"
)

type Users []User

func UnmarshalUsers(data []byte) (Users, error) {
	var r Users
	err := json.Unmarshal(data, &r)

	return r, err
}

func (r *Users) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type User struct {
	Password      string `json:"password"`
	TokenPolicies string `json:"token_policies"`
	Username      string `json:"username"`
}
