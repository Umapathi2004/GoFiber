package Routes

import (
	"Fleet_GoFiber/database"
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

type Vehicle struct {
	Vehicle_ID        string  `json:"vehicle_id"`
	Owner_ID          string  `json:"owner_id"`
	Model_ID          string  `json:"model_id"`
	Registration_No   string  `json:"registration_no"`
	Registration_Date string  `json:"registration_date"`
	Milage            float64 `json:"milage"`
	Total_VKT         float64 `json:"total_vkt"`
	Purpose           string  `json:"purpose"`
	Manufacturing_No  string  `json:"manufacturing_no"`
	Vehicle_Status    string  `json:"vehicle_status"`
}

func VehicleRoute(vehicle fiber.Router) {
	vehicle.Get("/", GetVehicle)
	vehicle.Post("/", PostVehicle)
}

func GetVehicle(c *fiber.Ctx) error {
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
		query_temp := fmt.Sprintf("SELECT * FROM vehicle WHERE %v", whereClause)
		re := regexp.MustCompile(`'(\d+(\.\d+)?)'`)
		query = re.ReplaceAllString(query_temp, `$1`)
	} else {
		query = "SELECT * FROM vehicle"
	}

	log.Println(query)
	var vehicles []Vehicle
	rows, err := Sqldb.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(
			&vehicle.Vehicle_ID,
			&vehicle.Owner_ID,
			&vehicle.Model_ID,
			&vehicle.Registration_No,
			&vehicle.Registration_Date,
			&vehicle.Milage,
			&vehicle.Total_VKT,
			&vehicle.Purpose,
			&vehicle.Manufacturing_No,
			&vehicle.Vehicle_Status,
		); err != nil {
			return c.Status(500).SendString(err.Error())
		}
		vehicles = append(vehicles, vehicle)
	}
	return c.JSON(vehicles)
}

func PostVehicle(c *fiber.Ctx) error {
	Sqldb, dberr = database.Database_connect()
	if dberr != nil {
		log.Println("DB Connection Error: ", dberr)
		return c.Status(500).JSON("Database connection error")
	}
	defer Sqldb.Close()
	var vehicle Vehicle
	if err := c.BodyParser(&vehicle); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	query := `
		INSERT INTO vehicle (vehicle_id, owner_id, model_id, registration_no, registration_date, milage, total_vkt, purpose, manufacturing_no, vehicle_status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?,?,?)
	`
	_, err := Sqldb.Exec(query, vehicle.Vehicle_ID, vehicle.Owner_ID, vehicle.Model_ID, vehicle.Registration_No, vehicle.Registration_Date, vehicle.Milage, vehicle.Total_VKT, vehicle.Purpose, vehicle.Manufacturing_No, vehicle.Vehicle_Status)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(vehicle)
}
