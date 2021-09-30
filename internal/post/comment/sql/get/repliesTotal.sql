SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PCC.parent = ?
  /* PCC.parent != PCC.child prevents query from counting parent's self */
  AND PCC.parent != PCC.child