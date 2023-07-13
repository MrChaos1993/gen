// Code generated by github.com/MrChaos1993/gen. DO NOT EDIT.
// Code generated by github.com/MrChaos1993/gen. DO NOT EDIT.
// Code generated by github.com/MrChaos1993/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePerson = "people"

// Person mapped from table <people>
type Person struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name           string         `gorm:"column:name" json:"name"`
	Age            int32          `gorm:"column:age" json:"age"`
	Flag           bool           `gorm:"column:flag" json:"flag"`
	AnotherFlag    int32          `gorm:"column:another_flag" json:"another_flag"`
	Commit         string         `gorm:"column:commit" json:"commit"`
	First          bool           `gorm:"column:First" json:"First"`
	Bit            []uint8        `gorm:"column:bit" json:"bit"`
	Small          int32          `gorm:"column:small" json:"small"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Score          float64        `gorm:"column:score" json:"score"`
	Number         int32          `gorm:"column:number" json:"number"`
	Birth          time.Time      `gorm:"column:birth;default:CURRENT_TIMESTAMP" json:"birth"`
	XMLHTTPRequest string         `gorm:"column:xmlHTTPRequest;default:' '" json:"xmlHTTPRequest"`
	JStr           string         `gorm:"column:jStr" json:"jStr"`
	Geo            string         `gorm:"column:geo" json:"geo"`
	Mint           int32          `gorm:"column:mint" json:"mint"`
	Blank          string         `gorm:"column:blank;default:' '" json:"blank"`
	Remark         string         `gorm:"column:remark" json:"remark"`
	LongRemark     string         `gorm:"column:long_remark" json:"long_remark"`
}

// TableName Person's table name
func (*Person) TableName() string {
	return TableNamePerson
}
