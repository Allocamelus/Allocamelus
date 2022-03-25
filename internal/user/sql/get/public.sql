SELECT userName,
  name,
  bio,
  type,
  created
FROM Users
WHERE userId = ?
LIMIT 1