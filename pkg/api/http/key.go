package http

import (
	"context"
	"net/http"
	"time"

	"github.com/SigNoz/signoz-otel-collector/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type KeyAPI struct {
	base    *API
	storage *storage.Storage
}

func NewKeyAPI(base *API, storage *storage.Storage) *KeyAPI {
	return &KeyAPI{
		base:    base,
		storage: storage,
	}
}

func (api *KeyAPI) Register() {
	group := api.base.Handler().Group("/keys")

	group.GET("", func(gctx *gin.Context) {
		_, cancel := context.WithTimeout(gctx, 5*time.Second)
		defer cancel()
	})

	group.POST("", func(gctx *gin.Context) {
		ctx, cancel := context.WithTimeout(gctx, 5*time.Second)
		defer cancel()

		type Req struct {
			Tenant struct {
				Name string `json:"name" binding:"required"`
			} `json:"tenant" binding:"required"`
			Name      string    `json:"name" binding:"required"`
			ExpiresAt time.Time `json:"expires_at" binding:"required"`
		}

		req := new(Req)
		if err := gctx.MustBindWith(req, binding.JSON); err != nil {
			api.base.SendErrorResponse(gctx, err)
			return
		}

		tenant, err := api.storage.DAO.Tenants().SelectByName(ctx, req.Tenant.Name)
		if err != nil {
			api.base.SendErrorResponse(gctx, err)
			return
		}

		type Res struct {
			Id string `json:"id"`
		}

		api.base.SendSuccessResponse(gctx, http.StatusCreated, Res{Id: tenant.Id.String()})
	})

	group.DELETE("", func(gctx *gin.Context) {
		_, cancel := context.WithTimeout(gctx, 5*time.Second)
		defer cancel()
	})

}
