package api

import (
	"errors"
	db "github.com/Dennisblay/ordering-app-server/internal/database/sqlc"
	"github.com/Dennisblay/ordering-app-server/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"net/http"
)

// getUsersController gets all users by pagination
func (s *Server) getUsersController(ctx *gin.Context) {
	var req db.UsersRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := db.GetUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := s.store.GetUsers(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, users)
	return
}

// getUserByIDController gets a user by Id
func (s *Server) getUserByIDController(ctx *gin.Context) {
	var req db.UserRequestByID
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserById(ctx, req.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
	return
}

// getUserControllerByEmailOrPhone gets user by email or phone
func (s *Server) getUserControllerByEmailOrPhone(ctx *gin.Context) {
	var req db.UserRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	args := db.GetUserByEmailOrPasswordParams{
		ID:    req.ID,
		Email: req.Email,
		Phone: req.Phone,
	}
	user, err := s.store.GetUserByEmailOrPassword(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

// getUserByEmailAndPasswordController get user by Email and Password
func (s *Server) getUserByEmailAndPasswordController(ctx *gin.Context) {
	var req db.UserRequestByEmailAndPassword
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.GetUserByEmailAndPassword(ctx, db.GetUserByEmailAndPasswordParams{Email: req.Email, PasswordHash: req.Password})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

// createUserController creates a new user
func (s *Server) createUserController(ctx *gin.Context) {
	var req db.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	hashPassword, err := util.HashPassword(req.Password)
	arg := db.CreateUserParams{
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Email:        req.Email,
		PasswordHash: hashPassword,
		Phone:        req.Phone,
		Address:      pgtype.Text{String: req.Address},
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		errCode := db.ErrorCode(err)
		if errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, rsp)
	return
}

// updateUserNameController updates user's name
func (s *Server) updateUserNameController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var req db.UpdateUserNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserName(ctx, db.UpdateUserNameParams{ID: reqID.ID, FirstName: req.FirstName, LastName: req.LastName})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)
	return
}

// updateUserEmailController update user email
func (s *Server) updateUserEmailController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var req db.UpdateUserEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserEmail(ctx, db.UpdateUserEmailParams{ID: reqID.ID, Email: req.Email})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)
	return
}

// updateUserPhoneController update user phone
func (s *Server) updateUserPhoneController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var req db.UpdateUserPhoneRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserPhone(ctx, db.UpdateUserPhoneParams{ID: reqID.ID, Phone: req.Phone})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)
	return
}

// updateUserAddressController updates user address
func (s *Server) updateUserAddressController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	var req db.UpdateUserAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	user, err := s.store.UpdateUserAddress(ctx, db.UpdateUserAddressParams{ID: reqID.ID, Address: pgtype.Text{String: req.Address}})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)
	return
}

// updateUserPasswordController updates user password
func (s *Server) updateUserPasswordController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
	}

	var reqPassword db.UpdateUserPasswordRequest
	if err := ctx.ShouldBind(&reqPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	// first Check if old password == to existing password
	args := db.UpdateUserPasswordParams{
		ID:           reqID.ID,
		PasswordHash: reqPassword.PasswordNew,
	}
	user, err := s.store.UpdateUserPassword(ctx, args)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}
	rsp := db.UserResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, rsp)
	return
}

// deleteUserController deletes user
func (s *Server) deleteUserController(ctx *gin.Context) {
	var reqID db.UserRequestByID
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	err := s.store.DeleteUser(ctx, reqID.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}
	ctx.Status(http.StatusNoContent)
	return
}
