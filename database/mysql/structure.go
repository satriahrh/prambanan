package mysql

type Base struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Province struct {
	*Base
}

type City struct {
	*Base
	ProvinceID string `json:"province_id"`
}

type District struct {
	*Base
	DistrictID string `json:"district_id"`
}
