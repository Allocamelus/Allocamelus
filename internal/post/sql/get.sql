SELECT userId,
  created,
  published,
  updated,
  content
FROM Posts
WHERE postId = ?
LIMIT 1