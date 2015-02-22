package mysql

import (
	"fmt"

	"github.com/morpheusxaut/evetax/misc"
	"github.com/morpheusxaut/evetax/models"

	// Blank import of the MySQL driver to use with sqlx
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DatabaseConnection provides an implementation of the Connection interface using a MySQL database
type DatabaseConnection struct {
	// Config stores the current configuration values being used
	Config *misc.Configuration

	conn *sqlx.DB
}

// Connect tries to establish a connection to the MySQL backend, returning an error if the attempt failed
func (c *DatabaseConnection) Connect() error {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", c.Config.DatabaseUser, c.Config.DatabasePassword, c.Config.DatabaseHost, c.Config.DatabaseSchema))
	if err != nil {
		return err
	}

	c.conn = conn

	return nil
}

// RawQuery performs a raw MySQL query and returns a map of interfaces containing the retrieve data. An error is returned if the query failed
func (c *DatabaseConnection) RawQuery(query string, v ...interface{}) ([]map[string]interface{}, error) {
	rows, err := c.conn.Query(query, v...)
	if err != nil {
		return nil, err
	}

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	var results []map[string]interface{}

	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		rows.Scan(valuePtrs...)

		resultRow := make(map[string]interface{})

		for i, col := range columns {
			resultRow[col] = values[i]
		}

		results = append(results, resultRow)
	}

	return results, nil
}

// LoadAllLootPastes retrieves all loot pastes from the MySQL database, returning an error if the query failed
func (c *DatabaseConnection) LoadAllLootPastes() ([]*models.LootPaste, error) {
	var lootpastes []*models.LootPaste

	err := c.conn.Select(&lootpastes, "SELECT id, charactername, rawpaste, pastecomment, totalvalue, taxamount, timestamp FROM lootpastes")
	if err != nil {
		return nil, err
	}

	return lootpastes, nil
}

// SaveLootPaste saves a loot paste to the MySQL database, returning the updated model or an error if the query failed
func (c *DatabaseConnection) SaveLootPaste(lootPaste *models.LootPaste) (*models.LootPaste, error) {
	if lootPaste.ID > 0 {
		_, err := c.conn.Exec("UPDATE lootpastes SET charactername=?, rawpaste=?, pastecomment=?, totalvalue=?, taxamount=?, timestamp=? WHERE id=?", lootPaste.CharacterName, lootPaste.RawPaste, lootPaste.PasteComment, lootPaste.TotalValue, lootPaste.TaxAmount, lootPaste.Timestamp, lootPaste.ID)
		if err != nil {
			return nil, err
		}
	} else {
		resp, err := c.conn.Exec("INSERT INTO lootpastes(charactername, rawpaste, pastecomment, totalvalue, taxamount, timestamp) VALUES(?, ?, ?, ?, ?, ?)", lootPaste.CharacterName, lootPaste.RawPaste, lootPaste.PasteComment, lootPaste.TotalValue, lootPaste.TaxAmount, lootPaste.Timestamp)
		if err != nil {
			return nil, err
		}

		lastInsertedID, err := resp.LastInsertId()
		if err != nil {
			return nil, err
		}

		lootPaste.ID = lastInsertedID
	}

	return lootPaste, nil
}
