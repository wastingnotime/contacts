package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateContact(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(echo.POST, "/contacts", strings.NewReader(`{"firstName":"John","lastName":"Doe","phoneNumber":"1234567890"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	seed := map[string]contact{}
	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.CreateContact(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.NotEmpty(t, rec.Header().Get(echo.HeaderLocation))

		id := strings.Replace(rec.Header().Get(echo.HeaderLocation), "/contacts/", "", -1)

		var co contact
		for _, value := range seed {
			co = value
			break
		}
		assert.Equal(t, id, co.ID)
		assert.Equal(t, "John", co.FirstName)
		assert.Equal(t, "Doe", co.LastName)
		assert.Equal(t, "1234567890", co.PhoneNumber)
	}
}

func TestCreateContactValidation(t *testing.T) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(echo.POST, "/contacts", strings.NewReader(`{"lastName":"Doe","phoneNumber":"1234567890"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	seed := map[string]contact{}
	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.CreateContact(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestUpdateContact(t *testing.T) {
	sample := contact{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890"}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(echo.PUT, "/contacts/"+sample.ID, strings.NewReader(`{"firstName":"John1","lastName":"Doe1","phoneNumber":"12345678901"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sample.ID)

	seed := map[string]contact{sample.ID: sample}

	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.UpdateContact(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		var co contact
		co = seed[sample.ID]
		assert.Equal(t, "John1", co.FirstName)
		assert.Equal(t, "Doe1", co.LastName)
		assert.Equal(t, "12345678901", co.PhoneNumber)
	}
}

func TestUpdateContactValidation(t *testing.T) {
	sample := contact{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890"}

	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	req := httptest.NewRequest(echo.PUT, "/contacts/"+sample.ID, strings.NewReader(`{"lastName":"Doe1","phoneNumber":"12345678901"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sample.ID)

	seed := map[string]contact{sample.ID: sample}

	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.UpdateContact(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var co contact
		co = seed[sample.ID]
		assert.Equal(t, "John", co.FirstName)
		assert.Equal(t, "Doe", co.LastName)
		assert.Equal(t, "1234567890", co.PhoneNumber)
	}
}

func TestDeleteContact(t *testing.T) {
	sample := contact{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890"}

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/contacts/"+sample.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sample.ID)

	seed := map[string]contact{sample.ID: sample}

	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.DeleteContact(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		var co contact
		co = seed[sample.ID]
		assert.Empty(t, co)
	}
}

func TestGetContact(t *testing.T) {
	sample := contact{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890"}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/contacts/"+sample.ID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(sample.ID)

	seed := map[string]contact{sample.ID: sample}

	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.GetContact(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := rec.Result()
		defer rsp.Body.Close()
		b, _ := io.ReadAll(rsp.Body)
		assert.Contains(t, string(b), sample.ID)
		assert.Contains(t, string(b), "John")
		assert.Contains(t, string(b), "Doe")
		assert.Contains(t, string(b), "1234567890")
	}
}

func TestGetContacts(t *testing.T) {
	sample := contact{ID: uuid.New().String(), FirstName: "John", LastName: "Doe", PhoneNumber: "1234567890"}
	sample1 := contact{ID: uuid.New().String(), FirstName: "John1", LastName: "Doe1", PhoneNumber: "12345678901"}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/contacts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	seed := map[string]contact{
		sample.ID:  sample,
		sample1.ID: sample1,
	}

	h := &handler{repo: NewFakeContactRepository(seed)}

	if assert.NoError(t, h.GetContacts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		rsp := rec.Result()
		defer rsp.Body.Close()
		b, _ := io.ReadAll(rsp.Body)
		assert.Contains(t, string(b), sample.ID)
		assert.Contains(t, string(b), "John")
		assert.Contains(t, string(b), "Doe")
		assert.Contains(t, string(b), "1234567890")
		assert.Contains(t, string(b), sample1.ID)
		assert.Contains(t, string(b), "John1")
		assert.Contains(t, string(b), "Doe1")
		assert.Contains(t, string(b), "12345678901")
	}
}
