package handler

import (
	"route-planner/api/internal/viewmodel"
	"route-planner/model"
	"route-planner/pkg/errorx"
	bmodel "route-planner/pkg/model"
	"route-planner/pkg/utils"
	"route-planner/service"

	"github.com/gofiber/fiber/v2"
)

type LocationHandler struct {
	LocationService service.ILocationService
}

func NewLocationHandler(LocationService service.ILocationService) LocationHandler {
	return LocationHandler{LocationService: LocationService}
}

func (h LocationHandler) GetAll(ctx *fiber.Ctx) error {

	data, dataCount, err := h.LocationService.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}
	var vm []viewmodel.LocationResponse
	for _, d := range data {
		vm = append(vm, viewmodel.LocationResponse{}.ToViewModel(d))
	}
	return ctx.Status(fiber.StatusOK).JSON(bmodel.Response{Data: vm, DataCount: dataCount})
}

func (h LocationHandler) GetByID(ctx *fiber.Ctx) error {
	id := utils.StrToInt64(ctx.Params("id"))

	location, err := h.LocationService.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}

	return ctx.Status(fiber.StatusOK).JSON(bmodel.Response{Data: location})
}

func (h LocationHandler) Create(ctx *fiber.Ctx) error {
	var vm viewmodel.LocationRequest
	if err := ctx.BodyParser(&vm); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(bmodel.Response{Error: err.Error()})
	}

	m := vm.ToDBModel(model.Location{})

	err := h.LocationService.Create(ctx.Context(), &m)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (h LocationHandler) Update(ctx *fiber.Ctx) error {
	id := utils.StrToInt64(ctx.Params("id"))

	location, err := h.LocationService.GetByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}

	var vm viewmodel.LocationRequest
	if err := ctx.BodyParser(&vm); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(bmodel.Response{Error: err.Error()})
	}

	m := vm.ToDBModel(location)

	err = h.LocationService.Update(ctx.Context(), m)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}

func (h LocationHandler) Delete(ctx *fiber.Ctx) error {
	id := utils.StrToInt64(ctx.Params("id"))

	err := h.LocationService.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(bmodel.Response{Error: errorx.InternalServerError})
	}

	return ctx.Status(fiber.StatusOK).JSON(nil)
}
