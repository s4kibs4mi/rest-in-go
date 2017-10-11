package utils

import "github.com/nu7hatch/gouuid"

func GetUUI() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return uuid.String()
}
