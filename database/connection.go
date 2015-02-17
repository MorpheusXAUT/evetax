package database

import (
	"fmt"

	"github.com/morpheusxaut/evetax/database/mysql"
	"github.com/morpheusxaut/evetax/misc"
	"github.com/morpheusxaut/evetax/models"
)

// Connection provides an interface for communicating with a database backend in order to retrieve and persist the needed information
type Connection interface {
	// Connect tries to establish a connection to the database backend, returning an error if the attempt failed
	Connect() error

	// RawQuery performs a raw database query and returns a map of interfaces containing the retrieve data. An error is returned if the query failed
	RawQuery(query string, v ...interface{}) ([]map[string]interface{}, error)

	// SaveLootPaste saves a loot paste to the database, returning the updated model or an error if the query failed
	SaveLootPaste(lootPaste *models.LootPaste) (*models.LootPaste, error)
}

// SetupDatabase parses the database type set in the configuration and returns an appropriate database implementation or an error if the type is unknown
func SetupDatabase(conf *misc.Configuration) (Connection, error) {
	var database Connection

	switch Type(conf.DatabaseType) {
	case TypeMySQL:
		database = &mysql.DatabaseConnection{
			Config: conf,
		}
		break
	default:
		return nil, fmt.Errorf("Unknown type #%d", conf.DatabaseType)
	}

	return database, nil
}
