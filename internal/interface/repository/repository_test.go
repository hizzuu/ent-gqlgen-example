package repository

import (
	"os"
	"path/filepath"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/DATA-DOG/go-txdb"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/hizzuu/plate-backend/ent/enttest"
	"github.com/hizzuu/plate-backend/internal/infrastructure/db"
	"github.com/hizzuu/plate-backend/test/utils"
	"github.com/hizzuu/plate-backend/utils/path"
)

func TestMain(m *testing.M) {
	utils.ReadConf()

	txdb.Register("txdb", "mysql", db.MysqlDSN())

	client := enttest.Open(&testing.T{}, dialect.MySQL, db.MysqlDSN())
	defer client.Close()

	prepare()

	code := m.Run()
	os.Exit(code)
}

func prepare() {
	db, err := db.NewMysqlDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory(filepath.Join(path.RootDir(), "/test/fixtures/posts")),
		testfixtures.Directory(filepath.Join(path.RootDir(), "/test/fixtures/users")),
		testfixtures.Directory(filepath.Join(path.RootDir(), "/test/fixtures/images")),
	)
	if err != nil {
		panic(err)
	}

	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}
