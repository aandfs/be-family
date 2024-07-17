package familyListController

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
	var familyList []models.FamilyList
	var FamilyListResponse []models.FamilyListResponse

	if err := config.DB.Joins("JOIN customers ON family_lists.cst_id = customers.cst_id").
		Joins("JOIN nationalities ON customers.nationality_id = nationalities.nationality_id").
		Select("family_lists.*, customers.cst_name, customers.nationality_id, nationalities.nationality_name").
		Find(&familyList).Find(&FamilyListResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "List Family List", FamilyListResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var familyList models.FamilyList
	if err := json.NewDecoder(r.Body).Decode(&familyList); err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	//check costumer
	var costumer models.Customers
	if err := config.DB.First(&costumer, familyList.CstId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "costumer not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&familyList).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create family list", nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var familyList models.FamilyList

	if err := config.DB.First(&familyList, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "family list not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var familyListPayload models.FamilyList
	if err := json.NewDecoder(r.Body).Decode(&familyListPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var costumer models.Customers
	if familyListPayload.CstId != 0 {
		if err := config.DB.First(&costumer, familyListPayload.CstId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Costumer not found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("fl_id = ?", id).Updates(&familyListPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success update family list", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var familyList models.FamilyList
	var familyListResponse models.FamilyListResponse

	if err := config.DB.Joins("Customers").First(&familyList, id).First(&familyListResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "familyList not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail familyList", familyListResponse)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var familyList models.FamilyList
	res := config.DB.Delete(&familyList, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "family list not found", nil)
		return
	}

	helper.Response(w, 200, "sucess delete family list", nil)
}
