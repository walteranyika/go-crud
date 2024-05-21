package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateMeasurement(measurement models.Measurement) (models.Measurement, error){
    db := storage.GetDB()
	sql := `INSERT INTO measurements (user_id, weight, height, body_fat, created_at)
	VALUES($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(sql, measurement.UserId, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&measurement.Id)
	if err !=nil{
		return measurement, err
	}
	return measurement, nil
}

func UpdateMeasurement(measurement models.Measurement, id int)(models.Measurement, error){
	db:=storage.GetDB()
	sql :=`UPDATE measurements SET
	       weight=$2, height=$3, body_fat=$4, created_at=$5
     	   WHERE id=$1 RETURNING id`
	err := db.QueryRow(sql, id, measurement.Weight, measurement.Height, measurement.BodyFat, time.Now()).Scan(&id)
	if  err !=nil {
		return models.Measurement{}, err
	}
	measurement.Id = id
	return	measurement, nil 
}