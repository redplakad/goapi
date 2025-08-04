package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redplakad/goapi/models"
)

type NominatifKreditHandler struct {
	repo *models.NominatifKreditRepository
}

func NewNominatifKreditHandler(repo *models.NominatifKreditRepository) *NominatifKreditHandler {
	return &NominatifKreditHandler{repo: repo}
}

// APIResponse represents a standard API response
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// PaginationMeta represents pagination metadata
type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
}

// GetAll handles GET /api/nominatif-kredit
func (h *NominatifKreditHandler) GetAll(c *gin.Context) {
	// Get pagination parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	perPage, _ := strconv.Atoi(c.DefaultQuery("per_page", "10"))
	
	if page < 1 {
		page = 1
	}
	if perPage < 1 || perPage > 100 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	// Get filter parameters
	filters := models.FilterParams{
		CAB:            c.Query("cab"),
		AO:             c.Query("ao"),
		KET_KD_PRD:     c.Query("ket_kd_prd"),
		TEMPAT_BEKERJA: c.Query("tempat_bekerja"),
	}

	// Get data with filters
	kredits, err := h.repo.GetAllWithFilters(perPage, offset, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to fetch nominatif kredit data",
		})
		return
	}

	// Get total count for pagination with filters
	total, err := h.repo.CountWithFilters(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to get total count",
		})
		return
	}

	lastPage := (total + perPage - 1) / perPage

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Nominatif kredit data retrieved successfully",
		Data:    kredits,
		Meta: PaginationMeta{
			CurrentPage: page,
			PerPage:     perPage,
			Total:       total,
			LastPage:    lastPage,
		},
	})
}

// GetByID handles GET /api/nominatif-kredit/:id
func (h *NominatifKreditHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Invalid ID format",
		})
		return
	}

	kredit, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "Nominatif kredit not found",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Nominatif kredit data retrieved successfully",
		Data:    kredit,
	})
}

// GetByNomorRekening handles GET /api/nominatif-kredit/rekening/:nomor_rekening
func (h *NominatifKreditHandler) GetByNomorRekening(c *gin.Context) {
	nomorRekening := c.Param("nomor_rekening")
	if nomorRekening == "" {
		c.JSON(http.StatusBadRequest, APIResponse{
			Success: false,
			Message: "Nomor rekening is required",
		})
		return
	}

	kredits, err := h.repo.GetByNomorRekening(nomorRekening)
	if err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Success: false,
			Message: "Failed to fetch nominatif kredit data",
		})
		return
	}

	if len(kredits) == 0 {
		c.JSON(http.StatusNotFound, APIResponse{
			Success: false,
			Message: "No data found for the given nomor rekening",
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Success: true,
		Message: "Nominatif kredit data retrieved successfully",
		Data:    kredits,
	})
}
