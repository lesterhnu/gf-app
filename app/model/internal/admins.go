// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Admins is the golang structure for table admins.
type Admins struct {
    Id          int64       `orm:"id,primary"   json:"id"`          //               
    CreatedAt   *gtime.Time `orm:"created_at"   json:"createdAt"`   //               
    UpdatedAt   *gtime.Time `orm:"updated_at"   json:"updatedAt"`   //               
    DeletedAt   *gtime.Time `orm:"deleted_at"   json:"deletedAt"`   //               
    Uuid        string      `orm:"uuid"         json:"uuid"`        // 用户UUID      
    Avatar      string      `orm:"avatar"       json:"avatar"`      // 用户头像      
    Nickname    string      `orm:"nickname"     json:"nickname"`    // 用户登录名    
    Username    string      `orm:"username"     json:"username"`    // 用户昵称      
    Password    string      `orm:"password"     json:"password"`    // 用户登录密码  
    AuthorityId string      `orm:"authority_id" json:"authorityId"` // 用户角色ID    
}