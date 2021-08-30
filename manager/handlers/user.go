/*
 *     Copyright 2020 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package handlers

import (
	"net/http"

	"d7y.io/dragonfly/v2/manager/types"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp user
// @Description signup by json config
// @Tags User
// @Accept json
// @Produce json
// @Param User body types.SignUpRequest true "User"
// @Success 200 {object} model.User
// @Failure 400
// @Failure 500
// @Router /user/signup [post]
func (h *Handlers) SignUp(ctx *gin.Context) {
	var json types.SignUpRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	user, err := h.Service.SignUp(json)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

// @Summary Reset Password For User
// @Description reset password by json config
// @Tags User
// @Accept json
// @Produce json
// @Param User body types.ResetPasswordRequest true "User"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/:id/reset_password [post]
func (h *Handlers) ResetPassword(ctx *gin.Context) {
	var params types.UserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	var json types.ResetPasswordRequest
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if err := h.Service.ResetPassword(params.ID, json); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Get User Roles
// @Description get roles by json config
// @Tags User
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} []string
// @Failure 400
// @Failure 500
// @Router /users/:id/roles [get]
func (h *Handlers) GetRolesForUser(ctx *gin.Context) {
	var params types.UserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	roles, err := h.Service.GetRolesForUser(params.ID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, roles)
}

// @Summary Add Role For User
// @Description add role to user by uri config
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/:id/roles/:role [put]
func (h *Handlers) AddRoleToUser(ctx *gin.Context) {
	var params types.AddRoleForUserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if ok, err := h.Service.AddRoleForUser(params); err != nil {
		ctx.Error(err)
		return
	} else if !ok {
		ctx.Status(http.StatusConflict)
		return
	}

	ctx.Status(http.StatusOK)
}

// @Summary Delete Role For User
// @Description delete role by uri config
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param role path string true "role"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /users/:id/roles/:role [delete]
func (h *Handlers) DeleteRoleForUser(ctx *gin.Context) {
	var params types.DeleteRoleForUserParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	if ok, err := h.Service.DeleteRoleForUser(params); err != nil {
		ctx.Error(err)
		return
	} else if !ok {
		ctx.Status(http.StatusNotFound)
		return
	}

	ctx.Status(http.StatusOK)
}
