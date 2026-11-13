package api

import (
	"apimockbuilder/projetos"
	"net/http"

	"github.com/gin-gonic/gin"
)

func get_corpo_string(c *gin.Context) string {
	corpo_bytes, err := c.GetRawData()
	if err != nil {
		return ""
	}
	return string(corpo_bytes)
}

//----

func config_debug(eng *gin.Engine) {
	eng.Any("/teste/*resto", func(c *gin.Context) {
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

func config_triviais(eng *gin.Engine) {
	eng.GET("/lista", func(c *gin.Context) {
		c.JSON(http.StatusOK, projetos.Lista_projetos())
	})

	eng.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "tudo ok")
	})
}

func config_projetos(eng *gin.Engine) {
	//criar projeto
	eng.POST("/projeto/:projeto", func(c *gin.Context) {
		nome_projeto := c.Param("projeto")

		if projetos.Novo_projeto(nome_projeto) != nil {
			c.String(http.StatusOK, "Projeto já existe!")
			return
		}

		c.String(http.StatusOK, "Criado novo projeto: "+nome_projeto)
	})

	//criar e atualizar endpoint no projeto
	eng.Any("/config/:projeto/*endpoint", func(c *gin.Context) {
		if c.Request.Method != "POST" && c.Request.Method != "PUT" {
			c.String(http.StatusMethodNotAllowed, "método sem suporte")
			return
		}

		projeto := c.Param("projeto")
		endpoint := c.Param("endpoint")
		corpo := get_corpo_string(c)

		err := projetos.Get_projeto(projeto)
		if err != nil {
			c.String(http.StatusNotFound, "projeto ainda nao existe..")
			return
		}

		if c.Request.Method != "POST" {
			err = projetos.Novo_endpoint(projeto, endpoint, corpo)
		} else {
			err = projetos.Mudar_endpoint(projeto, endpoint, corpo)
		}

		if err != nil {
			c.String(http.StatusInternalServerError, "deu ruim criando endpoint...")
			return
		}

		if c.Request.Method == "POST" {
			c.String(http.StatusOK, "endpoint criado.")
		} else {
			c.String(http.StatusOK,  "endpoint atualizado.")
		}

	})

	//acessar endpoint do projeto
	eng.Any("/live/:projeto/*endpoint", func(c *gin.Context) {
		projeto := c.Param("projeto")
		endpoint := c.Param("endpoint")

		err := projetos.Get_projeto(projeto)
		if err != nil {
			c.String(http.StatusOK, "projeto ainda nao existe..")
			return
		}

		//no futuro, listar endpoints disponiveis do projeto.
		c.String(http.StatusOK, "projeto funcionando: [ "+endpoint+" ]")
	})
}

func Start() {
	eng := gin.Default()

	config_debug(eng)
	config_triviais(eng)
	config_projetos(eng)

	eng.Run(":3000")
}
