SELECT PC.postId,
  PC.userId,
  PC.parent,
  PC.created,
  PC.updated,
  PC.content,
  PCC.depth
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.child = ?
ORDER BY PCC.depth DESC
LIMIT 1