SELECT EXISTS(
    SELECT *
    FROM UserTokens
    WHERE selector = ?
  )