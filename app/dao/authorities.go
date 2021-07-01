// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-app/app/dao/internal"
)

// authoritiesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type authoritiesDao struct {
	*internal.AuthoritiesDao
}

var (
	// Authorities is globally public accessible object for table authorities operations.
	Authorities = &authoritiesDao{
		internal.Authorities,
	}
)

// Fill with you ideas below.