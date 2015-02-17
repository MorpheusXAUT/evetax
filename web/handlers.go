package web

import (
	"fmt"
	"net/http"

	"github.com/morpheusxaut/evetax/misc"
	"github.com/morpheusxaut/evetax/models"
)

// IndexGetHandler displays the index page of the web app
func (controller *Controller) IndexGetHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["pageTitle"] = "Tax Calculator"

	response["status"] = 0
	response["result"] = nil

	controller.SendResponse(w, r, "index", response)
}

// TaxesPostHandler handles loot pastes and displays the appropriate tax amount
func (controller *Controller) TaxesPostHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["pageTitle"] = "Taxes"

	response["status"] = 0
	response["result"] = nil

	err := r.ParseForm()
	if err != nil {
		misc.Logger.Warnf("Failed to parse form: [%v]", err)

		response["status"] = 1
		response["result"] = fmt.Errorf("Failed to parse form, please try again!")

		controller.SendResponse(w, r, "taxes", response)

		return
	}

	character := r.FormValue("character")
	rawPaste := r.FormValue("rawPaste")
	comment := r.FormValue("comment")

	if len(character) == 0 || len(rawPaste) == 0 {
		misc.Logger.Warnf("Received empty character or rawPaste")

		response["status"] = 1
		response["result"] = fmt.Errorf("Empty character or loot paste, please try again!")

		controller.SendResponse(w, r, "taxes", response)

		return
	}

	lootPaste := models.NewLootPaste(character, rawPaste, comment)

	err = lootPaste.FetchValue()
	if err != nil {
		misc.Logger.Warnf("Failed to fetch value of loot paste: [%v]", err)

		response["status"] = 1
		response["result"] = fmt.Errorf("Failed to retrieve loot paste values, please try again!")

		controller.SendResponse(w, r, "taxes", response)

		return
	}

	misc.Logger.Println(lootPaste.TotalValue)
	misc.Logger.Println(controller.Config.TaxPercentage)
	misc.Logger.Println((lootPaste.TotalValue * controller.Config.TaxPercentage) / 100.0)

	lootPaste.TaxAmount = int((lootPaste.TotalValue * controller.Config.TaxPercentage) / 100.0)

	misc.Logger.Println(lootPaste.TaxAmount)

	lootPaste, err = controller.Database.SaveLootPaste(lootPaste)
	if err != nil {
		misc.Logger.Warnf("Failed to save loot paste: [%v]", err)

		response["status"] = 1
		response["result"] = fmt.Errorf("Failed to save loot paste, please try again!")

		controller.SendResponse(w, r, "taxes", response)

		return
	}

	response["status"] = 0
	response["result"] = nil
	response["lootPaste"] = lootPaste

	controller.SendResponse(w, r, "taxes", response)
}

// LegalGetHandler displays some legal information as well as copyright disclaimers and contact info
func (controller *Controller) LegalGetHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["pageTitle"] = "Legal"

	response["status"] = 0
	response["result"] = nil

	controller.SendResponse(w, r, "legal", response)
}
