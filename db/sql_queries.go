package db

import (
	"database/sql"
	"errors"

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

func GetNetExpenseDetailsForAUserForAGroup(payload models.ExpenseDetailsUserGroupPayload) (float64, error) {
	var balance float64

	// Start a transaction (to ensure consistency if needed)
	tx, err := db.Begin()
	if err != nil {
		logrus.Error("Error starting transaction: ", err)
		return 0.0, errors.New("failed to start transaction")
	}
	defer tx.Rollback() // Ensure rollback in case of failure

	// Call the stored procedure to set @balance
	_, err = tx.Exec("CALL GetUserBalance(?, ?, @balance)", payload.UserID, payload.GroupId)
	if err != nil {
		logrus.Error("Error calling GetUserBalance SP: ", err)
		return 0.0, errors.New("calling GetUserBalance SP failed")
	}

	// Retrieve the value of @balance
	err = tx.QueryRow("SELECT @balance").Scan(&balance)
	if err != nil {
		logrus.Error("Error scanning balance: ", err)
		return 0.0, errors.New("GetUserBalance SP passed, but could not retrieve the value")
	}

	// Commit the transaction (if no errors)
	if err := tx.Commit(); err != nil {
		logrus.Error("Error committing transaction: ", err)
		return 0.0, errors.New("failed to commit transaction")
	}

	return balance, nil
}

func logErrorAndRollBack(err error, TransactionInstance *sql.Tx) bool {
	if err != nil {
		logrus.Error("Error in Transaction ", err.Error())
		TransactionInstance.Rollback()
		return true
	}
	return false
}
