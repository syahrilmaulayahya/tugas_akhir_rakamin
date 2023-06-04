package mysql

import (
	"fmt"

	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/helper"
	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.User{}, &daos.Toko{}, &daos.Category{}, &daos.Alamat{}, &daos.Produk{}, &daos.FotoProduk{}, &daos.LogProduk{}, &daos.TRX{}, &daos.DetailTRX{}, &daos.LogFotoProduk{},
	)

	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelError, fmt.Sprintf("Database Migration Failed : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Database Migrated")
}
