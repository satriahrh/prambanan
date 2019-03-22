package prambanan

import (
	"context"
)

// Prambanan main struct of the project
type Prambanan struct {
	Database Database
}

// Database interface of the database
type Database interface {
	GetProvinceByID(ctx context.Context, id string) (province string, err error)
	GetCityByID(ctx context.Context, provinceID, id string) (city string, err error)
	GetDistrictByID(ctx context.Context, cityID, id string) (district string, err error)
}

// Result hold the result of nik decoder
type Result struct {
	Province   string `json:"provinsi"`
	City       string `json:"kota"`
	District   string `json:"kecamatan"`
	Gender     string `json:"jenis_kelamin"`
	BirthDate  string `json:"tanggal_lahir"`
	UniqueCode string `json:"kode_unik"`
}
