package handlers

import (
	//"bytes"
	"net/http"

	"Cloud-Log-Access-Service/aws/services"

	"github.com/gin-gonic/gin"
)

// S3Handler é a interface para handlers que interagem com o AWS S3.
type S3Handler interface {
	ListBucketObjects(c *gin.Context)
	GetObjectFromBucket(c *gin.Context)
}

// s3HandlerImpl implementa a interface S3Handler.
type s3HandlerImpl struct {
	s3Service services.S3Service
}

// NewS3Handler cria uma nova instância de S3Handler.
func NewS3Handler(s3Service services.S3Service) S3Handler {
	return &s3HandlerImpl{
		s3Service: s3Service,
	}
}

// ListBucketObjects godoc
// @Summary Lista objetos em um bucket S3
// @Description Retorna uma lista de objetos em um bucket S3 especificado
// @Tags s3
// @Accept json
// @Produce json
// @Param bucketName path string true "Nome do bucket S3"
// @Success 200 {array} object "Lista de objetos S3"
// @Failure 400 {object} map[string]string "Erro de requisição inválida"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /s3/bucket/{bucketName}/objects [get]
func (h *s3HandlerImpl) ListBucketObjects(c *gin.Context) {
	bucketName := c.Param("bucketName")
	if bucketName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome do bucket é obrigatório"})
		return
	}

	objects, err := h.s3Service.ListObjectsInBucket(bucketName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar objetos do bucket", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, objects)
}

// GetObjectFromBucket godoc
// @Summary Obtém um objeto de um bucket S3
// @Description Retorna o conteúdo de um objeto específico de um bucket S3
// @Tags s3
// @Accept json
// @Produce octet-stream
// @Param bucketName path string true "Nome do bucket S3"
// @Param objectKey path string true "Chave do objeto S3"
// @Success 200 {string} string "Conteúdo do objeto"
// @Failure 400 {object} map[string]string "Erro de requisição inválida"
// @Failure 404 {object} map[string]string "Objeto não encontrado"
// @Failure 500 {object} map[string]string "Erro interno do servidor"
// @Router /s3/bucket/{bucketName}/object/{objectKey} [get]
func (h *s3HandlerImpl) GetObjectFromBucket(c *gin.Context) {
	bucketName := c.Param("bucketName")
	objectKey := c.Param("objectKey")

	if bucketName == "" || objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome do bucket e chave do objeto são obrigatórios"})
		return
	}

	content, err := h.s3Service.GetObjectFromBucket(bucketName, objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter objeto do bucket", "details": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/octet-stream", content)
}
