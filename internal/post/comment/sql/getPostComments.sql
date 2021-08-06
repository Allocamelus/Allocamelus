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
        ORDER BY postCommentId DESC -- Newest first
        LIMIT ?, ?
      ) tmp
  )
  AND PCC.depth <= ?
ORDER BY PC.postCommentId ASC -- Comments are flipped back for sorting
LIMIT ?