--find friend 
SELECT distinct(t1.user_id)
FROM user_groups AS t1
INNER JOIN user_groups AS t2
  ON t1.group_id = t2.group_id
WHERE t2.user_id = 2
  AND t1.user_id != 2;

--------------------------------------------------------------------------------------------------------------------------------


-- overall debt/credit

SELECT COALESCE((SELECT SUM(
    CASE WHEN user_id = 1 THEN -1*amount
        ELSE amount
    END
) FROM expense_splits WHERE (creditor = 1 OR user_id = 1) AND (creditor != user_id) AND settled = 0), 0);

-- debt/credit for a group 

SELECT COALESCE((SELECT SUM(
    CASE WHEN user_id = 1 THEN -1*amount
        ELSE amount
    END
) FROM expense_splits WHERE (creditor = 1 OR user_id = 1) AND (creditor != user_id) AND settled = 0 AND group_id = 1), 0);

-- debt credit for each friends;

WITH friends AS (
    SELECT DISTINCT t1.user_id AS friend_id
    FROM user_groups AS t1
    INNER JOIN user_groups AS t2
    ON t1.group_id = t2.group_id
    WHERE t2.user_id = 3
    AND t1.user_id != 3
)
SELECT 
    f.friend_id,
    SUM(CASE
            WHEN es.creditor = 3 THEN es.amount
            ELSE -1 * es.amount
        END) AS total_amount_to_friend
FROM expense_splits es
JOIN friends f ON (es.user_id = f.friend_id OR es.creditor = f.friend_id)
WHERE ((es.user_id IN (SELECT friend_id FROM friends) AND es.creditor = 3)
   OR (es.user_id = 3 AND es.creditor IN (SELECT friend_id FROM friends))
   AND es.settled = 0)
GROUP BY f.friend_id;
