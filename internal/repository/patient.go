package repository

import (
	"time"

	"github.com/sumitst05/patiently/internal/models"
	"github.com/sumitst05/patiently/internal/utils"
)

func CreatePatient(patient *models.Patient) (*models.Patient, error) {
	err := DB.Create(patient).Error
	if err != nil {
		return nil, err
	}

	err = DB.Preload("CreatedBy").First(&patient, patient.ID).Error
	if err != nil {
		return nil, err
	}

	// json string for new snapshot of patient registration
	newVal, err := utils.SnapshotStruct(patient)

	logHistory := &models.RegistrationHistory{
		PatientID:   patient.ID,
		Action:      "create",
		ChangedByID: patient.CreatedByID,
		Timestamp:   time.Now(),
		NewValue:    newVal,
	}

	err = DB.Create(logHistory).Error
	if err != nil {
		return nil, err
	}

	err = DB.Preload("ChangedBy").First(&logHistory, logHistory.ID).Error
	if err != nil {
		return nil, err
	}

	return patient, nil
}

func GetAllPatients() ([]models.Patient, error) {
	patients := []models.Patient{}
	err := DB.Preload("CreatedBy").Find(&patients).Error
	return patients, err
}

func GetPatientById(id uint) (*models.Patient, error) {
	patient := models.Patient{}
	err := DB.Preload("CreatedBy").First(&patient, id).Error
	return &patient, err
}

func GetPatientRegistrationHistory(patientId uint) ([]models.RegistrationHistory, error) {
	history := []models.RegistrationHistory{}

	err := DB.Preload("Patient").Preload("ChangedBy").Where("patient_id = ?", patientId).Find(&history).Error
	if err != nil {
		return nil, err
	}

	return history, nil
}

func UpdatePatient(id uint, updatedPatient *models.Patient) (*models.Patient, error) {
	patient := models.Patient{}
	err := DB.First(&patient, id).Error
	if err != nil {
		return nil, err
	}

	err = DB.Preload("CreatedBy").First(&patient, patient.ID).Error
	if err != nil {
		return nil, err
	}

	// json string for old value for snapshot of patient registration
	oldVal, err := utils.SnapshotStruct(updatedPatient)

	// selectively updating the fields received in payload
	err = utils.UpdateStruct(&patient, updatedPatient)
	if err != nil {
		return nil, err
	}

	patient.UpdatedAt = time.Now()

	err = DB.Save(&patient).Error
	if err != nil {
		return nil, err
	}

	// json string for new value for snapshot of patient registration
	newVal, err := utils.SnapshotStruct(patient)

	logHistory := &models.RegistrationHistory{
		PatientID:   patient.ID,
		Action:      "update",
		ChangedByID: patient.CreatedByID,
		Timestamp:   time.Now(),
		OldValue:    oldVal,
		NewValue:    newVal,
	}

	err = DB.Create(logHistory).Error
	if err != nil {
		return nil, err
	}

	err = DB.Preload("ChangedBy").First(&logHistory, logHistory.ID).Error
	if err != nil {
		return nil, err
	}

	return &patient, err
}

func DeletePatient(id uint) error {
	patient := models.Patient{}

	err := DB.First(&patient, id).Error
	if err != nil {
		return err
	}

	err = DB.Delete(&patient).Error
	return err
}
