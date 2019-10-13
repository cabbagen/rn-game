package schemas

type User struct {
	Schema
	Username      string `gorm:"column:username; type:varchar(255); unique; not null" json:"username"`
	Nickname      string `gorm:"column:nickname; type:varchar(255); not null; default:\"\"" json:"nickname"`
	Password      string `gorm:"column:password; type:varchar(255); not null" json:"password"`
	Gender        int8   `gorm:"column:gender; type:smallint; not null;" json:"gender"`
	Avatar        string `gorm:"column:avatar; type:varchar(255); not null; default:\"\";" json:"avatar"`
	Score         int    `gorm:"column:score; type:int; not null; default: 0;" json:"score"`
	Extra         string `gorm:"column:extra; type:varchar(255); not null; default:\"\"" json:"extra"`
}
