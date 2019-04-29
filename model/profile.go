package model

import "encoding/json"

type Profile struct {
	Name string
	Gender string
	Age string
	JobTitle string
	Category string
	College string
	Email string
	Phone string
	Office string
}

func FromJsonObj(o interface{}) (Profile, error) {
	var profile Profile
	s, err := json.Marshal(o)
	if err != nil {
		return profile, err
	}

	err = json.Unmarshal(s, &profile)
	return profile, err
}

