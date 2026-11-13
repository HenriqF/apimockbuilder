package api

import (
	"apimockbuilder/projetos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func config_endpoints(eng *gin.Engine) {

	eng.GET("/lista", func(c *gin.Context) {
		c.JSON(http.StatusOK, projetos.Lista_projetos())
	})

	eng.POST("/novo/:nome", func(c *gin.Context) {
		nome_projeto := c.Param("nome")

		if projetos.Novo_projeto(nome_projeto) != nil {
			c.String(http.StatusOK, "Projeto já existe!")
			return
		}

		c.String(http.StatusOK, "Criado novo projeto: "+nome_projeto)
	})

	eng.Any("/live/*resto", func(c *gin.Context) {
		metodo := c.Request.Method
		caminho := c.Param("resto") // c.Request.URL.Path
		parametro := c.Query("info")
		header := c.GetHeader("Content-Type")

		corpo_bytes, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Corpo podre")
			return
		}
		corpo := string(corpo_bytes)

		c.JSON(http.StatusOK, gin.H{
			"corpo":     corpo,
			"metodo":    metodo,
			"caminho":   caminho,
			"parametro": parametro,
			"header":    header,
		})
	})

}

func Start() {
	eng := gin.Default()

	config_endpoints(eng)

	eng.Run(":3000")
}
