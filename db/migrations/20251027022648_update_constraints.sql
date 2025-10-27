-- migrate:up

ALTER TABLE UserAvatars DROP CONSTRAINT UserAvatars_userId_fkey;
ALTER TABLE UserAvatars ADD CONSTRAINT UserAvatars_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserEvents DROP CONSTRAINT UserEvents_userId_fkey;
ALTER TABLE UserEvents ADD CONSTRAINT UserEvents_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserKeys DROP CONSTRAINT UserKeys_userId_fkey;
ALTER TABLE UserKeys ADD CONSTRAINT UserKeys_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserEventKeys DROP CONSTRAINT UserEventKeys_userEventId_fkey;
ALTER TABLE UserEventKeys ADD CONSTRAINT UserEventKeys_userEventId_fkey FOREIGN KEY (userEventId) REFERENCES UserEvents (userEventId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserEventKeys DROP CONSTRAINT UserEventKeys_userKeyId_fkey;
ALTER TABLE UserEventKeys ADD CONSTRAINT UserEventKeys_userKeyId_fkey FOREIGN KEY (userKeyId) REFERENCES UserKeys (userKeyId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserFollows DROP CONSTRAINT UserFollows_userId_fkey;
ALTER TABLE UserFollows ADD CONSTRAINT UserFollows_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserFollows DROP CONSTRAINT UserFollows_followUserId_fkey;
ALTER TABLE UserFollows ADD CONSTRAINT UserFollows_followUserId_fkey FOREIGN KEY (followUserId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE UserTokens DROP CONSTRAINT UserTokens_userId_fkey;
ALTER TABLE UserTokens ADD CONSTRAINT UserTokens_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE Posts DROP CONSTRAINT Posts_userId_fkey;
ALTER TABLE Posts ADD CONSTRAINT Posts_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE PostMedia DROP CONSTRAINT PostMedia_postId_fkey;
ALTER TABLE PostMedia ADD CONSTRAINT PostMedia_postId_fkey FOREIGN KEY (postId) REFERENCES Posts (postId) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE PostMedia DROP CONSTRAINT PostMedia_postMediaFileId_fkey;
ALTER TABLE PostMedia ADD CONSTRAINT PostMedia_postMediaFileId_fkey FOREIGN KEY (postMediaFileId) REFERENCES PostMediaFiles (postMediaFileId) ON DELETE CASCADE ON UPDATE CASCADE;


-- migrate:down

ALTER TABLE UserAvatars DROP CONSTRAINT UserAvatars_userId_fkey;
ALTER TABLE UserAvatars ADD CONSTRAINT UserAvatars_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId);

ALTER TABLE UserEvents DROP CONSTRAINT UserEvents_userId_fkey;
ALTER TABLE UserEvents ADD CONSTRAINT UserEvents_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE;

ALTER TABLE UserKeys DROP CONSTRAINT UserKeys_userId_fkey;
ALTER TABLE UserKeys ADD CONSTRAINT UserKeys_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE;

ALTER TABLE UserEventKeys DROP CONSTRAINT UserEventKeys_userEventId_fkey;
ALTER TABLE UserEventKeys ADD CONSTRAINT UserEventKeys_userEventId_fkey FOREIGN KEY (userEventId) REFERENCES UserEvents (userEventId) ON DELETE CASCADE;

ALTER TABLE UserEventKeys DROP CONSTRAINT UserEventKeys_userKeyId_fkey;
ALTER TABLE UserEventKeys ADD CONSTRAINT UserEventKeys_userKeyId_fkey FOREIGN KEY (userKeyId) REFERENCES UserKeys (userKeyId) ON DELETE CASCADE;

ALTER TABLE UserFollows DROP CONSTRAINT UserFollows_userId_fkey;
ALTER TABLE UserFollows ADD CONSTRAINT UserFollows_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE;

ALTER TABLE UserFollows DROP CONSTRAINT UserFollows_followUserId_fkey;
ALTER TABLE UserFollows ADD CONSTRAINT UserFollows_followUserId_fkey FOREIGN KEY (followUserId) REFERENCES Users (userId) ON DELETE CASCADE;

ALTER TABLE UserTokens DROP CONSTRAINT UserTokens_userId_fkey;
ALTER TABLE UserTokens ADD CONSTRAINT UserTokens_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId) ON DELETE CASCADE;

ALTER TABLE Posts DROP CONSTRAINT Posts_userId_fkey;
ALTER TABLE Posts ADD CONSTRAINT Posts_userId_fkey FOREIGN KEY (userId) REFERENCES Users (userId);

ALTER TABLE PostMedia DROP CONSTRAINT PostMedia_postId_fkey;
ALTER TABLE PostMedia ADD CONSTRAINT PostMedia_postId_fkey FOREIGN KEY (postId) REFERENCES Posts (postId);

ALTER TABLE PostMedia DROP CONSTRAINT PostMedia_postMediaFileId_fkey;
ALTER TABLE PostMedia ADD CONSTRAINT PostMedia_postMediaFileId_fkey FOREIGN KEY (postMediaFileId) REFERENCES PostMediaFiles (postMediaFileId);
