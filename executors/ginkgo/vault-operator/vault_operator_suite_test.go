package vault_operator_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/lmittmann/tint"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestVaultOperator(t *testing.T) {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))
	slog.SetDefault(logger)

	RegisterFailHandler(Fail)
	RunSpecs(t, "VaultOperator Suite")
}
