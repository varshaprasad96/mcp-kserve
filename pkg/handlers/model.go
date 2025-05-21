package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/varshaprasad96/mcp-kserve/pkg/kserve"
)

// DeployModelRequest represents the request body for model deployment
type DeployModelRequest struct {
	Name      string `json:"name" binding:"required"`
	ModelURI  string `json:"modelURI" binding:"required"`
	Framework string `json:"framework" binding:"required"`
}

// ModelHandler handles model-related HTTP requests
type ModelHandler struct {
	kserveClient *kserve.Client
}

// NewModelHandler creates a new model handler
func NewModelHandler(kserveClient *kserve.Client) *ModelHandler {
	return &ModelHandler{
		kserveClient: kserveClient,
	}
}

// DeployModel handles model deployment requests
func (h *ModelHandler) DeployModel(c *gin.Context) {
	var req DeployModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.kserveClient.DeployModel(c.Request.Context(), req.Name, req.ModelURI, req.Framework)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Model deployment initiated",
		"name":    req.Name,
	})
}

// GetModelStatus handles model status requests
func (h *ModelHandler) GetModelStatus(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "model name is required"})
		return
	}

	service, err := h.kserveClient.GetModelStatus(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, service)
}

// ListModels handles model listing requests
func (h *ModelHandler) ListModels(c *gin.Context) {
	services, err := h.kserveClient.ListModels(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, services)
}

// DeleteModel handles model deletion requests
func (h *ModelHandler) DeleteModel(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "model name is required"})
		return
	}

	err := h.kserveClient.DeleteModel(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Model deleted successfully",
		"name":    name,
	})
}
