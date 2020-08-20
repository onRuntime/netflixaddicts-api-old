package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func New() *Data {
	return &Data{
		Sheets: map[string]*Sheet{},
	}
}

func (d *Data) Connect(addr string, user string, password string) {
	log.Print("Connecting to database...")
	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+addr+":3306)/netflixaddicts?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	log.Printf("Database connected!")
	d.db = db
	db.AutoMigrate(&Sheet{})

	log.Print("Loading Sheets...")
	var sheets []*Sheet
	db.Find(&sheets)
	for _, sheet := range sheets {
		d.Sheets[sheet.Name] = sheet
	}
	log.Printf("%d Sheets has been loaded!", len(d.Sheets))
}

func (d *Data) getSheets() ([]*Sheet, error) {
	var sheets []*Sheet
	d.db.Find(&sheets)

	return sheets, nil
}

func (d *Data) getSheet(id string) (Sheet, error) {
	return Sheet{}, nil
}

func (d *Data) addSheet(s Sheet) error {
	return nil
}

func (d *Data) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}
}

type Data struct {
	Sheets map[string]*Sheet
	db     *gorm.DB
}

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
