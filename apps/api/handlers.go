package main

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type handler struct {
	repo ContactRepository
}

func (h *handler) CreateContact(c echo.Context) error {
	payload := new(contact)
	if err := c.Bind(payload); err != nil {
		return err
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	payload.ID = uuid.New().String()
	if err := h.repo.Create(payload); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	c.Response().Header().Set(echo.HeaderLocation, "/contacts/"+payload.ID)

	return c.NoContent(http.StatusCreated)
}

func (h *handler) GetContacts(c echo.Context) error {
	contacts, err := h.repo.List()
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, contacts)
}

func (h *handler) GetContact(c echo.Context) error {
	id := c.Param("id")

	co, err := h.repo.Get(id)
	if errors.Is(err, ErrContactNotFound) {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, co)
}

func (h *handler) UpdateContact(c echo.Context) error {
	id := c.Param("id")

	payload := new(contact)
	if err := c.Bind(payload); err != nil {
		return err
	}

	if err := c.Validate(payload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.repo.Update(id, payload)
	if errors.Is(err, ErrContactNotFound) {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *handler) DeleteContact(c echo.Context) error {
	id := c.Param("id")

	err := h.repo.Delete(id)
	if errors.Is(err, ErrContactNotFound) {
		return c.NoContent(http.StatusNotFound)
	}
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusNoContent)
}
