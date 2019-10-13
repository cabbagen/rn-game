package schemas

import "time"

type Game struct {
	Schema
	Img             string       `gorm:"column:img; type:varchar(255); not null; default:\"\"" json:"img"`
	Name            string       `gorm:"column:name; type:varchar(255); not null; default:\"\"" json:"name"`
	Size            string       `gorm:"column:size; type:varchar(255); not null; default:\"\"" json:"size"`
	Status          int8         `gorm:"column:status; type:smallint; not null; default:1" json:"status"`
	Developer       string       `gorm:"column:developer; type:varchar(255); not null; default:\"\"" json:"developer"`
	CategoryId      int          `gorm:"column:category_id; type:int; not null" json:"categoryId"`
	Download        int          `gorm:"column:download; type:int; not null; default:0" json:"download"`
	Description      string      `gorm:"column:description; type:text; not null" json:"description"`
	Remark          string       `gorm:"column:remark; type:varchar(255); not null; default:\"\"" json:"remark"`
	PublishDate     time.Time    `gorm:"column:publish_date; type:datetime" json:"publishDate"`
}

type GameDetail struct {
	Game
	CategoryName    string       `gorm:"column:categoryName" json:"categoryName"`
}