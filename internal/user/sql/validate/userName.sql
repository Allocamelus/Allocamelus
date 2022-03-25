SELECT EXISTS(
    SELECT *
    FROM Users
    WHERE userName = ?
  )