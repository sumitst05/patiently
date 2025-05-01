package service

import (
	"errors"

	"github.com/sumitst05/patiently/internal/models"
	"github.com/sumitst05/patiently/internal/repository"
)

func CreatePatient(userId uint, role, name string, age int, gender, address, phone string) (*models.Patient, error) {
	if role != models.RoleReceptionist {
		return nil, errors.New("Only receptionists can create patients")
	}

	if name == "" || age <= 0 || phone == "" {
		return nil, errors.New("Invalid patient data")
	}

	patient := &models.Patient{
		Name:        name,
		Age:         age,
		Gender:      gender,
		Address:     address,
		Phone:       phone,
		CreatedByID: userId,
	}

	return repository.CreatePatient(patient)
}

func GetAllPatients(role string) ([]models.Patient, error) {
	if role != models.RoleReceptionist && role != models.RoleDoctor {
		return nil, errors.New("Invalid role")
	}

	return repository.GetAllPatients()
}

func GetPatientById(patientId uint, role string) (*models.Patient, error) {
	if role != models.RoleDoctor && role != models.RoleReceptionist {
		return nil, errors.New("Invalid role")
	}

	return repository.GetPatientById(patientId)
}

func GetPatientRegistrationHistory(patientId uint, role string) ([]models.RegistrationHistory, error) {
	if role != models.RoleReceptionist {
		return nil, errors.New("Only receptionists can view patient registration history")
	}

	return repository.GetPatientRegistrationHistory(patientId)
}

func UpdatePatient(patientId uint, role string, name string, age int, gender, address, phone string) (*models.Patient, error) {
	if role != models.RoleDoctor && role != models.RoleReceptionist {
		return nil, errors.New("Invalid role")
	}

	if age < 0 {
		return nil, errors.New("Invalid patient age")
	}

	patient, err := repository.GetPatientById(patientId)
	if err != nil {
		return nil, err
	}

	if name != "" {
		patient.Name = name
	}
	if age != 0 {
		patient.Age = age
	}
	if gender != "" {
		patient.Gender = gender
	}
	if address != "" {
		patient.Address = address
	}
	if phone != "" {
		patient.Phone = phone
	}

	return repository.UpdatePatient(patientId, patient)
}

func DeletePatient(patientId uint, role string) error {
	if role != models.RoleReceptionist {
		return errors.New("Only receptionists can create patients")
	}

	return repository.DeletePatient(patientId)
}
