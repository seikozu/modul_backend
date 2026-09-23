package service

import (
	"testing"
	"modul6/app/model"
)

func TestValidateRegister_WeakPassword(t *testing.T) {
	req := model.RegisterRequest{
		Username: "sari",
		Email:    "sari@example.com",
		Password: "password1",
	}

	errs := ValidateRegister(req)
	if msg, ok := errs["password"]; !ok || msg != "password terlalu umum" {
		t.Errorf("diharapkan error 'password terlalu umum', didapat: %v", errs)
	}
}

func TestValidateRegister_ShortPassword(t *testing.T) {
	req := model.RegisterRequest{
		Username: "sari",
		Email:    "sari@example.com",
		Password: "abc",
	}

	errs := ValidateRegister(req)
	if msg, ok := errs["password"]; !ok || msg != "minimal 8 karakter" {
		t.Errorf("diharapkan error 'minimal 8 karakter', didapat: %v", errs)
	}
}

func TestValidateRegister_NoDigit(t *testing.T) {
	req := model.RegisterRequest{
		Username: "sari",
		Email:    "sari@example.com",
		Password: "passwordhanya",
	}

	errs := ValidateRegister(req)
	if msg, ok := errs["password"]; !ok || msg != "harus memuat huruf dan angka" {
		t.Errorf("diharapkan error 'harus memuat huruf dan angka', didapat: %v", errs)
	}
}