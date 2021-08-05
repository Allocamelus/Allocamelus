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
WHERE PCC.parent IN (
    SELECT postCommentId
    FROM (
        SELECT postCommentId
        FROM PostComments
        WHERE postId = ?
          AND parent = 0
        LIMIT ?, ?
      ) tmp
  )
  AND PCC.depth <= ?
LIMIT ?