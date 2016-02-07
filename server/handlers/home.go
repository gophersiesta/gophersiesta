package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHome(c *gin.Context) {

	logo := `
.      .__                 __
  _____|__| ____   _______/  |______
 /  ___/  |/ __ \ /  ___/\   __\__  \
 \___ \|  \  ___/ \___ \  |  |  / __ \_
/____  >__|\___  >____  > |__| (____  /
     \/        \/     \/            \/
`
	c.String(http.StatusOK, logo)

}
