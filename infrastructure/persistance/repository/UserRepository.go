package repository

import (
	"Training/go-crud-with-oracle/domain"
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
)

type UserRepository struct {
	db *sql.DB
}

func BuildUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

var _ domain.IUserRepository = &UserRepository{}

func (r *UserRepository) Create(registerDTO *domain.RegisterDTO) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := r.db.PrepareContext(ctx, `INSERT INTO m_users (NAMA_USER, PASSWORD, TELP, EMAIL, CREATEBY, CREATEDATE) VALUES (:1, :2, :3, :4, :5, to_date(:6,'DD-MM-YYYY'))`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.ExecContext(ctx,
		registerDTO.NamaUser,
		registerDTO.Password,
		registerDTO.Telp,
		registerDTO.Email,
		registerDTO.CreateBy,
		registerDTO.CreateDate)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) FindById(id uint64) (*domain.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.PrepareContext(ctx, `SELECT ID_USER, NAMA_USER, PASSWORD, TELP, EMAIL, CREATEBY, CREATEDATE, UPDATEBY, UPDATEDATE FROM m_users WHERE ID_USER=:1 ORDER BY ID_USER ASC`)
	defer query.Close()

	if err != nil {
		return nil, err
	}

	var user domain.User
	query.QueryRowContext(ctx, id).Scan(
		&user.ID,
		&user.NamaUser,
		&user.Password,
		&user.Telp,
		&user.Email,
		&user.CreateBy,
		&user.CreateDate,
		&user.UpdateBy,
		&user.UpdateDate,
	)

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.PrepareContext(ctx, `SELECT ID_USER, NAMA_USER, PASSWORD, TELP, EMAIL, CREATEBY, CREATEDATE, UPDATEBY, UPDATEDATE FROM m_users WHERE EMAIL=:1 ORDER BY ID_USER ASC`)
	defer query.Close()

	if err != nil {
		return nil, err
	}

	var user domain.User
	query.QueryRowContext(ctx, email).Scan(
		&user.ID,
		&user.NamaUser,
		&user.Password,
		&user.Telp,
		&user.Email,
		&user.CreateBy,
		&user.CreateDate,
		&user.UpdateBy,
		&user.UpdateDate,
		)

	return &user, nil
}

func (r *UserRepository) FindAll() (*[]domain.User, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.QueryContext(ctx, `SELECT ID_USER, NAMA_USER, PASSWORD, TELP, EMAIL, CREATEBY, CREATEDATE, UPDATEBY, UPDATEDATE FROM m_users ORDER BY ID_USER ASC`)
	defer query.Close()

	if err != nil {
		return nil, err
	}

	var users []domain.User

	for query.Next() {
		var user domain.User

		err = query.Scan(
			&user.ID,
			&user.NamaUser,
			&user.Password,
			&user.Telp,
			&user.Email,
			&user.CreateBy,
			&user.CreateDate,
			&user.UpdateBy,
			&user.UpdateDate)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return &users, nil
}

func (r *UserRepository) UpdateById(id uint64, updateDTO *domain.UpdateUserDTO) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := r.db.PrepareContext(ctx, `UPDATE m_users
		SET TELP=:1, EMAIL=:2, UPDATEBY=:3, UPDATEDATE=to_date(:4,'DD-MM-YYYY') WHERE ID_USER=:5`)
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.ExecContext(ctx,
		updateDTO.Telp,
		updateDTO.Email,
		updateDTO.UpdateBy,
		updateDTO.UpdateDate,
		id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) DeleteById(id uint64) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	row, err := r.db.PrepareContext(ctx, "DELETE FROM m_users WHERE ID_USER=:1")
	defer row.Close()

	if err != nil {
		return false, err
	}

	_, err = row.ExecContext(ctx, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) GetDataView() (*[]domain.Instansi, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query, err := r.db.QueryContext(ctx, "SELECT * FROM view_instansi ORDER BY ID_INSTANSI ASC")
	defer query.Close()

	if err != nil {
		return nil, err
	}

	var instansi []domain.Instansi

	for query.Next() {
		var instan domain.Instansi

		err = query.Scan(
			&instan.ID,
			&instan.Instansi)

		if err != nil {
			return nil, err
		}

		instansi = append(instansi, instan)
	}

	return &instansi, nil
}

func (r *UserRepository) GetDataFunction(rangeDate *domain.DateRangeDTO) (*[]domain.Pemakaian, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf(`
		BEGIN
			:l_cursor := PKG_MANAGE.GETPRODUCTLIST('%s','%s');
		END;
	`, rangeDate.DateFrom, rangeDate.DateTo)

	stmt, err := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		return nil, err
	}

	var rows driver.Rows

	_, err = stmt.ExecContext(ctx, sql.Out{Dest: &rows})

	if err != nil {
		return nil, err
	}

	rw := make([]driver.Value, len(rows.Columns()))

	columns := make([]string, len(rw))

	for i, col := range rows.Columns() {
		columns[i] = col
	}

	fields := make([]interface{}, len(rows.Columns()))
	var myMap = make(map[string]interface{})

	var Pemakaian []domain.Pemakaian

	for {
		if err := rows.Next(rw); err != nil {
			if err == io.EOF {
				break
			}
			rows.Close()
			return nil, err
		}
		for j, col := range rw {
			fields[j] = col
		}

		for k, data := range fields {
			myMap[columns[k]] = data
		}
		jsonString, _ := json.Marshal(myMap)

		s := domain.Pemakaian{}
		json.Unmarshal(jsonString, &s)

		Pemakaian = append(Pemakaian, s)
	}

	return &Pemakaian, nil
}

func (r *UserRepository) SetDataProcedure(rangeDate *domain.DateRangeDTO) (bool, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	query := fmt.Sprintf(`
			BEGIN
				SETDATAPRODUCT('%s', '%s');
			END;`,
			rangeDate.DateFrom,
			rangeDate.DateTo)

	stmt, err := r.db.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx)

	if err != nil {
		return false, err
	}

	return true, nil
}