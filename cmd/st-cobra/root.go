package st

/*
O sistema deverá gerar um relatório com informações específicas após a execução dos testes.

Entrada de Parâmetros via CLI:

--url: URL do serviço a ser testado.
--requests: Número total de requests.
--concurrency: Número de chamadas simultâneas.
*/

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/winstonjr/goexpert-desafio-stress-test/internal/entity"
	"log"
	"net/url"
	"os"
)

var stressTestConfig *entity.StressTestConfig

var rootCmd = &cobra.Command{
	Use:   "go-stress-test",
	Short: "A stress tester CLI like Apache ab",
	Long:  `A stress tester CLI like Apache ab`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		flagUrl, _ := cmd.Flags().GetString("url")

		if flagUrl == "" {
			return fmt.Errorf("flag url is required")
		}
		_, err := url.ParseRequestURI(flagUrl)
		if err != nil {
			return fmt.Errorf("url flag is invalid")
		}

		flagConcurrency, _ := cmd.Flags().GetUint64("concurrency")
		flagRequests, _ := cmd.Flags().GetUint64("requests")

		if flagRequests <= uint64(0) {
			return fmt.Errorf("requests must be greater or equals 1")
		}

		if flagConcurrency <= uint64(0) {
			return fmt.Errorf("concurrency must be greater or equals 1")
		}

		if flagRequests < flagConcurrency {
			return fmt.Errorf("requests must be greater or equal concurrent")
		}

		stressTestConfig = entity.NewStressTestConfig(flagUrl, flagConcurrency, flagRequests)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		//err := s.Run()
		//if err != nil {
		//	panic(err)
		//}
		//s.PrintReport()

		log.Println("url: ", stressTestConfig.Url,
			"concurrency: ", stressTestConfig.Concurrency,
			"requests: ", stressTestConfig.TotalRequests)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var flagUrl string
	var flagConcurrency, flagRequests uint64
	rootCmd.Flags().StringVarP(&flagUrl, "url", "u", "", "URL to stress")
	rootCmd.Flags().Uint64VarP(&flagConcurrency, "concurrency", "c", 1, "URL to stress")
	rootCmd.Flags().Uint64VarP(&flagRequests, "requests", "r", 1, "URL to stress")
}
