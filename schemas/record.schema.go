package schemas

type Record struct {
	Schema
	UserId               int      `gorm:"column:user_id; type:int; not null" json:"userId"`
	GameId               int      `gorm:"column:game_id; type:int; not null; unique key" json:"gameId"`
	Remark               string   `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
}

type RecordDetail struct {
	Id                   int       `gorm:"column:id" json:"id"`
	Username             string    `gorm:"column:username" json:"username"`
	Nickname             string    `gorm:"column:nickname" json:"nickname"`
	Gender               int8      `gorm:"column:gender" json:"gender"`
	Avatar               string    `gorm:"column:avatar" json:"avatar"`
	GameId               int       `gorm:"column:gameId" json:"gameId"`
	GameImg              string    `gorm:"column:gameImg" json:"gameImg"`
	GameName             string    `gorm:"column:gameName" json:"gameName"`
	GamePlayLink         string    `gorm:"column:gamePlayLink" json:"gamePlayLink"`
	GameCategoryName     string    `gorm:"column:gameCategoryName" json:"gameCategoryName"`                 
}
