package listobjects

import (
	"testing"

	openfgav1 "github.com/openfga/api/proto/openfga/v1"
	"go.uber.org/goleak"

	"github.com/openfga/openfga/internal/server/config"
	"github.com/openfga/openfga/pkg/testutils"
	"github.com/openfga/openfga/tests"
)

func TestListObjectsPostgres(t *testing.T) {
	testRunAll(t, "postgres")
}

func TestListObjectsMySQL(t *testing.T) {
	testRunAll(t, "mysql")
}

func testRunAll(t *testing.T, engine string) {
	t.Cleanup(func() {
		goleak.VerifyNone(t)
	})
	cfg := config.MustDefaultConfig()
	cfg.Experimentals = append(cfg.Experimentals, "enable-check-optimizations")
	cfg.Log.Level = "error"
	cfg.Datastore.Engine = engine

	tests.StartServer(t, cfg)

	conn := testutils.CreateGrpcConnection(t, cfg.GRPC.Addr)

	RunAllTests(t, openfgav1.NewOpenFGAServiceClient(conn))
}
