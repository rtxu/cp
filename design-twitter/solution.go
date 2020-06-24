type Twitter struct {
    Tweets  []int
    Posts   map[int][]int
    F       map[int]map[int]bool    // F[A][B] = true, means A follow B
}


/** Initialize your data structure here. */
func Constructor() Twitter {
    return Twitter{
        Tweets : make([]int, 0),
        Posts : make(map[int][]int), 
        F     : make(map[int]map[int]bool),
    }
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
    seqId := len(this.Tweets)
    this.Tweets = append(this.Tweets, tweetId)
    this.Posts[userId] = append(this.Posts[userId], seqId)
}

type Node struct {
    seqIdx      int
    seqId       int
    fromUserId  int
}

// An IntHeap is a min-heap of ints.
type Heap []Node

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].seqId > h[j].seqId }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
    h := make(Heap, 0)
    push := func(user int) {
        postCnt := len(this.Posts[user])
        if postCnt > 0 {
            heap.Push(&h, Node{
                seqId: this.Posts[user][postCnt-1],
                seqIdx: postCnt-1,
                fromUserId: user,
            })
        }
    }
    followee := this.F[userId]
    if followee == nil {
        followee = make(map[int]bool)
    }
    followee[userId] = true
    for user, _ := range followee {
        push(user)
    }
    
    result := make([]int, 0)
    for i := 0; i < 10 && h.Len() > 0; i++ {
        top := heap.Pop(&h).(Node)
        result = append(result, this.Tweets[top.seqId])
        if top.seqIdx > 0 {
            heap.Push(&h, Node{
                seqId: this.Posts[top.fromUserId][top.seqIdx-1],
                seqIdx: top.seqIdx-1,
                fromUserId: top.fromUserId,
            })
        }
    }
    return result
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.F[followerId] == nil {
        this.F[followerId] = map[int]bool{}
    }
    this.F[followerId][followeeId] = true
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.F[followerId], followeeId)
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
