package apis

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var (
	invalidRequestError = errors.New("Invalid request.")
	unauthorizedError   = errors.New("Unauthorized.")
	forbiddenError      = errors.New("Forbidden.")
	notFoundError       = errors.New("Not found.")
	unimplementedError  = errors.New("Unimplemented.")
)

func handleError(
	app *pocketbase.PocketBase,
	e *core.ServeEvent,
	c echo.Context,
	err error,
) error {
	if errors.Is(err, invalidRequestError) {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Invalid request.",
			},
		)
	}

	if errors.Is(err, unauthorizedError) {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Unauthorized.",
			},
		)
	}

	if errors.Is(err, forbiddenError) {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Forbidden.",
			},
		)
	}

	if errors.Is(err, notFoundError) {
		return c.JSON(
			http.StatusNotFound,
			errorResponse{
				response: response{
					Success: false,
				},
				Message: "Not found.",
			},
		)
	}

	app.Logger().Error(
		"Internal server error occur.",
		"route", c.RouteInfo().Path(),
		"pathParams", c.PathParams(),
		"queryParams", c.QueryParams(),
		"error", err,
	)
	return c.JSON(
		http.StatusInternalServerError,
		errorResponse{
			response: response{
				Success: false,
			},
			Message: "Internal server error.",
		},
	)
}
