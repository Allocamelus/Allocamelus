SELECT PC.postCommentId,
  PC.postId,
  PC.userId,
  PC.parent,
  PC.created,
  PC.updated,
  PC.content,
  PCC.depth
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.parent = ?
  AND PCC.depth <= ?
  AND PCC.parent != PCC.child
  LIMIT ?, ?