// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAmountDetail = "amount_details"

// AmountDetail mapped from table <amount_details>
type AmountDetail struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID         int64     `gorm:"column:user_id;not null" json:"user_id"`
	Type           string    `gorm:"column:type;not null;comment:提现 or 提成" json:"type"`                     // 提现 or 提成
	CorrelationID  int64     `gorm:"column:correlation_id;not null;comment:关联ID" json:"correlation_id"`     // 关联ID
	OriginalAmount string    `gorm:"column:original_amount;not null;comment:原始金额 分" json:"original_amount"` // 原始金额 分
	OperateAmount  string    `gorm:"column:operate_amount;not null;comment:操作金额" json:"operate_amount"`     // 操作金额
	CurrentAmount  string    `gorm:"column:current_amount;not null;comment:当前金额" json:"current_amount"`     // 当前金额
	Remarks        string    `gorm:"column:remarks;not null;comment:备注" json:"remarks"`                     // 备注
	Status         int32     `gorm:"column:status;not null;default:1;comment:1-正常" json:"status"`           // 1-正常
	CreateTime     time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime     time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
	IsDelete       int32     `gorm:"column:is_delete;not null" json:"is_delete"`
}

// TableName AmountDetail's table name
func (*AmountDetail) TableName() string {
	return TableNameAmountDetail
}
