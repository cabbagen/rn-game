package schemas

type Category struct {
	Schema
	Name        string `gorm:"column:name; type:varchar(255); not null; default:\"\"" json:"name"`
	Status      int8   `gorm:"column:status; type:smallint; not null; default:1" json:"status"`
	Remark      string `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
}
