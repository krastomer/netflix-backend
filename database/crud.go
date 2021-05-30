package database

import (
	"sync"

	"github.com/krastomer/netflix-backend/models"
	"gorm.io/gorm"
)

var wg sync.WaitGroup

func AddUser(u models.User) error {
	d := GetDB()
	result := d.Create(&u)
	return result.Error
}

func GetUser(u string) models.User {
	d := GetDB()
	user := models.User{}
	d.First(&user, "email=?", u)
	return user
}

func GetUserPayment(u string) models.UserPayment {
	d := GetDB()
	payment := models.UserPayment{}
	d.First(&payment, "email=?", u)
	return payment
}

func SetUserPayment(payment models.UserPayment) error {
	d := GetDB()
	result := d.Where("email=?", payment.Email).Save(&payment)
	return result.Error
}

func GetUserProfile(u string) models.UserProfile {
	d := GetDB()
	user := models.UserProfile{}
	d.First(&user, "email=?", u)
	return user
}

func GetMovieDetail(id int) models.MovieDetail {
	d := GetDB()
	m := models.MovieDetail{}
	wg.Add(4)
	go getActorList(d, id, &m)
	go getDirectorList(d, id, &m)
	go getGenresList(d, id, &m)
	go getDetailList(d, id, &m)
	wg.Wait()
	return m
}

func getActorList(d *gorm.DB, m int, md *models.MovieDetail) {
	defer wg.Done()
	castRows, _ := d.Raw("SELECT DISTINCT cast.name FROM `movie_and_series` JOIN `casting` ON `casting`.`id_movie` = `movie_and_series`.`id_movie` JOIN `cast` ON `cast`.`id_cast` = `casting`.`id_cast` WHERE movie_and_series.id_movie = ?", m).Rows()
	defer castRows.Close()
	for castRows.Next() {
		var castName string
		castRows.Scan(&castName)
		md.AddActor(castName)
	}
}

func getDirectorList(d *gorm.DB, m int, md *models.MovieDetail) {
	defer wg.Done()
	directorRows, _ := d.Raw("SELECT DISTINCT director.name FROM `movie_and_series` JOIN `director_movie` ON `director_movie`.`id_movie` = `movie_and_series`.`id_movie` JOIN `director` ON `director_movie`.`id_director` = `director`.`id_director` WHERE movie_and_series.id_movie = ?", m).Rows()
	for directorRows.Next() {
		var directorName string
		directorRows.Scan(&directorName)
		md.AddDirector(directorName)
	}
}

func getGenresList(d *gorm.DB, m int, md *models.MovieDetail) {
	defer wg.Done()
	genresRows, _ := d.Raw("SELECT DISTINCT genres.name FROM `movie_and_series` JOIN `genres_movie` ON `genres_movie`.`id_movie` = `movie_and_series`.`id_movie` JOIN `genres` ON `genres`.`id_genres` = `genres_movie`.`id_genres` WHERE movie_and_series.id_movie = ?", m).Rows()
	for genresRows.Next() {
		var genresName string
		genresRows.Scan(&genresName)
		md.AddGenres(genresName)
	}
}

func getDetailList(d *gorm.DB, s int, m *models.MovieDetail) {
	defer wg.Done()
	row := d.Raw("SELECT movie_and_series.name,movie_and_series.is_series, season.year, movie_and_series.rate, season.description, COUNT(movie_and_series.id_movie) FROM `movie_and_series` JOIN `season` on `movie_and_series`.`id_movie` = `season`.`id_movie` WHERE movie_and_series.id_movie =?", s).Row()
	row.Scan(&m.Name, &m.IsSeries, &m.Year, &m.Rate, &m.Description, &m.NumberSeason)
}
