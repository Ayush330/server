DELIMITER $$

CREATE PROCEDURE GetUserBalance(
    IN input_user_id BIGINT UNSIGNED,
    IN input_group_id BIGINT UNSIGNED,
    OUT user_balance DECIMAL(10,2)
)
BEGIN
    DECLARE total_paid DECIMAL(10,2) DEFAULT 0;
    DECLARE total_owed DECIMAL(10,2) DEFAULT 0;

    -- Calculate total paid by the user in the group
    SELECT SUM(e.total_amount)
    INTO total_paid
    FROM expenses e
    WHERE e.paid_by = input_user_id AND e.group_id = input_group_id;

    -- Calculate total owed by the user in the group
    SELECT SUM(es.amount)
    INTO total_owed
    FROM expense_splits es
    JOIN expenses e ON es.expense_id = e.expense_id
    WHERE es.user_id = input_user_id AND e.group_id = input_group_id;

    -- Calculate the balance
    SET user_balance = IFNULL(total_paid, 0) - IFNULL(total_owed, 0);
END$$

DELIMITER ;




DELIMITER $$

CREATE PROCEDURE GetUserBalancesForGroup(
    IN input_user_id BIGINT UNSIGNED,
    IN input_group_id BIGINT UNSIGNED
)
BEGIN
    -- Temporary table to store results
    CREATE TEMPORARY TABLE IF NOT EXISTS UserBalances (
        user_id BIGINT UNSIGNED,
        user_balance DECIMAL(10,2)
    );

    -- Insert balances for all users who have transactions with the input user
    INSERT INTO UserBalances (user_id, user_balance)
    SELECT 
        es.user_id AS user_id,
        SUM(IFNULL(e.total_amount, 0)) - SUM(IFNULL(es.amount, 0)) AS user_balance
    FROM 
        expense_splits es
    JOIN 
        expenses e ON es.expense_id = e.expense_id
    WHERE 
        e.group_id = input_group_id
        AND (es.user_id = input_user_id OR e.paid_by = input_user_id)
    GROUP BY 
        es.user_id;

    -- Output results
    SELECT * FROM UserBalances;

    -- Cleanup
    DROP TEMPORARY TABLE IF EXISTS UserBalances;
END$$

DELIMITER ;
