// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"awesome-project/internal/dao"
	"awesome-project/internal/service"
	"awesome-project/internal/server/grpc"
	"awesome-project/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
