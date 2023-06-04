package utils

import (
	"fmt"
	"github.com/syahrilmaulayahya/tugas_akhir_rakamin/internal/pkg/dto"
	"runtime"
	"sync"
	"testing"
)

func BenchmarkCallAPI(b *testing.B) {
	b.Run("With goroutine", func(b *testing.B) {

		// fetching from 3rd party api
		var wg sync.WaitGroup
		// get province data from 3rd party api
		var getProvince dto.Province
		var errGetProvince error
		wg.Add(1)
		go func() {
			defer wg.Done()
			getProvince, errGetProvince = GetProvince("11")
			fmt.Println(getProvince, errGetProvince)
			fmt.Println(runtime.NumGoroutine())

		}()
		// get regency data from 3rd party api
		var getRegency dto.Regency
		var errGetRegency error
		wg.Add(1)
		go func() {
			defer wg.Done()
			getRegency, errGetRegency = GetRegency("1101")
			fmt.Println(getRegency, errGetRegency)
			fmt.Println(runtime.NumGoroutine())

		}()
		wg.Wait()

	})
	b.Run("Without goroutine", func(b *testing.B) {

		getProvince, errGetProvince := GetProvince("11")
		fmt.Println(getProvince, errGetProvince)
		fmt.Println(runtime.NumGoroutine())
		getRegency, errGetRegency := GetRegency("1101")
		fmt.Println(getRegency, errGetRegency)
		fmt.Println(runtime.NumGoroutine())

	})

}

func TestGetAllProvince(t *testing.T) {
	result, err := GetAllProvince()
	fmt.Println(result, err)
}
