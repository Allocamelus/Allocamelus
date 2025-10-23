-- name: InsertUserToken :exec
INSERT INTO UserTokens (
    userId,
    tokenType,
    selector,
    token,
    created,
    expiration
  )
VALUES ($1, $2, $3, $4, $5, $6);

-- name: UserTokenSelectorExist :one
SELECT EXISTS(
    SELECT *
    FROM UserTokens
    WHERE selector = $1
  );

-- name: GetUserToken :one
SELECT userTokenId,
  userId,
  tokenType,
  token,
  expiration
FROM UserTokens
WHERE selector = $1
LIMIT 1;

-- name: DeleteUserToken :exec
DELETE FROM UserTokens
WHERE userTokenId = $1;

-- name: DeleteUserTokenByUIDAndType :exec
DELETE FROM UserTokens
WHERE userId = $1
  AND tokenType = $2;

-- name: GetUserIDByUserName :one
SELECT userId
FROM Users
WHERE userName = $1
LIMIT 1;

-- name: InsertUser :one
INSERT INTO Users (
    userName,
    name,
    email,
    bio,
    type,
    permissions,
    created
  )
VALUES ($1, '', $2, '', $3, $4, $5) RETURNING userId;

-- name: GetUserIDByEmail :one
SELECT userId
FROM Users
WHERE email = $1
LIMIT 1;

-- name: GetUserEmailByID :one
SELECT email
FROM Users
WHERE userId = $1
LIMIT 1;

-- name: GetUserNameByID :one
SELECT userName FROM Users WHERE userId = $1 LIMIT 1;

-- name: ValidateUserName :one
SELECT EXISTS(
    SELECT *
    FROM Users
    WHERE userName = $1
  );

-- name: UpdateUserType :exec
UPDATE Users
SET type = $1
WHERE userId = $2;

-- name: UpdateUserBio :exec
UPDATE Users
SET bio = $1
WHERE userId = $2;

-- name: ListFollowRequests :many
SELECT userId
FROM UserFollows
WHERE followUserId = $1
  AND accepted = false;

-- name: UpdateUserName :exec
UPDATE Users
SET name = $1
WHERE userId = $2;

-- name: ListFollowers :many
SELECT userId
FROM UserFollows
WHERE followUserId = $1
  AND accepted = true;

-- name: IsFollowing :one
SELECT accepted
FROM UserFollows
WHERE userId = $1
  AND followUserId = $2;

-- name: UnfollowUser :exec
DELETE FROM UserFollows
WHERE userId = $1
  AND followUserId = $2;

-- name: CountFollowers :one
SELECT COUNT(userFollowId)
FROM UserFollows
WHERE followUserId = $1
  AND accepted = true;

-- name: FollowUser :exec
INSERT INTO UserFollows (userId, followUserId, accepted, created)
VALUES ($1, $2, $3, $4);

-- name: ListFollowing :many
SELECT followUserId
FROM UserFollows
WHERE userId = $1
  AND accepted = true;

-- name: AcceptAllFollowRequests :exec
UPDATE UserFollows
SET accepted = true
WHERE followUserId = $1
  AND accepted = false;

-- name: AcceptFollowRequest :exec
UPDATE UserFollows
SET accepted = true
WHERE userId = $1
  AND followUserId = $2;

-- name: GetUserType :one
SELECT type
FROM Users
WHERE userId = $1
LIMIT 1;

-- name: GetPublicUser :one
SELECT userName,
  name,
  bio,
  type,
  created
FROM Users
WHERE userId = $1
LIMIT 1;

-- name: ValidateUserEmail :one
SELECT EXISTS(
    SELECT *
    FROM Users
    WHERE email = $1
  );

-- name: GetUserAuthKeySalt :one
SELECT authKeySalt
FROM UserKeys
WHERE userId = $1
ORDER BY userKeyId DESC
LIMIT 1;

-- name: UpdateUserPermissions :exec
UPDATE Users
SET permissions = $1
WHERE userId = $2;

-- name: GetUserPermissions :one
SELECT permissions
FROM Users
WHERE userId = $1
LIMIT 1;

-- name: GetUserAuthKeyHash :one
SELECT authKeyHash
FROM UserKeys
WHERE userId = $1
ORDER BY userKeyId DESC
LIMIT 1;

-- name: InsertUserKey :exec
INSERT INTO UserKeys (
    userId,
    created,
    authKeyHash,
    authKeySalt,
    publicArmored,
    privateArmored,
    recoveryKeyHash,
    recoveryArmored
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
  );

-- name: InsertUserEvent :one
INSERT INTO UserEvents (userId, eventType, created, info)
VALUES ($1, $2, $3, $4) RETURNING userEventId;

