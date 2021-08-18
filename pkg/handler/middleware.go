package handler

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationHeader = "Authorization"

func (h *Handler) userIdentity(c *gin.Context) {

	// header
	header := c.GetHeader(authorizationHeader)

	// check header exists
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, 2000, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	// check Bearer in headers
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, 2010, "Invalid auth header")
		return
	}

	// check token exists
	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, 2010, "Token is empty")
		return
	}

	// Ccheck token
	if headerParts[1] != os.Getenv("AUTH_TOKEN") {
		newErrorResponse(c, http.StatusUnauthorized, 3000, "Invalid auth token")
		return
	}

}
