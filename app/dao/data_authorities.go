// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-app/app/dao/internal"
)

// dataAuthoritiesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dataAuthoritiesDao struct {
	*internal.DataAuthoritiesDao
}

var (
	// DataAuthorities is globally public accessible object for table data_authorities operations.
	DataAuthorities = &dataAuthoritiesDao{
		internal.DataAuthorities,
	}
)

// Fill with you ideas below.