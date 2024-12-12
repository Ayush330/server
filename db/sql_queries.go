package db

import (
	"github.com/Ayush330/server/models"
	"github.com/sirupsen/logrus"
)

func CreateNewUser(payload models.CreateUserPayload) bool {
	_, err := db.Exec("INSERT INTO users (username, email, hashed_password) VALUES (?, ?, ?)", payload.UserName, payload.Email, payload.HashedPassword)
	return err == nil
}

func CreateNewGroup(payload models.CreateGroupPayload) bool {
	logrus.Warn("The payload is: ", payload)
	_, err := db.Exec("INSERT INTO group_mapping (group_name, created_by) VALUES (?, ?)", payload.GroupName, payload.CreatorID)
	if err != nil {
		logrus.Error("Error in CreateNewGroup: ", err.Error())
	}
	return err == nil
}
