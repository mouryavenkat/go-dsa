package main

type tweet struct {
	sequence int
	tweetID  int
	left     *tweet
	right    *tweet
}
type userInfo struct {
	followers  map[int]bool
	following  map[int]bool
	tweetCount int
	head       *tweet
	last       *tweet
	tweets     *tweet
}
type Twitter struct {
	UserMap     map[int]userInfo
	totalTweets int
}

func createTweetNode(tweetID int, sequence int) *tweet {
	return &tweet{
		sequence: sequence,
		tweetID:  tweetID,
	}
}

func pickTopTen(tweetHeads []*tweet, capacity int) (newhead, last *tweet, unfollowedCapacity int) {
	var nilCount int

	for capacity > 0 && nilCount < capacity {
		nilCount := 0
		best := -1
		bestIndex := -1
		for index, tweatHead := range tweetHeads {
			if tweatHead == nil {
				nilCount++
			} else if tweatHead.sequence > best {
				best = tweatHead.sequence
				bestIndex = index
			}
		}
		if bestIndex < 0 {
			break
		}
		bestNode := tweetHeads[bestIndex]
		node := createTweetNode(bestNode.tweetID, bestNode.sequence)
		if newhead == nil {
			newhead = node
			last = node
		} else {
			node.left = last
			last.right = node
			last = node
		}
		unfollowedCapacity++
		tweetHeads[bestIndex] = tweetHeads[bestIndex].right
		capacity--
	}
	return
}

func mergeLinkedLists(list1Copy, list2Copy *tweet, capacity int) (newhead, last *tweet, mergedCapacity int) {
	list1 := list1Copy
	list2 := list2Copy
	mergedCapacity = capacity
	for capacity > 0 && (list1 != nil || list2 != nil) {
		var bestNode *tweet
		if list2 == nil || (list1 != nil && list1.sequence > list2.sequence) {
			bestNode = createTweetNode(list1.tweetID, list1.sequence)
			list1 = list1.right
		} else {
			bestNode = createTweetNode(list2.tweetID, list2.sequence)
			list2 = list2.right
		}
		if newhead == nil {
			newhead = bestNode
			last = bestNode
		} else {
			last.right = bestNode
			bestNode.left = last
			last = bestNode
		}
		capacity--
	}
	return newhead, last, mergedCapacity - capacity
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		UserMap: make(map[int]userInfo),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.totalTweets++
	userData, exists := this.UserMap[userId]
	if !exists {
		userData = userInfo{
			followers: map[int]bool{userId: true},
		}
		this.UserMap[userId] = userData
	}

	for followingUser := range userData.followers {
		followingUserInfo := this.UserMap[followingUser]

		node := createTweetNode(tweetId, this.totalTweets)
		if followingUserInfo.head == nil {
			followingUserInfo.head = node
			followingUserInfo.last = node
			followingUserInfo.tweetCount++
		} else {
			if followingUserInfo.tweetCount == 10 {
				followingUserInfo.last = followingUserInfo.last.left
				followingUserInfo.last.right = nil
			} else {
				followingUserInfo.tweetCount++
			}

			node.right = followingUserInfo.head
			followingUserInfo.head.left = node
			followingUserInfo.head = node

		}
		this.UserMap[followingUser] = followingUserInfo
	}
	userData, _ = this.UserMap[userId]
	node := createTweetNode(tweetId, this.totalTweets)
	if userData.tweets == nil {
		userData.tweets = node
	} else {
		node.right = userData.tweets
		userData.tweets.left = node
		userData.tweets = node
	}
	this.UserMap[userId] = userData
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) (feed []int) {
	userInfo, _ := this.UserMap[userId]
	temp := userInfo.head
	for temp != nil {
		feed = append(feed, temp.tweetID)
		temp = temp.right
	}
	return
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	followerInfo, exists := this.UserMap[followerId]
	_, followerExists := followerInfo.following[followeeId]
	if followerExists {
		return
	}
	if !exists {
		followerInfo.followers = map[int]bool{followerId: true}

		this.UserMap[followerId] = followerInfo
	}
	followeeInfo, exists := this.UserMap[followeeId]
	if !exists {
		followeeInfo.followers = map[int]bool{followeeId: true}
		this.UserMap[followeeId] = followeeInfo

	}

	newHead, newLast, newCapacity := mergeLinkedLists(followerInfo.head, followeeInfo.tweets, 10)
	followerInfo.head = newHead
	followerInfo.last = newLast
	followerInfo.tweetCount = newCapacity
	if followerInfo.following == nil {
		followerInfo.following = map[int]bool{followeeId: true}
	} else {
		followerInfo.following[followeeId] = true
	}

	followeeInfo.followers[followerId] = true
	this.UserMap[followerId] = followerInfo
	this.UserMap[followeeId] = followeeInfo
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	followerInfo, exists := this.UserMap[followerId]
	if !exists {
		return
	}
	followeeInfo, exists := this.UserMap[followeeId]
	if !exists {
		return
	}
	tweetHeads := []*tweet{}

	for followingUser := range followerInfo.following {
		if followingUser == followeeId {
			continue
		}
		followingUserInfo, _ := this.UserMap[followingUser]
		tweetHeads = append(tweetHeads, followingUserInfo.tweets)
	}
	tweetHeads = append(tweetHeads, followerInfo.tweets)
	newhead, newLast, newCapacity := pickTopTen(tweetHeads, 10)
	followerInfo.head = newhead
	followerInfo.last = newLast
	followerInfo.tweetCount = newCapacity
	delete(followerInfo.following, followeeId)

	delete(followeeInfo.followers, followerId)
	this.UserMap[followeeId] = followeeInfo
	this.UserMap[followerId] = followerInfo
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
