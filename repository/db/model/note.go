package model

// 定义 Note 结构体作为数据库模型
type Note struct {
	Rank    uint   `gorm:"type:uint"`
	Content string `gorm:"primaryKey;type:varchar(52)"` // 如果主键是数字类型，您可以使用 内联条件 来检索对象。 当使用字符串时，需要额外的注意来避免SQL注入；查看 Security 部分来了解详情。
	Done    bool   `gorm:"default:false"`
}
