// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
    "github.com/gogf/gf/os/gtime"
)

// BreakpointContinue is the golang structure for table breakpoint_continue.
type BreakpointContinue struct {
    Id         uint64      `orm:"id,primary"  json:"id"`         //   
    CreatedAt  *gtime.Time `orm:"created_at"  json:"createdAt"`  //   
    UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updatedAt"`  //   
    DeletedAt  *gtime.Time `orm:"deleted_at"  json:"deletedAt"`  //   
    FileMd5    string      `orm:"file_md5"    json:"fileMd5"`    //   
    FileName   string      `orm:"file_name"   json:"fileName"`   //   
    FilePath   string      `orm:"file_path"   json:"filePath"`   //   
    IsFinish   int         `orm:"is_finish"   json:"isFinish"`   //   
    ChunkTotal int64       `orm:"chunk_total" json:"chunkTotal"` //   
}