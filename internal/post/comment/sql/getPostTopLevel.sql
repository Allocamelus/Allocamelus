SELECT COUNT(*)
FROM PostComments
WHERE postId = ?
  AND parent = 0