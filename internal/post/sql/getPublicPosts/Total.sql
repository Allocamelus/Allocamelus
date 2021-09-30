SELECT COUNT(postId)
FROM Posts
WHERE userId IN (
    SELECT userId
    FROM (
        SELECT userId
        FROM Users
        WHERE type = ?
      ) tmp
  )
  AND published != 0