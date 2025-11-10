-- migrate:up
CREATE TABLE Users (
  userId BIGSERIAL PRIMARY KEY,
  userName VARCHAR(64) NOT NULL UNIQUE,
  name VARCHAR(128) NOT NULL,
  email VARCHAR(254) NOT NULL UNIQUE,
  bio VARCHAR(255) NOT NULL,
  type SMALLINT NOT NULL DEFAULT 0,
  permissions BIGINT NOT NULL DEFAULT 0,
  created INTEGER NOT NULL -- Time
);

CREATE INDEX Users_created_idx ON Users (created);

CREATE TABLE UserAvatars (
  userAvatarId BIGSERIAL PRIMARY KEY,
  userId BIGINT NOT NULL,
  active BOOLEAN NOT NULL DEFAULT TRUE, -- Bool
  created INTEGER NOT NULL, -- Time
  fileType INTEGER NOT NULL,
  hash VARCHAR(128) NOT NULL,
  CONSTRAINT UserAvatars_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId)
);

CREATE INDEX UserAvatars_userId_idx ON UserAvatars (userId);
CREATE INDEX UserAvatars_active_idx ON UserAvatars (active);
CREATE INDEX UserAvatars_fileType_idx ON UserAvatars (fileType);
CREATE INDEX UserAvatars_hash_idx ON UserAvatars (hash);

CREATE TABLE UserEvents (
  userEventId SERIAL PRIMARY KEY,
  userId BIGINT NOT NULL,
  eventType SMALLINT NOT NULL,
  created INTEGER NOT NULL, -- Time
  info VARCHAR(4096) NOT NULL, -- Encrypted Bytes
  CONSTRAINT UserEvents_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE
);

CREATE INDEX UserEvents_userId_idx ON UserEvents (userId);
CREATE INDEX UserEvents_created_idx ON UserEvents (created);

CREATE TABLE UserKeys (
  userKeyId BIGSERIAL PRIMARY KEY,
  userId BIGINT NOT NULL,
  created INTEGER NOT NULL, -- Time
  replaced INTEGER NOT NULL DEFAULT 0, -- Time
  authKeyHash VARCHAR(128) NOT NULL,
  authKeySalt VARCHAR(128) NOT NULL,
  publicArmored VARCHAR(2048) NOT NULL,
  privateArmored VARCHAR(4096) NOT NULL, -- Encrypted
  recoveryKeyHash VARCHAR(128) NOT NULL,
  recoveryArmored VARCHAR(4096) NOT NULL, -- Encrypted w/ Recovery Key
  CONSTRAINT UserKeys_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE
);

CREATE INDEX UserKeys_userId_idx ON UserKeys (userId);
CREATE INDEX UserKeys_created_idx ON UserKeys (created);
CREATE INDEX UserKeys_replaced_idx ON UserKeys (replaced);
CREATE INDEX UserKeys_recoveryKeyHash_idx ON UserKeys (recoveryKeyHash);

CREATE TABLE UserEventKeys (
  userEventKeyId BIGSERIAL PRIMARY KEY,
  userEventId INTEGER NOT NULL,
  userKeyId BIGINT NOT NULL,
  infoKey VARCHAR(1024) NOT NULL, -- Encrypted AES Key w/ Public key
  CONSTRAINT UserEventKeys_userEventId_fkey FOREIGN KEY (userEventId) REFERENCES UserEvents (userEventId) ON DELETE CASCADE,
  CONSTRAINT UserEventKeys_userKeyId_fkey FOREIGN KEY (userKeyId) REFERENCES UserKeys (userKeyId) ON DELETE CASCADE
);

CREATE INDEX UserEventKeys_userEventId_idx ON UserEventKeys (userEventId);
CREATE INDEX UserEventKeys_userKeyId_idx ON UserEventKeys (userKeyId);

CREATE TABLE UserFollows (
  userFollowId BIGSERIAL PRIMARY KEY,
  userId BIGINT NOT NULL, -- Follower
  followUserId BIGINT NOT NULL, -- Following
  accepted BOOLEAN NOT NULL, -- Bool
  created INTEGER NOT NULL, -- Time
  CONSTRAINT UserFollows_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE,
  CONSTRAINT UserFollows_followUserId_fkey FOREIGN KEY (followUserId) REFERENCES Users (userId) ON DELETE CASCADE
);

