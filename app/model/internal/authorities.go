// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Authorities is the golang structure for table authorities.
type Authorities struct {
    CreatedAt     *gtime.Time `orm:"created_at"           json:"createdAt"`     //           
    UpdatedAt     *gtime.Time `orm:"updated_at"           json:"updatedAt"`     //           
    DeletedAt     *gtime.Time `orm:"deleted_at"           json:"deletedAt"`     //           
    ParentId      string      `orm:"parent_id"            json:"parentId"`      // 父角色ID  
    AuthorityId   string      `orm:"authority_id,primary" json:"authorityId"`   // 角色ID    
    AuthorityName string      `orm:"authority_name"       json:"authorityName"` // 角色名    
    DefaultRouter string      `orm:"default_router"       json:"defaultRouter"` // 默认菜单  
}