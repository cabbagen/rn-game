package schemas

type Record struct {
	Schema
	UserId      int      `gorm:"column:user_id; type:int; not null" json:"userId"`
	Status      int8     `gorm:"column:status; type:smallint; not null; default:1" json:"status"`
	Score       int      `gorm:"column:score; type:int; not null; default:0" json:"score"`
	Lemmata     string   `gorm:"column:lemmata; type:varchar(255); not null; default:\"\"" json:"lemmata"`
	Remark      string   `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
}
