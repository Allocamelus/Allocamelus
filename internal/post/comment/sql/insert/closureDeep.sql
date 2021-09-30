INSERT INTO PostCommentClosures(parent, child, depth)
SELECT p.parent,
  c.child,
  p.depth + c.depth + 1
FROM PostCommentClosures p,
  PostCommentClosures c
WHERE p.child = ?
  AND c.parent = ?