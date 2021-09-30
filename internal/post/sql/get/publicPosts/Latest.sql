SELECT postId,
    userId,
    published,
    updated,
    content
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
ORDER BY published DESC
LIMIT ?, ?