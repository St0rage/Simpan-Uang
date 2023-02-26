package manager

import (
	"fmt"
	"log"

	"github.com/St0rage/Simpan-Uang/config"
	"github.com/St0rage/Simpan-Uang/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type InfraManager interface {
	DbConn() *sqlx.DB
}

type infraManager struct {
	db  *sqlx.DB
	cfg config.Config
}

func (infra *infraManager) initDb() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", infra.cfg.Host, infra.cfg.Port, infra.cfg.User, infra.cfg.Password, infra.cfg.Name)
	db, err := sqlx.Connect("postgres", psqlconn)
	utils.PanicIfError(err)
	defer func() {
		if err := recover(); err != nil {
			log.Println("Application Failed to run", err)
		}
	}()

	err = db.Ping()
	utils.PanicIfError(err)

	infra.db = db
}

func (infra *infraManager) DbConn() *sqlx.DB {
	return infra.db
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initDb()
	return &infra
}
