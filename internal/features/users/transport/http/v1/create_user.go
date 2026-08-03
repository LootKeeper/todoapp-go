package users_transport_http_v1

import (
	"net/http"

	"github.com/LootKeeper/todoapp-go/internal/core/domain"
	core_logger "github.com/LootKeeper/todoapp-go/internal/core/logger"
	core_http_request "github.com/LootKeeper/todoapp-go/internal/core/transport/http/request"
	core_http_response "github.com/LootKeeper/todoapp-go/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	Name        string  `json:"name" validate:"required,min=3,max=100"`
	Email       string  `json:"email" validate:"required,min=3,max=150,email"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min10,max=15,startswith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPRequestHandler(logger, rw)

	logger.Debug("invoke create user handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request body")
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.Name, dto.Email, dto.PhoneNumber)
}
