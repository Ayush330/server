package db

import (
	"database/sql"

	"github.com/Ayush330/server/models"
	"github.com/sirupsen/logrus"
)

func CreateNewUser(payload models.CreateUserPayload) bool {
	_, err := db.Exec("INSERT INTO users (username, email, hashed_password) VALUES (?, ?, ?)", payload.UserName, payload.Email, payload.HashedPassword)
	return err == nil
}

func CreateNewGroup(payload models.CreateGroupPayload) bool {
	Tx, err := db.Begin()
	if err != nil {
		logrus.Error("Failed to begin transaction for CreateNewGroup")
		return false
	}
	var Result sql.Result
	Result, err = Tx.Exec("INSERT INTO group_mapping (group_name, created_by) VALUES (?, ?)", payload.GroupName, payload.CreatorID)
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	var InsertID int64
	InsertID, err = Result.LastInsertId()
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	_, err = Tx.Exec("INSERT INTO user_groups (user_id, group_id) VALUES (?, ?)", payload.CreatorID, InsertID)
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	Tx.Commit()
	return true
}

func AddToGroup(payload models.AddToGroupPayload) bool {
	_, err := db.Exec("INSERT INTO user_groups (group_id, user_id) VALUES (?, ?)", payload.GroupId, payload.UserId)
	if err != nil {
		logrus.Error("AddToGroup ", err.Error())
	}
	return err == nil
}

func AddExpense(payload models.AddExpensePayload) bool {
	Tx, err := db.Begin()
	if err != nil {
		logrus.Error("Failed to begin transaction for AddExpense ", err.Error())
		return false
	}
	var Result sql.Result
	Result, err = Tx.Exec("INSERT INTO expenses (group_id, paid_by, description, total_amount) VALUES (?, ?, ?, ?)", payload.GroupId, payload.PaidBy, payload.Description, payload.TotalAmount)
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	var ExpenseID int64
	ExpenseID, err = Result.LastInsertId()
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	var PreparedStatement *sql.Stmt
	PreparedStatement, err = Tx.Prepare("INSERT INTO expense_splits (expense_id, user_id, amount) VALUES (?, ?, ?)")
	if logErrorAndRollBack(err, Tx) {
		return false
	}
	for _, splitList := range payload.UserListToSplit {
		_, err = PreparedStatement.Exec(ExpenseID, splitList.UserId, splitList.Amount)
		if logErrorAndRollBack(err, Tx) {
			return false
		}
	}
	Tx.Commit()
	return true
}

func logErrorAndRollBack(err error, TransactionInstance *sql.Tx) bool {
	if err != nil {
		logrus.Error("Error in Transaction ", err.Error())
		TransactionInstance.Rollback()
		return true
	}
	return false
}
