package models

import (
	"github.com/guregu/null"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type LogAudit struct {
	ID            int         `gorm:"column:id;primary_key" json:"id"`
	CompanyId     int         `gorm:"column:company_id" json:"company_id"`
	LoginUsername string      `gorm:"column:login_username" json:"login_username"`
	LoginIp       string      `gorm:"column:login_ip" json:"login_ip"`
	Module        null.String `gorm:"column:module" json:"module"`
	AuditEvent    string      `gorm:"column:audit_event" json:"audit_event"`
	EventAttr     string      `gorm:"column:event_attr" json:"event_attr"`
	EventResult   string      `gorm:"column:event_result" json:"event_result"`
	InsertTm      time.Time   `gorm:"column:insert_tm" json:"insert_tm"`
}

func (m *LogAudit) TableName() string {
	return "tb_log_audit"
}
