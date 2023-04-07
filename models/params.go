package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

//存放请求参数的结构体

type ParamsSignUp struct {
	UserName   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

//登陆结构体
type ParamsLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//ParamVoteData 投票数据
type ParamVoteData struct {
	//UserID 从请求中获取
	PostID    string `json:"post_id,string" binding:"required"`                //帖子id
	Direction int8   `json:"direction,string" binding:"required" oneof:1 0 -1` //赞成票(1)反对票(-1)取消票(0)
}

//获取帖子列表
type ParamArticleLists struct {
	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}
