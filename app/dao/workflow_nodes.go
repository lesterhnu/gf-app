// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"gf-app/app/dao/internal"
)

// workflowNodesDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type workflowNodesDao struct {
	*internal.WorkflowNodesDao
}

var (
	// WorkflowNodes is globally public accessible object for table workflow_nodes operations.
	WorkflowNodes = &workflowNodesDao{
		internal.WorkflowNodes,
	}
)

// Fill with you ideas below.