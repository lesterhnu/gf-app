// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal



// AuthoritiesMenus is the golang structure for table authorities_menus.
type AuthoritiesMenus struct {
    AuthorityId string `orm:"authority_id,primary" json:"authorityId"` // 角色ID  
    MenuId      uint64 `orm:"menu_id,primary"      json:"menuId"`      //         
}