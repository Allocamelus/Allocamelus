INSERT INTO UserTokens (
    userId,
    tokenType,
    selector,
    token,
    created,
    expiration
  )
VALUES (?, ?, ?, ?, ?, ?)