package models

// Relation role and menu
type Relation struct {
	ID     int64 `xorm:"pk id"`
	MenuID int64 `xorm:"menuid"`
	RoleID int64 `xorm:"roleid"`
}
