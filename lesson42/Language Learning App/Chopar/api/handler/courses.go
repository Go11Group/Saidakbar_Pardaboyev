package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) Create(c *gin.Context) {
	body := c.Request.Body
	url := "http://localhost:8080/" + c.Request.URL.String()
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request for create: %s", err),
		})
		log.Printf("Error with creating request for create: %s", err)
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
			"message": fmt.Sprintf("Error with creating request for create: %s", err),
		})
		log.Printf("Error with creating request for create: %s", err)
		return
	}
	c.JSON(http.StatusOK, string(message))
}

// Read all
func (h *Handler) Get(c *gin.Context) {
	url := "http://localhost:8080" + c.Request.URL.String()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request for get: %s", err),
		})
		log.Printf("Error with creating request for get: %s", err)
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
			"message": fmt.Sprintf("Error with getting data from the url: %s", err),
		})
		log.Printf("Error with getting data from the url: %s", err)
		return
	}
	c.JSON(http.StatusOK, string(message))
}

// Update
func (h *Handler) Update(c *gin.Context) {
	url := "http://localhost:8080" + c.Request.URL.String()
	body := c.Request.Body

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request for update: %s", err),
		})
		log.Printf("Error with creating request for update: %s", err)
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
			"message": fmt.Sprintf("Error with getting data from the url: %s", err),
		})
		log.Printf("Error with getting data from the url: %s", err)
		return
	}
	c.JSON(http.StatusOK, string(message))
}

// Delete
func (h *Handler) Delete(c *gin.Context) {
	url := "http://localhost:8080" + c.Request.URL.String()
	body := c.Request.Body

	req, err := http.NewRequest("DELETE", url, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "StatusBadRequest",
			"message": fmt.Sprintf("Error with creating request for delete: %s", err),
		})
		log.Printf("Error with creating request for delete: %s", err)
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
			"message": fmt.Sprintf("Error with getting data from the url: %s", err),
		})
		log.Printf("Error with getting data from the url: %s", err)
		return
	}
	c.JSON(http.StatusOK, string(message))
}
