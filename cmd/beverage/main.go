package beverage

import (
	"context"
	"os"
	"path"

	"github.com/logrhythm/go-framework-poc/internal/pkg/authz"
	"github.com/logrhythm/go-framework-poc/internal/pkg/beverage/svc"
)

func main() {

	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	rootDir := path.Dir(exePath)

	authorizer, err := authz.NewAuthorizer(path.Join(rootDir, "model.conf"), path.Join(rootDir, "policy.csv"))
	if err != nil {
		panic(err)
	}
	beerService := svc.NewBeerServer()
	authorizer.Authorize(context.Background(), beerService.GetBeer)
}
