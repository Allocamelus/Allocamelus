SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PC.postId = ?
  AND PCC.depth <= ?
  AND PC.parent = 0