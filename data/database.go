package data

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/netflixaddicts/Go-API/structs"
	"log"
)

func New() *Data {
	return &Data{
		Sheets: map[string]*structs.Sheet{},
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
	db.AutoMigrate(&structs.Sheet{})

	log.Print("Loading Sheets...")
	var sheets []*structs.Sheet
	db.Find(&sheets)
	for _, sheet := range sheets {
		d.Sheets[sheet.Name] = sheet
	}
	log.Printf("%d Sheets has been loaded!", len(d.Sheets))
}

func (d *Data) getSheets() ([]*structs.Sheet, error) {
	var sheets []*structs.Sheet
	d.db.Find(&sheets)

	return sheets, nil
}

func (d *Data) getSheet(id string) (structs.Sheet, error) {
	return structs.Sheet{}, nil
}

func (d *Data) addSheet(s structs.Sheet) error {
	return nil
}

func (d *Data) Close() {
	err := d.db.Close()
	if err != nil {
		panic(err)
	}
}

type Data struct {
	Sheets map[string]*structs.Sheet
	db     *gorm.DB
}
