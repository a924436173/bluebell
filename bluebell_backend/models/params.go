package models

// 定义请求参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username,omitempty" binding:"required"`
	Password   string `json:"password,omitempty" binding:"required"`
	RePassword string `json:"re_password,omitempty" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从请求中获取当前用户
	PostID    string `json:"post_id" binding:"required"`              // 帖子id
	Direction int8   `json:"direction,string" binding:"oneof=-1 0 1"` // 1 -> 赞成票  -1 -> 反对票 0 -> 取消投票
}

// ParamPostList 获取帖子列表query string参数
type ParamPostList struct {
	CommunityID int64  `json:"community_id" form:"community_id"`   // 可以为空
	Page        int64  `json:"page" form:"page" example:"1"`       // 页码
	Size        int64  `json:"size" form:"size" example:"10"`      // 每页数据量
	Order       string `json:"order" form:"order" example:"score"` // 排序依据
}
