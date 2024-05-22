package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	passwordHashed, _ := hashPassword(user.Password)

	sql := `INSERT INTO users(name, email, password) 
  VALUES($1, $2, $3) RETURNING id`
	err := db.QueryRow(sql, user.Name, user.Email, passwordHashed).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	db := storage.GetDB()
	sql := `UPDATE users
           SET name= $2, email=$3, password=$4 , updated_at=$5
		   WHERE id = $1
		   RETURNING id`
	err := db.QueryRow(sql, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
	if err != nil {
		return models.User{}, err
	}
	user.Id = id
	return user, nil
}

func GetUser(id int) (models.User, error) {
	db := storage.GetDB()
	sql := `SELECT  id, name, email, created_at, updated_at  FROM users WHERE id=$1`
	var user models.User
	err := db.QueryRow(sql, id).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAT)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetAll() ([]models.User, error) {
	db := storage.GetDB()
	sql := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := db.Query(sql)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAT); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func AuthenticateUser(username,password string) (models.User, bool){
	db := storage.GetDB()
	sql := `SELECT  id, name, email, password  FROM users WHERE email=$1`
	var user models.User
	err := db.QueryRow(sql, username).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, false
	}
    isPasswordCorrect := checkPassword(user.Password, password)

	if isPasswordCorrect  {
        return user, true
	}
    return models.User{}, false
}

func hashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string)bool{
  err:= bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
  return err==nil
}