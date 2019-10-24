package schemas

import "time"

type Schema struct {
	ID           int        `gorm:"column:id; type:int; not null; primary_key" json:"id"`
	CreatedAt    time.Time  `gorm:"column:created_at; type:datetime; not null; default now()" json:"createdAt"`
 	UpdatedAt    time.Time  `gorm:"column:updated_at; type:datetime; not null; default now()" json:"updatedAt"`
}