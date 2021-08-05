DELETE PC
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.parent = ?