-- name: InsertUserEventKey :exec
INSERT INTO UserEventKeys (userEventId, userKeyId, infoKey)
VALUES ($1, $2, $3);

-- name: CountUserEvents :one
SELECT COUNT(userEventId)
FROM UserEvents
WHERE eventType = $1
  AND userId = $2
  AND created > $3;

-- name: GetUserPrivateArmoredKey :one
SELECT privateArmored
FROM UserKeys
WHERE userId = $1
ORDER BY userKeyId DESC
LIMIT 1;

-- name: GetUserPublicKeys :many
SELECT userKeyId,
  publicArmored
FROM UserKeys
WHERE userId = $1
  AND (
    replaced > $2
    OR replaced = false
  )
ORDER BY userKeyId DESC;

-- name: InsertPost :one
INSERT INTO Posts (userId, created, published, content)
VALUES ($1, $2, $3, $4) RETURNING postId;

-- name: PublishPost :exec
UPDATE Posts
SET published = $1
WHERE postId = $2;

-- name: UpdatePostContent :exec
UPDATE Posts
SET updated = $1,
  content = $2
WHERE postId = $3;

-- name: InsertPostMediaFile :one
INSERT INTO PostMediaFiles (
    created,
    fileType,
    width,
    height,
    hash,
    newHash
  )
VALUES ($1, $2, $3, $4, $5, $6) RETURNING postMediaFileId;

-- name: GetPostUserID :one
SELECT userId
FROM Posts
WHERE postId = $1
LIMIT 1;

-- name: GetPostMedia :many
SELECT PM.alt,
  PMF.width,
  PMF.height,
  PMF.hash
FROM PostMedia PM
  JOIN PostMediaFiles PMF ON (PM.postMediaFileId = PMF.postMediaFileId)
WHERE postId = $1
  AND active = true
ORDER BY postMediaId ASC
LIMIT 4;

-- name: PostMediaHashCheck :one
SELECT hash
FROM PostMediaFiles
WHERE hash = $1
  OR newHash = $1
LIMIT 1;

-- name: DeletePostComment :exec
DELETE FROM PostComments AS PC
USING PostCommentClosures AS PCC
WHERE PC.postCommentId = PCC.child
  AND PCC.parent = $1;

-- name: UpdatePostCommentContent :exec
UPDATE PostComments
SET updated = $1,
  content = $2
WHERE postCommentId = $3;

-- name: InsertPostMedia :exec
INSERT INTO PostMedia (
    postId,
    added,
    active,
    alt,
    postMediaFileId
  )
VALUES ($1, $2, true, $3, $4);

-- name: GetPostPublishedStatus :one
SELECT published
FROM Posts
WHERE postId = $1
LIMIT 1;

-- name: GetPost :one
SELECT userId,
  created,
  published,
  updated,
  content
FROM Posts
WHERE postId = $1
LIMIT 1;

-- name: GetPostMediaFileIDByHash :one
SELECT postMediaFileId
FROM PostMediaFiles
WHERE hash = $1
LIMIT 1;

-- name: GetPostCanView :one
SELECT userId,
  published
FROM Posts
WHERE postId = $1
LIMIT 1;

-- name: InsertPostCommentClosureSelf :exec
INSERT INTO PostCommentClosures(parent, child, depth)
VALUES ($1, $2, 0);

-- name: InsertPostComment :one
INSERT INTO PostComments (
    postId,
    userId,
    parent,
    created,
    content
  )
VALUES ($1, $2, $3, $4, $5) RETURNING postCommentId;

-- name: GetPostCommentUserID :one
SELECT userId
FROM PostComments
WHERE postCommentId = $1
LIMIT 1;

-- name: CountPostCommentReplies :one
SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PCC.parent = $1
  /* PCC.parent != PCC.child prevents query from counting parent's self */
  AND PCC.parent != PCC.child;

-- name: GetPostCommentReplies :many
SELECT sqlc.embed(PC),
  PCC.depth
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.parent = $1
  AND PCC.depth <= $2
  AND PCC.parent != PCC.child
LIMIT $4 OFFSET $3;

-- name: GetPostCommentPostUserID :one
SELECT postId,
  userId
FROM PostComments
WHERE postCommentId = $1
LIMIT 1;

-- name: CountPostCommentsTotal :one
SELECT COUNT(*)
FROM PostComments
WHERE postId = $1;

-- name: CountPostCommentsTotalDepth :one
SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PC.postId = $1
  AND PCC.depth <= $2
  AND PC.parent = 0;

-- name: CountPostCommentsTopLevel :one
SELECT COUNT(*)
FROM PostComments
WHERE postId = $1
  AND parent = 0;

-- name: GetPostCommentPostID :one
SELECT postId
FROM PostComments
WHERE postCommentId = $1
LIMIT 1;

