package usecase

import "github.com/winstonjr/goexpert-desafio-stress-test/internal/entity"

type ExecuteStressTestUseCase struct{}

func NewExecuteStressTestUseCase() *ExecuteStressTestUseCase {
	return &ExecuteStressTestUseCase{}
}

func (uc *ExecuteStressTestUseCase) ExecuteStressTest(input *entity.StressTestConfig) (*entity.Results, error) {
	var requestsPerConcurrency uint64
	var remainderRequests uint64

}
