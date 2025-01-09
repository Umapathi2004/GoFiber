package Routes

import (
	"Fleet_GoFiber/database"
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Driver struct {
	DriverID          string  `json:"driver_id"`
	DriverName        string  `json:"driver_name"`
	Experience        float64 `json:"experience"`
	DOB               string  `json:"dob"`
	LicenceNO         string  `json:"licence_no"`
	LicenceExpireDate string  `json:"licence_expire_date"`
	DriverStatus      string  `json:"driver_status"`
	DriverRating      float64 `json:"driver_rating"`
}

var Sqldb *sql.DB
var dberr error

func DriverRoute(driver fiber.Router) {
	driver.Get("/", GetDriver)
	driver.Post("/", PostDriver)
}

func GetDriver(c *fiber.Ctx) error {
	Sqldb, dberr = database.Database_connect()
	if dberr != nil {
		log.Println("DB Connection Error: ", dberr)
		return c.Status(500).SendString("Database connection error")
	}
	defer Sqldb.Close()

	var conditions []string
	Query := c.Queries()
	for key, value := range Query {
		conditions = append(conditions, fmt.Sprintf("%s = '%s'", key, strings.ReplaceAll(value, "'", "")))

	}
	var query string
	whereClause := strings.Join(conditions, " AND ")

	if len(conditions) > 0 {
		query_temp := fmt.Sprintf("SELECT * FROM driver WHERE %v", whereClause)
		re := regexp.MustCompile(`'(\d+(\.\d+)?)'`)
		query = re.ReplaceAllString(query_temp, `$1`)
		log.Println(query)
	} else {
		query = "SELECT * FROM driver"
	}

	var drivers []Driver
	rows, err := Sqldb.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var driver Driver
		if err := rows.Scan(
			&driver.DriverID,
			&driver.DriverName,
			&driver.Experience,
			&driver.DOB,
			&driver.LicenceNO,
			&driver.LicenceExpireDate,
			&driver.DriverStatus,
			&driver.DriverRating,
		); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		drivers = append(drivers, driver)
	}
	return c.JSON(drivers)
}

func PostDriver(c *fiber.Ctx) error {
	Sqldb, dberr = database.Database_connect()
	if dberr != nil {
		log.Println("DB Connection Error: ", dberr)
		return c.Status(500).JSON("Database connection error")
	}
	defer Sqldb.Close()
	var driver Driver
	if err := c.BodyParser(&driver); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	query := `
		INSERT INTO driver (driver_id, driver_name, experience, dob, licence_no, licence_expire_date, driver_status, driver_rating)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := Sqldb.Exec(query, driver.DriverID, driver.DriverName, driver.Experience, driver.DOB, driver.LicenceNO, driver.LicenceExpireDate, driver.DriverStatus, driver.DriverRating)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(driver)
}
