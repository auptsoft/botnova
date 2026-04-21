package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ScalarDocsHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
            <!doctype html>
            <html>
            <head>
                <title>API Docs</title>
                <meta charset="utf-8" />
                <meta name="viewport" content="width=device-width, initial-scale=1" />
            </head>
            <body>
                <script
                    id="api-reference"
                    data-url="/swagger/doc.json"
                    src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
            </body>
            </html>
        `))
}
