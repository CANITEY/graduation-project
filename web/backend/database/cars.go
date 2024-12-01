package database

import (
	"database/sql"
	"gp-backend/models"
)

func AddCar(d *sql.DB, carInfo models.CarInfo) (error) {
	stmt := `INSERT INTO cars(uuid, longitude, latitude, driver_status, served) VALUES ($1, $2, $3, $4, false)`
	_, err := d.Exec(stmt, carInfo.UUID, carInfo.Longitude, carInfo.Latitude, carInfo.DriverStatus)
	if err != nil {
		return err
	}

	return nil
}
