package api

import (
	"encoding/json"

	"github.com/memerelics/asana/utils"
)

type Me_t struct {
	Id         int
	Name       string
	Email      string
	Workspaces []Base
	Photo      map[string]string
}

func Me() Me_t {
	var me map[string]Me_t
	err := json.Unmarshal(Get("/api/1.0/users/me", nil), &me)
	utils.Check(err)
	return me["data"]
}
