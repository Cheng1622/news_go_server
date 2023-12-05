package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Cheng1622/news_go_server/internal/model"
	"github.com/Cheng1622/news_go_server/internal/model/requ"
	"github.com/Cheng1622/news_go_server/pkg/casbin"
	"github.com/Cheng1622/news_go_server/pkg/mysql"
)

type RoleService interface {
	GetRoles(req *requ.RoleListRequest) ([]model.Role, int64, error)     // 获取角色列表
	GetRolesByIds(roleIds []uint) ([]*model.Role, error)                 // 根据角色ID获取角色
	CreateRole(role *model.Role) error                                   // 创建角色
	UpdateRoleById(roleId uint, role *model.Role) error                  // 更新角色
	GetRoleMenusById(roleId uint) ([]*model.Menu, error)                 // 获取角色的权限菜单
	UpdateRoleMenus(role *model.Role) error                              // 更新角色的权限菜单
	GetRoleApisByRoleKeyword(roleKeyword string) ([]*model.Api, error)   // 根据角色关键字获取角色的权限接口
	UpdateRoleApis(roleKeyword string, reqRolePolicies [][]string) error // 更新角色的权限接口（先全部删除再新增）
	BatchDeleteRoleByIds(roleIds []uint) error                           // 删除角色
}

type Role struct{}

func NewRoleService() RoleService {
	return Role{}
}

// GetRoles 获取角色列表
func (rd Role) GetRoles(req *requ.RoleListRequest) ([]model.Role, int64, error) {
	var list []model.Role
	db := mysql.DB.Model(&model.Role{}).Order("created_at DESC")

	name := strings.TrimSpace(req.Name)
	if name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		db = db.Where("keyword LIKE ?", fmt.Sprintf("%%%s%%", keyword))
	}
	status := req.Status
	if status != 0 {
		db = db.Where("status = ?", status)
	}
	// 当pageNum > 0 且 pageSize > 0 才分页
	//记录总条数
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	pageNum := int(req.PageNum)
	pageSize := int(req.PageSize)
	if pageNum > 0 && pageSize > 0 {
		err = db.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}

// GetRolesByIds 根据角色ID获取角色
func (rd Role) GetRolesByIds(roleIds []uint) ([]*model.Role, error) {
	var list []*model.Role
	err := mysql.DB.Where("id IN (?)", roleIds).Find(&list).Error
	return list, err
}

// CreateRole 创建角色
func (rd Role) CreateRole(role *model.Role) error {
	err := mysql.DB.Create(role).Error
	return err
}

// UpdateRoleById 更新角色
func (rd Role) UpdateRoleById(roleId uint, role *model.Role) error {
	err := mysql.DB.Model(&model.Role{}).Where("id = ?", roleId).Updates(role).Error
	return err
}

// GetRoleMenusById 获取角色的权限菜单
func (rd Role) GetRoleMenusById(roleId uint) ([]*model.Menu, error) {
	var role model.Role
	err := mysql.DB.Where("id = ?", roleId).Preload("Menus").First(&role).Error
	return role.Menus, err
}

// UpdateRoleMenus 更新角色的权限菜单
func (rd Role) UpdateRoleMenus(role *model.Role) error {
	err := mysql.DB.Model(role).Association("Menus").Replace(role.Menus)
	return err
}

// GetRoleApisByRoleKeyword 根据角色关键字获取角色的权限接口
func (rd Role) GetRoleApisByRoleKeyword(roleKeyword string) ([]*model.Api, error) {
	policies := casbin.CasbinEnforcer.GetFilteredPolicy(0, roleKeyword)
	// 获取所有接口
	var apis []*model.Api
	err := mysql.DB.Find(&apis).Error
	if err != nil {
		return apis, errors.New("获取角色的权限接口失败")
	}

	accessApis := make([]*model.Api, 0)

	for _, policy := range policies {
		path := policy[1]
		method := policy[2]
		for _, api := range apis {
			if path == api.Path && method == api.Method {
				accessApis = append(accessApis, api)
				break
			}
		}
	}

	return accessApis, err

}

// UpdateRoleApis 更新角色的权限接口（先全部删除再新增）
func (rd Role) UpdateRoleApis(roleKeyword string, reqRolePolicies [][]string) error {
	// 先获取path中的角色ID对应角色已有的police(需要先删除的)
	err := casbin.CasbinEnforcer.LoadPolicy()
	if err != nil {
		return errors.New("角色的权限接口策略加载失败")
	}
	rmPolicies := casbin.CasbinEnforcer.GetFilteredPolicy(0, roleKeyword)
	if len(rmPolicies) > 0 {
		isRemoved, _ := casbin.CasbinEnforcer.RemovePolicies(rmPolicies)
		if !isRemoved {
			return errors.New("更新角色的权限接口失败")
		}
	}
	isAdded, _ := casbin.CasbinEnforcer.AddPolicies(reqRolePolicies)
	if !isAdded {
		return errors.New("更新角色的权限接口失败")
	}
	err = casbin.CasbinEnforcer.LoadPolicy()
	if err != nil {
		return errors.New("更新角色的权限接口成功，角色的权限接口策略加载失败")
	} else {
		return err
	}
}

// BatchDeleteRoleByIds 删除角色
func (rd Role) BatchDeleteRoleByIds(roleIds []uint) error {
	var roles []*model.Role
	err := mysql.DB.Where("id IN (?)", roleIds).Find(&roles).Error
	if err != nil {
		return err
	}
	err = mysql.DB.Select("Users", "Menus").Unscoped().Delete(&roles).Error
	// 删除成功就删除casbin policy
	if err == nil {
		for _, role := range roles {
			roleKeyword := role.Keyword
			rmPolicies := casbin.CasbinEnforcer.GetFilteredPolicy(0, roleKeyword)
			if len(rmPolicies) > 0 {
				isRemoved, _ := casbin.CasbinEnforcer.RemovePolicies(rmPolicies)
				if !isRemoved {
					return errors.New("删除角色成功, 删除角色关联权限接口失败")
				}
			}
		}

	}
	return err
}
