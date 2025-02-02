package common

type NormalizeResp struct {
	Code    int32  `json:"code"`    // 状态码, 0-成功, 其他值-失败
	Message string `json:"message"` // 返回状态描述
}

type PaginateResp struct {
	Total     int64 `json:"total"`      // 总记录数
	Page      int64 `json:"page"`       // 页数
	PageSize  int64 `json:"page_size"`  // 每页记录数
	TotalPage int64 `json:"total_page"` // 总页数
}

type User struct {
	Id             int64  `json:"id"`              // 用户id
	Name           string `json:"name"`            // 用户名称
	FollowCount    int64  `json"follow_count`      // 关注总数
	FollowerCount  int64  `json:"follower_count"`  // 粉丝总数
	IsFollow       bool   `json:"is_follow"`       // true-已关注, false-未关注
	TotalFavorited int64  `json:"total_favorited"` // 获赞数量
	WorkCount      int64  `json:"work_count"`      // 作品数量
	FavoriteCount  int64  `json:"favorite_count"`  // 点赞数量
}

type Video struct {
	Id            int64  `json:"id"`            // 视频唯一标识
	PlayUrl       string `json:"play_url"`      // 视频播放地址
	CoverUrl      string `json:"cover_url"`     // 视频封面地址
	Title         string `json:"title"`         // 视频标题
	FavoriteCount int64  `json:"favorite_count` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count"` // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite"`   // true-已点赞, false-未点赞
	Author        User   `json:"author"`        // 视频作者信息
}

type VideoListResp struct {
	NormalizeResp
	PaginateResp
	VideoList Video `json:"video_list"` // 视频列表
}
