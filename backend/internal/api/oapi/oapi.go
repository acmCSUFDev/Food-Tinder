package oapi

import (
	"encoding/json"
	"fmt"

	"github.com/acmCSUFDev/Food-Tinder/backend/foodtinder"
	"github.com/bwmarrin/snowflake"
)

//go:generate goapi-gen --config=config.json ../../../../openapi/foodtinder.jsonc

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(foodtinder.ID(id).String()))
}

func (id *ID) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	s, err := snowflake.ParseString(str)
	if err != nil {
		return fmt.Errorf("invalid snowflake: %v", err)
	}
	*id = ID(s)
	return nil
}

// RespErr returns an Error response value from an error.
func RespErr(err error) Error {
	return Error{err.Error()}
}
