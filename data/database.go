package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/netflixaddicts/Go-API/structs"
)

func New() *Data {
	return &Data{
		Sheets: map[string]*structs.Sheet{},
	}
}

func (d *Data) Connect(user string, password string, addr string) {
	db, err := sql.Open("mysql", user+":"+password+"@tcp("+addr+":3306)/netflixaddicts?parseTime=true")
	if err != nil {
		panic(err)
	}
	d.db = db

	sheets, err := d.getSheets()
	if err != nil {
		panic(err)
	}
	for _, sheet := range sheets {
		d.Sheets[sheet.Name] = sheet
	}
}

func (d *Data) getSheets() ([]*structs.Sheet, error) {
	sheets := make([]*structs.Sheet, 0)
	result, err := d.db.Query("SELECT * FROM sheet")
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var sheet structs.Sheet
		err := result.Scan(&sheet.ID, &sheet.Name, &sheet.Title, &sheet.Image, &sheet.Note, &sheet.Styles, &sheet.Synopsis, &sheet.CreatedAt, &sheet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		sheets = append(sheets, &sheet)
	}
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
	db     *sql.DB
}