-- name: GetPostComment :one
SELECT sqlc.embed(PC),
  PCC.depth
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.child = $1
ORDER BY PCC.depth DESC
LIMIT 1;

-- name: GetPostComments :many
SELECT sqlc.embed(PC),
  PCC.depth
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.child)
WHERE PCC.parent IN (
    SELECT postCommentId
    FROM (
        SELECT postCommentId
        FROM PostComments
        WHERE PostComments.postId = $1
          AND parent = 0
        ORDER BY postCommentId DESC -- Newest first
        LIMIT $3 OFFSET $2
      ) tmp
  )
  AND PCC.depth <= $4
ORDER BY PC.postCommentId ASC -- Comments are flipped back for sorting
LIMIT $5;

-- name: CountPostCommentRepliesTotalDepth :one
SELECT COUNT(*)
FROM PostComments PC
  JOIN PostCommentClosures PCC ON (PC.postCommentId = PCC.parent)
WHERE PCC.parent = $1
  AND PCC.depth <= $2
  AND PCC.parent != PCC.child;

-- name: InsertPostCommentClosureDeep :exec
INSERT INTO PostCommentClosures(parent, child, depth)
SELECT p.parent,
  c.child,
  p.depth + c.depth + 1
FROM PostCommentClosures p,
  PostCommentClosures c
WHERE p.child = $1
  AND c.parent = $2;

-- name: InsertUserAvatar :exec
INSERT INTO UserAvatars (userID, created, fileType, hash)
VALUES ($1, $2, $3, $4);

-- name: DeactivateUserAvatar :exec
UPDATE UserAvatars
SET active = false
WHERE userID = $1
  AND active = true;

-- name: UserHasAvatar :one
SELECT EXISTS(
    SELECT *
    FROM UserAvatars
    WHERE userId = $1
      AND active = true
  );

-- name: DeleteUserAvatarByFile :exec
DELETE FROM UserAvatars
WHERE fileType = $1
  AND hash = $2;

-- name: CountPublicPostsByUser :one
SELECT COUNT(postId)
FROM Posts
WHERE published != 0
    AND userId = $1;

-- name: GetLatestPublicPostsByUser :many
SELECT postId,
    published,
    updated,
    content
FROM Posts
WHERE published != 0
    AND userId = $1
ORDER BY published DESC
LIMIT $3 OFFSET $2;

-- name: GetUserAvatarURLHash :one
SELECT hash
FROM UserAvatars
WHERE userId = $1
  AND active = true
ORDER BY userAvatarId DESC
LIMIT 1;

-- name: CountPublicPostsForUser :one
SELECT COUNT(postId)
FROM Posts
WHERE Posts.userId IN (
    SELECT followUserId
    FROM (
        SELECT followUserId
        FROM UserFollows
        WHERE UserFollows.userId = $1
          AND accepted = true
      ) tmp
  )
  AND published != 0
  OR Posts.userId = $1;

-- name: GetOldUserAvatars :many
SELECT UA.fileType,
  UA.hash
FROM UserAvatars AS UA
WHERE UA.userId = $1
  AND UA.active = false
  AND NOT EXISTS (
    SELECT *
    FROM UserAvatars AS UA2
    WHERE UA.hash = UA2.hash
      AND UA2.active = true
  );

-- name: GetLatestPublicPostsForUser :many
SELECT postId,
    userId,
    published,
    updated,
    content
FROM Posts
WHERE Posts.userId IN (
        SELECT followUserId
        FROM (
                SELECT followUserId
                FROM UserFollows
                WHERE UserFollows.userId = $1
                    AND accepted = true
            ) tmp
    )
    AND published != 0
    OR Posts.userId = $1
ORDER BY published DESC
LIMIT $3 OFFSET $2;

-- name: DeactivateOldUserAvatars :exec
UPDATE UserAvatars
SET active = false
WHERE userAvatarId IN (
        SELECT userAvatarId
        FROM (
            SELECT userAvatarId
            FROM UserAvatars
            ORDER BY userAvatarId DESC
            OFFSET 1
          ) tmp
    )
  AND UserAvatars.userId = $1
  AND active = true;

-- name: InsertSession :exec
INSERT INTO Sessions (
  key,
  data,
  expiration
)
VALUES ($1, $2, $3);

-- name: GetSession :one
SELECT data, expiration
FROM Sessions
WHERE key = $1
LIMIT 1;

-- name: SessionExist :one
SELECT EXISTS(
    SELECT *
    FROM Sessions
    WHERE key = $1
  );

-- name: UpdateSession :exec
UPDATE Sessions
SET data = $1,
  expiration = $2
WHERE key = $3;

-- name: DeleteSession :exec
DELETE FROM Sessions
WHERE key = $1;

-- name: DeleteOldSessions :exec
DELETE FROM Sessions
WHERE expiration < $1;