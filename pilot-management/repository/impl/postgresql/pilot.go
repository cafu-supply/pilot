package postgresql

import (
	"../../../domain/entity"
	"upper.io/db.v3/lib/sqlbuilder"
)

type PilotRepo struct {
	readConn  sqlbuilder.Database
	writeConn sqlbuilder.Database
}

type Pilot struct {
	Id          string `db:"id"`
	UserId      string `db:"user_id"`
	CodeName    string `db:"code_name"`
	SupplierId  string `db:"supplier_id"`
	MarketId    string `db:"market_id"`
	ServiceId   string `db:"service_id"`
	Status      string `db:"status"`
}

func MakePostgresPilotRepo() PilotRepo {
	return PilotRepo{
		readConn:  getReadConn(),
		writeConn: getWriteConn(),
	}
}

func (repo *PilotRepo) ListPilots() ([]entity.Pilot, error) {
	resultSet := make([]Pilot, 0)
	err := repo.readConn.Collection("pilots").Find().All(&resultSet)
	if err != nil {
		return nil, err
	}
	pilots := make([]entity.Pilot, 0)
	for _, pilot := range resultSet {
		pilots = append(pilots, entity.Pilot(pilot))
	}
	return pilots, nil
}

func (repo *PilotRepo) UpdatePilotStatus(pilot *entity.Pilot) (*entity.Pilot, error) {

	res := repo.readConn.Collection("pilots").Find(pilot.Id)
	err := res.One(&pilot)
	if err != nil {
		return nil, err
	}

	q := repo.writeConn.Update("pilots").
		Set("status = ?, ", pilot.Status).
		Where("id = ?", pilot.Id)

	_ , err = q.Exec()

	if err != nil {
		return nil, err
	}

	return pilot, nil
}
