SELECT postId,
    published,
    updated,
    content
FROM Posts
WHERE published != 0
    AND userId = ?
ORDER BY published DESC
LIMIT ?, ?