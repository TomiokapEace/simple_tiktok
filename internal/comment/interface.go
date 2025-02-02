package comment

type ICommentService interface {
	// 视频评论操作
	CommentAction(req *CommentActionReq) (*CommentActionResp, error)

	// 视频评论列表
	CommentList(req *CommentListReq) (*CommentListResp, error)
}
