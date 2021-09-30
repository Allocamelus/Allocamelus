SELECT postId,
    userId,
    published,
    updated,
    content
FROM Posts
WHERE userId IN (
        SELECT followUserId
        FROM (
                SELECT followUserId
                FROM UserFollows
                WHERE userId = ?
                    AND accepted = 1
            ) tmp
    )
    AND published != 0
    OR userId = ?
ORDER BY published DESC
LIMIT ?, ?