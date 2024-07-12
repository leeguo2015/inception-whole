// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"inception-whole/internal/dao/internal"
)

// interactDao is the data access object for table gf_interact.
// You can define custom methods on it to extend its functionality as you wish.
type interactDao struct {
	*internal.InteractDao
}

var (
	// Interact is globally public accessible object for table gf_interact operations.
	Interact = interactDao{
		internal.NewInteractDao(),
	}
)

// Fill with you ideas below.
