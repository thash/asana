package api

import (
	"encoding/json"

	"github.com/thash/asana/utils"
)

type Me_t struct {
	Gid        string            `json:"gid"`
	Name       string            `json:"name"`
	Email      string            `json:"email"`
	Workspaces []Base            `json:"workspaces"`
	Photo      map[string]string `json:"photo"`
}

func Me() Me_t {
	var me map[string]Me_t
	err := json.Unmarshal(Get("/api/1.0/users/me", nil), &me)
	utils.Check(err)
	return me["data"]
}
