package database

import (
	"errors"
	"sync"
	"time"

	"github.com/krastomer/netflix-backend/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var wg sync.WaitGroup
var (
	errMaxViewer      = errors.New("maximum size viewer")
	errNotFoundViewer = errors.New("not found viewer")
)

const (
	emptyViewer   = 0
	maximumViewer = 5
)

func GetMaxViewerError() error {
	return errMaxViewer
}

func GetNotFoundViewerError() error {
	return errNotFoundViewer
}

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

// TODO: NextBilling invalid data type
func GetUserPayment(u string) models.UserPayment {
	d := GetDB()
	payment := models.UserPayment{}
	d.First(&payment, "email=?", u)
	return payment
}

func SetUserPayment(payment models.UserPayment) error {
	d := GetDB()
	result := d.Model(&payment).Where("email=?", payment.Email).Updates(
		map[string]interface{}{
			"firstname":     payment.Firstname,
			"lastname":      payment.Lastname,
			"phone_number":  payment.PhoneNumber,
			"card_number":   payment.CardNumber,
			"exp_date":      payment.ExpDate,
			"security_code": payment.SecurityCode,
		})
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
	castRows, _ := d.Raw("SELECT DISTINCT cast.id_cast, cast.name FROM `movie_and_series` JOIN `casting` ON `casting`.`id_movie` = `movie_and_series`.`id_movie` JOIN `cast` ON `cast`.`id_cast` = `casting`.`id_cast` WHERE movie_and_series.id_movie = ?", m).Rows()
	defer castRows.Close()
	for castRows.Next() {
		var id int
		var name string
		castRows.Scan(&id, &name)
		md.AddActor(models.People{ID: id, Name: name})
	}
}

func getDirectorList(d *gorm.DB, m int, md *models.MovieDetail) {
	defer wg.Done()
	directorRows, _ := d.Raw("SELECT DISTINCT director.id_director,director.name FROM `movie_and_series` JOIN `director_movie` ON `director_movie`.`id_movie` = `movie_and_series`.`id_movie` JOIN `director` ON `director_movie`.`id_director` = `director`.`id_director` WHERE movie_and_series.id_movie = ?", m).Rows()
	for directorRows.Next() {
		var id int
		var name string
		directorRows.Scan(&id, &name)
		md.AddDirector(models.People{ID: id, Name: name})
	}
}

func getGenresList(d *gorm.DB, m int, md *models.MovieDetail) {
	defer wg.Done()
	genresRows, _ := d.Raw("SELECT DISTINCT genres.id_genres, genres.name FROM `movie_and_series` JOIN `genres_movie` ON `genres_movie`.`id_movie` = `movie_and_series`.`id_movie` JOIN `genres` ON `genres`.`id_genres` = `genres_movie`.`id_genres` WHERE movie_and_series.id_movie = ?", m).Rows()
	for genresRows.Next() {
		var id int
		var name string
		genresRows.Scan(&id, &name)
		md.AddGenres(models.People{ID: id, Name: name})
	}
}

func getDetailList(d *gorm.DB, s int, m *models.MovieDetail) {
	defer wg.Done()
	row := d.Raw("SELECT movie_and_series.id_movie, movie_and_series.name,movie_and_series.is_series, season.year, movie_and_series.rate, season.description, COUNT(movie_and_series.id_movie) FROM `movie_and_series` JOIN `season` on `movie_and_series`.`id_movie` = `season`.`id_movie` WHERE movie_and_series.id_movie =?", s).Row()
	row.Scan(&m.ID, &m.Name, &m.IsSeries, &m.Year, &m.Rate, &m.Description, &m.NumberSeason)
}

func GetListMovieFromActor(id int) models.MovieList {
	ml := models.MovieList{ID: id}
	wg.Add(2)
	go func(d *gorm.DB, s int) {
		defer wg.Done()
		listRows, _ := d.Raw("SELECT movie_and_series.id_movie, movie_and_series.name from cast join casting on cast.id_cast = casting.id_cast JOIN movie_and_series on movie_and_series.id_movie = casting.id_movie WHERE cast.id_cast = ?", id).Rows()
		for listRows.Next() {
			var id int
			var name string
			listRows.Scan(&id, &name)
			ml.AddMovie(models.People{ID: id, Name: name})
		}
	}(GetDB(), id)
	go func(d *gorm.DB, s int) {
		defer wg.Done()
		row := d.Raw("SELECT cast.name FROM `cast` WHERE cast.id_cast = ?", s).Row()
		row.Scan(&ml.Name)
	}(GetDB(), id)
	wg.Wait()
	return ml
}

func ReBillingPayment(payment *models.UserPayment, u string) error {
	d := GetDB()
	result := d.Model(&payment).Where("email", u).Update("next_billing", time.Now().AddDate(0, 1, 0).Format("2006-01-02"))
	return result.Error
}

func SetReceiptPayment(u models.UserProfile) error {
	d := GetDB()
	bill := models.UserBilling{BillingDate: time.Now().Local(), IDAccount: u.IDAccount}
	result := d.Create(&bill)
	return result.Error
}

func CancelMemberShip(u models.UserPayment) error {
	d := GetDB()
	result := d.Model(&u).Where("email=?", u.Email).Updates(
		map[string]interface{}{
			"firstname":     gorm.Expr("NULL"),
			"lastname":      gorm.Expr("NULL"),
			"phone_number":  gorm.Expr("NULL"),
			"card_number":   gorm.Expr("NULL"),
			"exp_date":      gorm.Expr("NULL"),
			"security_code": gorm.Expr("NULL"),
		})
	return result.Error
}

func getSizeViewer(u string) int {
	d := GetDB()
	var size int
	row := d.Raw("SELECT COUNT(id_viewer) FROM `viewer` WHERE id_account = (SELECT id_account FROM user WHERE email = ?) ", u).Row()
	row.Scan(&size)
	return size
}

func GetListViewer(u string) []models.Viewer {
	d := GetDB()
	if getSizeViewer(u) == emptyViewer {
		CreateViewer(u, models.Viewer{Name: "You"})
	}
	listViewer := []models.Viewer{}
	viewersRows, _ := d.Raw("SELECT viewer.id_viewer,viewer.id_account,  viewer.name ,viewer.pin_number,viewer.is_kid FROM viewer JOIN user ON viewer.id_account = user.id_account WHERE user.email = ?", u).Rows()
	for viewersRows.Next() {
		viewer := models.Viewer{}
		viewersRows.Scan(&viewer.IDViewer, &viewer.IDAccount, &viewer.Name, &viewer.PinNumber, &viewer.IsKid)
		listViewer = append(listViewer, viewer)
	}
	return listViewer
}

func CreateViewer(u string, v models.Viewer) error {
	d := GetDB()
	if getSizeViewer(u) == maximumViewer {
		return errMaxViewer
	}
	var pinNumber clause.Expr
	if v.PinNumber == "" {
		pinNumber = gorm.Expr("Null")
	} else {
		pinNumber = gorm.Expr(v.PinNumber)
	}
	result := d.Exec("INSERT INTO viewer(id_account, pin_number, name, is_kid) VALUES((SELECT id_account FROM user WHERE email=?), ?, ?, ?)", u, pinNumber, v.Name, v.IsKid)
	return result.Error
}

func DeleteViewer(b models.BodyViewer) error {
	d := GetDB()
	viewers := GetListViewer(b.Email)
	for _, viewer := range viewers {
		if viewer.IDViewer == b.IDViewer {
			result := d.Where("id_viewer=?", b.IDViewer).Delete(&models.BodyViewer{})
			return result.Error
		}
	}
	return errNotFoundViewer
}
