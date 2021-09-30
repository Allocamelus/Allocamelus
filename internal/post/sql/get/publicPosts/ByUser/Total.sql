SELECT COUNT(postId)
FROM Posts
WHERE published != 0
    AND userId = ?