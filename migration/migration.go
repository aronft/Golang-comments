package migration

import (
	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/configuration"
	"github.com/blazte/GoED/10ProyectoFInal/Golang-comments/models"
)

func Migrate() {
	db := configuration.GetConnection()
	defer db.Close()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.Comment{})
	db.CreateTable(&models.Vote{})
	db.Model(&models.Vote{}).AddUniqueIndex("comment_id_user_id_unique", "comment_id", "user_id")
}
