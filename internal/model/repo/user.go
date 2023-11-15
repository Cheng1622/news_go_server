package repo

import "github.com/Cheng1622/news_go_server/internal/model"

// 返回给前端的当前用户信息
type UserInfoResp struct {
	Userid       int64         `json:"userid,string"`
	Username     string        `json:"username"`
	Mobile       string        `json:"mobile"`
	Avatar       string        `json:"avatar"`
	Nickname     string        `json:"nickname"`
	Introduction string        `json:"introduction"`
	Roles        []*model.Role `json:"roles"`
}

func ToUserInfoResp(user model.User) UserInfoResp {
	return UserInfoResp{
		Userid:       user.Userid,
		Username:     user.Username,
		Mobile:       user.Mobile,
		Avatar:       user.Avatar,
		Nickname:     *user.Nickname,
		Introduction: *user.Introduction,
		Roles:        user.Roles,
	}
}

// 返回给前端的用户列表
type UsersResp struct {
	Userid       int64   `json:"userid,string"`
	Username     string  `json:"username"`
	Mobile       string  `json:"mobile"`
	Avatar       string  `json:"avatar"`
	Nickname     string  `json:"nickname"`
	Introduction string  `json:"introduction"`
	Status       uint    `json:"status"`
	Creator      string  `json:"creator"`
	RoleIds      []int64 `json:"roleIds"`
}

func ToUsersResp(userList []*model.User) []UsersResp {
	var users []UsersResp
	for _, user := range userList {
		userResp := UsersResp{
			Userid:       user.Userid,
			Username:     user.Username,
			Mobile:       user.Mobile,
			Avatar:       user.Avatar,
			Nickname:     *user.Nickname,
			Introduction: *user.Introduction,
			Status:       user.Status,
			Creator:      user.Creator,
		}
		roleIds := make([]int64, 0)
		for _, role := range user.Roles {
			roleIds = append(roleIds, int64(role.ID))
		}
		userResp.RoleIds = roleIds
		users = append(users, userResp)
	}

	return users
}
