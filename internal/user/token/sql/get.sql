SELECT userTokenId,
  userId,
  tokenType,
  token,
  expiration
FROM UserTokens
WHERE selector = ?
LIMIT 1