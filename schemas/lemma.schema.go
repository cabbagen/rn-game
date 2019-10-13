package schemas

type Lemma struct {
	Schema
	Name      string    `gorm:"column:name; type:varchar(255); not null; unique" json:"name"`
	Link      string    `gorm:"column:link; type:varchar(255); not null; default:\"\"" json:"link"`
	Remark    string    `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
}
