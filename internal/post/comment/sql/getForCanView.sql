SELECT postId,
  userId
FROM PostComments
WHERE postCommentId = ?
LIMIT 1