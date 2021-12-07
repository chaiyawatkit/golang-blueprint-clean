package utils

import (
	"fmt"
	"golang-blueprint-clean/app/constants"
)

func GetHumanErrorCode(keyMassage string) string {
	if constants.HumanErrorCode[keyMassage] == nil {
		keyMassage = "default"
	}
	humanMessage := fmt.Sprintf("%v", constants.HumanErrorCode[keyMassage])
	return humanMessage
}

func GetHumanSuccessCode(keyMassage string) string {
	if constants.HumanSuccessCode[keyMassage] == nil {
		keyMassage = "default"
	}
	humanMessage := fmt.Sprintf("%v", constants.HumanSuccessCode[keyMassage])
	return humanMessage
}
