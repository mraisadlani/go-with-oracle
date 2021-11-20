package domain

import "time"

type User struct {
	ID uint `json:"ID_USER"`
	NamaUser string `json:"NAMA_USER"`
	Password string `json:"PASSWORD"`
	Telp string `json:"TELP"`
	Email string `json:"EMAIL"`
	CreateBy *string `json:"CREATEBY"`
	CreateDate *time.Time `json:"CREATEDATE"`
	UpdateBy *string `json:"UPDATEBY"`
	UpdateDate *time.Time `json:"UPDATEDATE"`
}

type Instansi struct {
	ID uint `json:"ID_INSTANSI"`
	Instansi string `json:"INSTANSI"`
}

type Pemakaian struct {
	ID          string `json:"NO_REC_PAKAI"`
	Tanggal     string `json:"TANGGAL"`
	VolumePakai string `json:"VOLUME_PAKAI"`
	CreateDate  *string `json:"CREATE_DATE"`
	Status      *string `json:"STATUS"`
}