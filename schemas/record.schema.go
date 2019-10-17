package schemas

type Record struct {
	Schema
	UserId      int      `gorm:"column:user_id; type:int; not null" json:"userId"`
	GameId      int      `gorm:"column:game_id; type:int; not null" json:"gameId"`
	Remark      string   `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
}
