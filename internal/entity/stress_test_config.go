package entity

type StressTestConfig struct {
	Url           string
	Concurrency   uint64
	TotalRequests uint64
}

func NewStressTestConfig(url string, concurrency, requests uint64) *StressTestConfig {
	return &StressTestConfig{
		Url:           url,
		Concurrency:   concurrency,
		TotalRequests: requests,
	}
}

func (st *StressTestConfig) GetRequestsPerConcurrentExecution() (uint64, uint64) {
	if st.Concurrency == 1 {
		return st.TotalRequests, 0
	} else if st.Concurrency > 1 {
		reqs := st.TotalRequests / st.Concurrency
		remaining := st.TotalRequests % st.Concurrency

		return reqs, remaining
	} else {
		return 0, 0
	}
}
