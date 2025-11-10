package user

import (
	"context"
	"errors"
	"time"

	"github.com/allocamelus/allocamelus/internal/db"
	"github.com/allocamelus/allocamelus/internal/g"
	"github.com/allocamelus/allocamelus/internal/pkg/compare"
	"github.com/jackc/pgx/v5"
)

// Following is userId following followUserId
func Following(userId, followUserId int64) (follow FollowStruct, err error) {
	if compare.EqualInt(userId, followUserId) {
		follow.Following = true
		return
	}

	follow.Following, err = g.Data.Queries.IsFollowing(context.Background(), db.IsFollowingParams{
		Userid:       userId,
		Followuserid: followUserId,
	})
	if !errors.Is(err, pgx.ErrNoRows) {
		if err != nil {
			return
		}
		if !follow.Following {
			follow.Requested = true
		}
	}
	return follow, nil
}

// ListFollowing return a slice of userId(s) that userId follows
func ListFollowing(userId int64) ([]int64, error) {
	return g.Data.Queries.ListFollowing(context.Background(), userId)
}

func FollowExt(userId int64, followUserId int64, accepted bool) error {
	following, err := Following(userId, followUserId)
	if err != nil || following.Following {
		// return if following silently
		return err
	}

	return g.Data.Queries.FollowUser(context.Background(), db.FollowUserParams{
		Userid:       userId,
		Followuserid: followUserId,
		Accepted:     accepted,
		Created:      time.Now().Unix(),
	})
}

// Follow userId follow followUserId
func Follow(userId int64, followUserId int64) error {
	t, err := GetType(followUserId)
	if err != nil {
		return err
	}

	// Only Public Users are auto followed
	if t.Public() {
		return FollowExt(userId, followUserId, true)
	}
	// Has the user being followed requested to follow user
	follow, err := Following(followUserId, userId)
	if err != nil {
		return err
	}
	// if so Accept request
	if follow.Requested {
		return Accept(userId, followUserId)
	}
	// else request to follow
	return FollowExt(userId, followUserId, false)
}

// Accept userId Accept followerUserId request
func Accept(userId, followerUserId int64) error {
	// Is follower following already
	follow, err := Following(followerUserId, userId)
	if err != nil || follow.Following {
		// return if following silently
		return err
	}

	err = g.Data.Queries.AcceptFollowRequest(context.Background(), db.AcceptFollowRequestParams{
		Userid:       followerUserId,
		Followuserid: userId,
	})
	if err != nil {
		return err
	}

	t, err := GetType(userId)
	if err != nil {
		return err
	}

	// Auto Follow back is like fb friends
	if t.Private() {
		t, err = GetType(followerUserId)
		if err != nil {
			return err
		}
		// Is follower Private? Follow if so
		if t.Private() {
			return FollowExt(userId, followerUserId, true)
		}
	}
	return nil
}

// Decline userId Decline followerUserId request
func Decline(userId, followerUserId int64) error {
	return unfollowDB(followerUserId, userId)
}

// AcceptAll userId Accept all follow request
// Called when accounts goes from private to public
func AcceptAll(userId int64) error {
	return g.Data.Queries.AcceptAllFollowRequests(context.Background(), userId)
}

// unfollowDB userId unfollow followUserId
func unfollowDB(userId, followUserId int64) error {
	follow, err := Following(userId, followUserId)
	if err != nil || (!follow.Following && !follow.Requested) {
		// return if !following silently
		return err
	}

	return g.Data.Queries.UnfollowUser(context.Background(), db.UnfollowUserParams{
		Userid:       userId,
		Followuserid: followUserId,
	})
}

func Unfollow(userId, unfollowId int64) error {
	if err := unfollowDB(userId, unfollowId); err != nil {
		return err
	}

	// Unfriend if unfollowId and userId are private
	if t, err := GetType(unfollowId); err != nil || !t.Private() {
		return err
	}
	if t, err := GetType(userId); err != nil || !t.Private() {
		return err
	}

	return unfollowDB(unfollowId, userId)
}

// Followers count userId followers
func Followers(userId int64) (int64, error) {
	return g.Data.Queries.CountFollowers(context.Background(), userId)
}

// ListFollows return a slice of userId(s) that follow userId
func ListFollowers(userId int64) ([]int64, error) {
	return g.Data.Queries.ListFollowers(context.Background(), userId)
}

// ListRequests return a ordered map of userId(s) that request to follow userId
// TODO limit
func ListRequests(userId int64) (map[int64]int64, error) {
	requests, err := g.Data.Queries.ListFollowRequests(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	result := map[int64]int64{}
	for i, requester := range requests {
		result[int64(i)] = requester
	}
	return result, nil
}
