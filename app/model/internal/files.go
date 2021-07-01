// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// Files is the golang structure for table files.
type Files struct {
    Id        uint64      `orm:"id,primary" json:"id"`        //           
    CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //           
    UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //           
    DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //           
    Url       string      `orm:"url"        json:"url"`       // 文件地址  
    Tag       string      `orm:"tag"        json:"tag"`       // 文件标签  
    Key       string      `orm:"key"        json:"key"`       // 编号      
    Name      string      `orm:"name"       json:"name"`      // 文件名    
}