package handlers

import (
	"net/http"

	"github.com/Anshualawa/school-management/internal/services"
	"github.com/Anshualawa/school-management/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ Service *services.UserService }

func NewAuthorHandler(s *services.UserService) *AuthHandler { return &AuthHandler{Service: s} }

type SignupReq struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Role     string `json:"role,omitempty"`
}

func (h *AuthHandler) Signup(g *gin.Context) {
	var req SignupReq

	if err := g.Bind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload " + err.Error()})
		return
	}

	if utils.IsEmpty(req.Name) || utils.IsEmpty(req.Email) || utils.IsEmpty(req.Password) {
		g.JSON(http.StatusBadRequest, gin.H{"error": "email,name,password are required"})
		return
	}

	u, err := h.Service.Register(req.Name, req.Email, req.Password, req.Role)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, utils.APIResponse{Success: true, Message: "user create successfull", Data: u})
}

func (h *AuthHandler) Login(g *gin.Context) {
	var req SignupReq

	if err := g.Bind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload " + err.Error()})
		return
	}

	if utils.IsEmpty(req.Email) || utils.IsEmpty(req.Password) {
		g.JSON(http.StatusBadRequest, gin.H{"error": "email,password are required"})
		return
	}

	_, token, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		g.JSON(http.StatusInternalServerError, utils.APIResponse{Success: false, Message: "invalid creadential", Data: map[string]any{"error": err.Error()}})
	}

	g.JSON(http.StatusOK, utils.APIResponse{Success: true, Message: "user login successfull", Data: map[string]string{"token": token}})
}
