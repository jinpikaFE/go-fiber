package models

type Districts struct {
	citycode  string
	adcode    string
	name      string
	level     string
	districts []*Districts
}

type Region struct {
	status    string
	info      string
	infocode  string
	count     string
	districts []*Districts
}

