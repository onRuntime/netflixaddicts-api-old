package structs

import (
	"time"
)

type Sheet struct {
	ID        int       `json:"id" gorm:"Type:INT(11); NOT NULL; AUTO_INCREMENT; PRIMARY_KEY"`
	Name      string    `json:"name" gorm:"Type:VARCHAR(64); NOT NULL;"`
	Title     string    `json:"title" gorm:"Type:VARCHAR(64); NOT NULL;"`
	Image     string    `json:"image" gorm:"Type:VARCHAR(255)"`
	Note      int       `json:"note" gorm:"Type:INT(11); NOT NULL; DEFAULT:'-1'"`
	Styles    []uint8   `json:"styles" gorm:"Type:ENUM('Humour', 'Effrayant', 'Comédie');"`
	Synopsis  string    `json:"synopsis" gorm:"Type:LONGTEXT; NOT NULL;"`
	CreatedAt time.Time `json:"created_at" gorm:"Type:DATETIME; NOT NULL; DEFAULT:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"Type:DATETIME; NOT NULL; DEFAULT:CURRENT_TIMESTAMP"`
}
