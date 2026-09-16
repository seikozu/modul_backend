package service

import (
	"testing"

	"modul5/app/model"
)

func TestValidateCreate(t *testing.T) {
	req := model.CreateStudentRequest{
		NIM:   "",
		Name:  "Test Student",
		Grade: 105,
	}

	errs := ValidateCreate(req)

	if _, exists := errs["nim"]; !exists {
		t.Error("nim kosong seharusnya menghasilkan error")
	}
	if _, exists := errs["grade"]; !exists {
		t.Error("grade > 100 seharusnya menghasilkan error")
	}
}

func TestValidateReplace(t *testing.T) {
	req := model.ReplaceStudentRequest{
		NIM:   "342414001",
		Name:  "",
		Grade: -5,
	}

	errs := ValidateReplace(req)

	if _, exists := errs["name"]; !exists {
		t.Error("name kosong pada PUT seharusnya menghasilkan error")
	}
	if _, exists := errs["grade"]; !exists {
		t.Error("grade negatif pada PUT seharusnya menghasilkan error")
	}
}

func TestApplyPatch(t *testing.T) {
	initial := model.Student{
		ID:       1,
		NIM:      "342414001",
		Name:     "Yae miko",
		Grade:    95,
		IsActive: true,
	}

	newGrade := 98.5
	result, errs := ApplyPatch(initial, model.PatchStudentRequest{Grade: &newGrade})

	if len(errs) != 0 {
		t.Fatalf("tidak seharusnya ada error: %v", errs)
	}
	if result.Grade != 98.5 {
		t.Error("grade seharusnya berubah menjadi 98.5")
	}
	if result.Name != "Yae miko" {
		t.Error("field name yang tidak dikirim seharusnya tidak berubah")
	}
}