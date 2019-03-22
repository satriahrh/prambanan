package mysql

import (
	"context"
	"database/sql"
	"github.com/prometheus/common/log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/helloproclub/prambanan/errors"
)

type MySql struct {
	Database *sql.DB
}

func NewMySql(dataSourceName string) (mySql *MySql, err error) {
	db, err := sql.Open("mysql", dataSourceName)
	mySql = &MySql{
		Database: db,
	}
	return
}

func (m *MySql) GetProvinceByID(ctx context.Context, id string) (province string, err error) {
	if err = isID(id); err != nil {
		return
	}

	query := `SELECT * FROM provinces WHERE id = ?;`

	rows, err := m.Database.QueryContext(ctx, query, id)
	if err != nil {
		log.Errorln(err)
		return
	}

	var provinceData Province
	if err = rows.Scan(&provinceData); err != nil {
		return
	}

	province = provinceData.Name

	return
}

func (m *MySql) GetCityByID(ctx context.Context, provinceID, id string) (city string, err error) {
	return
}

func (m *MySql) GetDistrictByID(ctx context.Context, cityID, id string) (district string, err error) {
	return
}

func isID(id string) (err error) {
	idBytes := []byte(id)
	for _, idByte := range idBytes {
		if 48 <= idByte && idByte <= 57 {
			err = errors.ErrInvalidID
		}
	}
	return
}
