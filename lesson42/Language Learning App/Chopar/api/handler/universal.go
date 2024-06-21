package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) Universal(c *gin.Context) {
	url := "http://localhost:8080" + c.Request.URL.String()
	req, err := http.NewRequest(c.Request.Method, url, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request: %s", err),
		})
		log.Printf("Error with creating request: %s", err)
		return
	}

	resp, err := h.Client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Printf("%s\n", err)
		return
	}
	defer resp.Body.Close()

	message, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request: %s", err),
		})
		log.Printf("Error with creating request: %s", err)
		return
	}
	c.JSON(http.StatusOK, string(message))
}
