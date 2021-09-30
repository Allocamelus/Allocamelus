SELECT userId,
  published
FROM Posts
WHERE postId = ?
LIMIT 1