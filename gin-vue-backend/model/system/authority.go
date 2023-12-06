package system

import "time"

type SysAuthority struct {
	CreatedAt     time.Time     // 创建时间
	UpdatedAt     time.Time     // 更新时间
	DeletedAt     time.Time     // 删除时间
	AuthorityId   uint          `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string        `json:"authorityName" gorm:"comment:角色名"`
	SysBaseMenus  []SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus"`
}
