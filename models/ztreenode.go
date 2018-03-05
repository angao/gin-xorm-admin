package models

// ZTreeNode ztree plugin
type ZTreeNode struct {
	ID      int64  `xorm:"id" json:"id"`
	Pid     int64  `xorm:"pId" json:"pId"`
	Name    string `json:"name"`
	Open    bool   `xorm:"open" json:"open"`
	Checked bool   `json:"checked"`
}
