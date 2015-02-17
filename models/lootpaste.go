package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

// LootPaste stores all required information about the loot paste
type LootPaste struct {
	ID         int64
	Character  string
	RawPaste   string
	Comment    string
	TotalValue int
	TaxAmount  int
}

// NewLootPaste creates a new loot paste with the given values
func NewLootPaste(character string, rawPaste string, comment string) *LootPaste {
	lootPaste := &LootPaste{
		ID:         -1,
		Character:  character,
		RawPaste:   rawPaste,
		Comment:    comment,
		TotalValue: 0,
		TaxAmount:  0,
	}

	return lootPaste
}

func (lootPaste *LootPaste) FetchValue() error {
	lootID, err := lootPaste.PasteLoot()
	if err != nil {
		return err
	}

	err = lootPaste.RetrieveLootValue(lootID)
	if err != nil {
		return err
	}

	return nil
}

func (lootPaste *LootPaste) PasteLoot() (string, error) {
	data := url.Values{}
	data.Set("raw_paste", lootPaste.RawPaste)

	req, err := http.NewRequest("POST", "http://evepraisal.com/estimate", bytes.NewBufferString(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	reg := regexp.MustCompile("Result #([0-9]+)")
	resultID := reg.FindStringSubmatch(string(body))

	return resultID[1], nil
}

func (lootPaste *LootPaste) RetrieveLootValue(lootID string) error {
	resp, err := http.Get(fmt.Sprintf("http://evepraisal.com/e/%s.json", lootID))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var evePraisal EvePraisal

	err = json.NewDecoder(resp.Body).Decode(&evePraisal)
	if err != nil {
		return err
	}

	lootPaste.TotalValue = int(evePraisal.GetTotalBuyValue() * 100)

	return nil
}
