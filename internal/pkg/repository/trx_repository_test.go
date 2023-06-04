package repository

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/daos"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/container"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/infrastructure/mysql"
	"sync"
	"testing"
)

// test race condition stok
func TestCreateTRX(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	repo := NewTRXRepository(containerConf.Mysqldb)
	//repo.CreateTRX(context.Background(), daos.TRX{}, []int{1, 2, 3})
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response, err := repo.CreateTRX(context.Background(), daos.TRX{
				UserID:      1,
				AlamatID:    1,
				HargaTotal:  0,
				KodeInvoice: "INV-1670061868",
				MethodBayar: "bca",
			}, []daos.ProdukIDKuantitas{
				{
					ProdukID:  1,
					Kuantitas: 100,
				},
				{
					ProdukID:  2,
					Kuantitas: 100,
				},
				{
					ProdukID:  3,
					Kuantitas: 100,
				},
			})
			fmt.Println(err)
			fmt.Println(response)
		}()
	}
	wg.Wait()

}

func TestGetProductID(t *testing.T) {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	repo := NewTRXRepository(containerConf.Mysqldb)

	trx, err := repo.GetTRXByID(context.Background(), 2, 1)
	fmt.Println(err)
	fmt.Println(trx)
	fmt.Println(trx.DetailTRX)

}
