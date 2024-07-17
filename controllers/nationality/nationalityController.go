package nationalityController

import (
	"be-family/config"
	"be-family/helper"
	"be-family/models"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var nationality []models.Nationality

	if err := config.DB.Find(&nationality).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Nationalities", nationality)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var nationality models.Nationality

	if err := json.NewDecoder(r.Body).Decode(&nationality); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := nationality.Validate(); err != nil {
		helper.Response(w, 400, err.Error(), nil)
		return
	}

	var existingNationality models.Nationality
	if err := config.DB.Where("nationality_name = ?", nationality.NationalityName).First(&existingNationality).Error; err == nil {
		helper.Response(w, 400, "Nationality name already exists", nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&nationality).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create nationality", nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var nationality models.Nationality

	if err := config.DB.First(&nationality, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "nationality not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&nationality); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("nationality_id = ?", id).Updates(&nationality).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success update nationality", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var nationality models.Nationality
	res := config.DB.Delete(&nationality, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "nationality not found", nil)
		return
	}

	helper.Response(w, 200, "sucess delete nationality", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var nationality models.Nationality

	if err := config.DB.First(&nationality, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "nationality not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail nationality", nationality)
}