CREATE INDEX UserFollows_userId_idx ON UserFollows (userId);
CREATE INDEX UserFollows_followUserId_idx ON UserFollows (followUserId);

CREATE TABLE UserTokens (
  userTokenId BIGSERIAL PRIMARY KEY,
  userId BIGINT NOT NULL,
  tokenType SMALLINT NOT NULL,
  selector VARCHAR(12) NOT NULL,
  token VARCHAR(128) NOT NULL, -- Hashed
  created INTEGER NOT NULL, -- Time
  expiration INTEGER NOT NULL, -- Time
  CONSTRAINT UserTokens_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE
);

CREATE INDEX UserTokens_userId_idx ON UserTokens (userId);
CREATE INDEX UserTokens_selector_idx ON UserTokens (selector);
CREATE INDEX UserTokens_expiration_idx ON UserTokens (expiration);

CREATE TABLE Posts (
  postId BIGSERIAL PRIMARY KEY,
  userId BIGINT NOT NULL,
  created INTEGER NOT NULL, -- Time
  published INTEGER NOT NULL, -- Time
  updated INTEGER NOT NULL DEFAULT 0, -- Time
  content TEXT NOT NULL,
  CONSTRAINT Posts_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId)
);

CREATE INDEX Posts_userId_idx ON Posts (userId);

CREATE TABLE PostComments (
  postCommentId BIGSERIAL PRIMARY KEY,
  postId BIGINT NOT NULL,
  userId BIGINT NOT NULL,
  parent BIGINT NOT NULL DEFAULT 0, -- 0 or postCommentId
  created INTEGER NOT NULL, -- Time
  updated INTEGER NOT NULL DEFAULT 0, -- Time
  content VARCHAR(4096) NOT NULL,
  CONSTRAINT PostComments_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT PostComments_postId_fkey FOREIGN KEY (postId) REFERENCES Posts (postId) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX PostComments_parent_idx ON PostComments (parent);
CREATE INDEX PostComments_userId_idx ON PostComments (userId);
CREATE INDEX PostComments_postId_idx ON PostComments (postId);

CREATE TABLE PostCommentClosures (
  parent BIGINT NOT NULL,
  child BIGINT NOT NULL,
  depth BIGINT NOT NULL,
  PRIMARY KEY (parent, child),
  UNIQUE (parent, depth, child),
  UNIQUE (child, parent, depth),
  CONSTRAINT PostCommentClosures_parent_fkey FOREIGN KEY (parent) REFERENCES PostComments (postCommentId) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT PostCommentClosures_child_fkey FOREIGN KEY (child) REFERENCES PostComments (postCommentId) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE INDEX PostCommentClosures_parent_idx ON PostCommentClosures (parent);
CREATE INDEX PostCommentClosures_child_idx ON PostCommentClosures (child);
CREATE INDEX PostCommentClosures_depth_idx ON PostCommentClosures (depth);

CREATE TABLE PostMediaFiles (
  postMediaFileId BIGSERIAL PRIMARY KEY,
  created INTEGER NOT NULL, -- Time
  fileType INTEGER NOT NULL,
  width INTEGER NOT NULL,
  height INTEGER NOT NULL,
  hash VARCHAR(128) NOT NULL UNIQUE,
  newHash VARCHAR(128) NOT NULL UNIQUE -- Hash after imagedit
);

CREATE TABLE PostMedia (
  postMediaId BIGSERIAL PRIMARY KEY,
  postId BIGINT NOT NULL,
  added INTEGER NOT NULL, -- Time
  active BOOLEAN NOT NULL DEFAULT TRUE, -- Bool
  alt VARCHAR(512) NOT NULL,
  postMediaFileId BIGINT NOT NULL DEFAULT 1,
  CONSTRAINT PostMedia_postId_fkey FOREIGN KEY (postId) REFERENCES Posts (postId),
  CONSTRAINT PostMedia_postMediaFileId_fkey FOREIGN KEY (postMediaFileId) REFERENCES PostMediaFiles (postMediaFileId)
);

CREATE INDEX PostMedia_postId_idx ON PostMedia (postId);
CREATE INDEX PostMedia_active_idx ON PostMedia (active);
CREATE INDEX PostMedia_postMediaFileId_idx ON PostMedia (postMediaFileId);

-- migrate:down

