SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PCC.parent = ?
  AND PCC.parent != PCC.child