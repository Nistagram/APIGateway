package dto

type FeedRequestDTO struct {
	FollowedUsers []uint64
	MutedUsers    []uint64
}
