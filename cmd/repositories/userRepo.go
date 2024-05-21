package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"
)

func CreateUser(user models.User) (models.User, error){
  db := storage.GetDB()

  sql := `INSERT INTO users(name, email, password) 
  VALUES($1, $2, $3) RETURNING id`
  err := db.QueryRow(sql, user.Name, user.Email, user.Password).Scan(&user.Id)
  if err != nil{
	return user, err
  }
  return user, nil
}

func UpdateUser(user models.User, id int)(models.User, error){
   db :=storage.GetDB()
   sql := `UPDATE users
           SET name= $2, email=$3, password=$4 , updated_at=$5
		   WHERE id = $1
		   RETURNING id`
    err := db.QueryRow(sql, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)	   
    if err!=nil{
		return models.User{}, err
	}
	user.Id =id
	return user, nil
}