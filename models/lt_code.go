package models

type LtCode struct {
	Id         int    `xorm:"not null pk autoincr INT(10)"`
	GiftId     int    `xorm:"not null default 0 comment('奖品ID') index INT(10)"`
	Code       string `xorm:"not null default '' comment('虚拟券编码') unique VARCHAR(255)"`
	SysCreated int    `xorm:"not null default 0 comment('创建时间') INT(10)"`
	SysUpdated int    `xorm:"not null default 0 comment('更新时间') INT(10)"`
	SysStatus  int    `xorm:"not null default 0 comment('状态 0：正常 1：作废 2：已发送') TINYINT(3)"`
}
