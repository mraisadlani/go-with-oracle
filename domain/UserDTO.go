package domain

import (
	"Training/go-crud-with-oracle/technical_service/security/md5"
	"strings"
	"time"
)

type RegisterDTO struct {
	NamaUser string `json:"nama_user"`
	Password string `json:"password"`
	Telp string `json:"telp"`
	Email string `json:"email"`
	CreateBy string `json:"create_by"`
	CreateDate string `json:"create_date"`
}

func (d *RegisterDTO) SetDataUser(create RegisterDTO) *RegisterDTO {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Format("02-01-2006")

	d.NamaUser = strings.ToUpper(create.NamaUser)
	d.Password = md5.EncryptToMd5(create.Password)
	d.Telp = strings.Replace(create.Telp, "08", "+628", 1)
	d.Email = create.Email
	d.CreateBy = d.NamaUser
	d.CreateDate = now

	return d
}

type UpdateUserDTO struct {
	ID uint `json:"id"`
	Telp string `json:"telp"`
	Email string `json:"email"`
	UpdateBy string `json:"update_by"`
	UpdateDate string `json:"update_date"`
}

func (d *UpdateUserDTO) SetUpdateUser(create UpdateUserDTO) *UpdateUserDTO {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc).Format("02-01-2006")

	d.Telp = strings.Replace(create.Telp, "08", "+628", 1)
	d.Email = create.Email
	d.UpdateDate = now

	return d
}

type DateRangeDTO struct {
	DateFrom string `json:"date_from"`
	DateTo string `json:"date_to"`
}