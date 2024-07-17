package customerController

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
	var customers []models.Customers

	if err := config.DB.Preload("Nationality").
		Joins("JOIN nationalities ON customers.nationality_id = nationalities.nationality_id").
		Select("customers.*, nationalities.nationality_name").
		Find(&customers).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var response []models.CustomersResponse
	for _, customer := range customers {
		cstDOB := customer.CstDOB.Format("02-01-2006")
		response = append(response, models.CustomersResponse{
			CstId:           customer.CstId,
			NationalityId:   customer.NationalityId,
			NationalityName: customer.Nationality.NationalityName,
			CstName:         customer.CstName,
			CstPhoneNum:     customer.CstPhoneNum,
			CstDOB:          cstDOB,
			CstEmail:        customer.CstEmail,
		})
	}

	helper.Response(w, 200, "List Customers", response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var customer models.Customers
	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		helper.Response(w, 500, err.Error(), nil)
	}

	defer r.Body.Close()

	//check nationality
	var nationality models.Nationality
	if err := config.DB.First(&nationality, customer.NationalityId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "nationality not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&customer).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success create customer", nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var customer models.Customers

	if err := config.DB.First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "customer not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	var customerPayload models.Customers
	if err := json.NewDecoder(r.Body).Decode(&customerPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	var nationality models.Nationality
	if customerPayload.NationalityId != 0 {
		if err := config.DB.First(&nationality, customerPayload.NationalityId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Nationality not found", nil)
				return
			}

			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Where("cst_id = ?", id).Updates(&customerPayload).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "success update customer", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var customer models.Customers

	if err := config.DB.Preload("Nationality").First(&customer, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Customer not found", nil)
			return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	cstDOB := customer.CstDOB.Format("02-01-2006")
	response := map[string]interface{}{
		"cst_id":         customer.CstId,
		"nationality_id": customer.NationalityId,
		"cst_name":       customer.CstName,
		"cst_dob":        cstDOB,
		"cst_phone_num":  customer.CstPhoneNum,
		"cst_email":      customer.CstEmail,
		"nationality": map[string]interface{}{
			"nationality_id":   customer.Nationality.NationalityId,
			"nationality_name": customer.Nationality.NationalityName,
			"nationality_code": customer.Nationality.NationalityCode,
		},
	}

	helper.Response(w, 200, "Detail customer", response)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var costumer models.Customers
	res := config.DB.Delete(&costumer, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "costumer not found", nil)
		return
	}

	helper.Response(w, 200, "sucess delete costumer", nil)
}
