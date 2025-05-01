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

func UpdatePatient(id uint, patient *models.Patient) (*models.Patient, error) {
	existing := models.Patient{}
	if err := DB.First(&existing, id).Error; err != nil {
		return nil, err
	}

	// json string for old snapshot of patient registration
	oldVal, err := utils.SnapshotStruct(&existing)
	if err != nil {
		return nil, err
	}

	patient.UpdatedAt = time.Now()
	if err := DB.Save(patient).Error; err != nil {
		return nil, err
	}

	// json string for new snapshot of patient registration
	newVal, err := utils.SnapshotStruct(patient)
	if err != nil {
		return nil, err
	}

	log := &models.RegistrationHistory{
		PatientID:   patient.ID,
		Action:      "update",
		ChangedByID: patient.CreatedByID,
		Timestamp:   time.Now(),
		OldValue:    oldVal,
		NewValue:    newVal,
	}

	if err := DB.Create(log).Error; err != nil {
		return nil, err
	}

	if err := DB.Preload("ChangedBy").First(log, log.ID).Error; err != nil {
		return nil, err
	}

	return patient, nil
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
