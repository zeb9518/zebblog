// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Message string `json:"message"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	//Smscode  string `json:"smscode"` //短信码
}

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRep struct {
	Auid     int64  `json:"auid"`
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
	Nickname string `json:"nickname"`
	Openid   string `json:"openid"`
	Avator   string `json:"avator"`
}

type RegisterRep struct {
	Auid int64 `json:"auid"`
	Uid  int64 `json:"uid"`
}