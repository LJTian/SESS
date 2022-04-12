package model

import (
	"gorm.io/gorm"
	"time"
)

// BaseModelStart 基础模板头部
type BaseModelStart struct {
	ID uint `gorm:"primarykey"`
}

// BaseModeEnd 基础模板尾部
type BaseModeEnd struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

/*
1. 密文 2. 密文不可反解
	1. 对称加密
	2. 非对称加密
	3. md5 信息摘要算法
	密码如果不可以反解，用户找回密码
*/

type User struct {
	BaseModelStart
	Mobile   string     `gorm:"column:mobile;index:idx_mobile;unique;type:varchar(11);not null comment '手机号'"`
	NickName string     `gorm:"column:nick_name;type:varchar(20) comment '昵称'"`
	PassWord string     `gorm:"column:pass_word;type:varchar(100) comment '加密后的密码'"`
	Birthday *time.Time `gorm:"column:birthday;type:datetime comment '生日'"`
	Gender   string     `gorm:"column:gender;default:0;type:varchar(1) comment '性别 0:男 1:女'"`
	Role     string     `gorm:"column:role;default:1;type:varchar(1) comment '1:普通用户 2:管理员'"`
	BaseModeEnd
}
