package entity

import "time"

/*
Apresentar um relatório ao final dos testes contendo:
Tempo total gasto na execução
Quantidade total de requests realizados.
Quantidade de requests com status HTTP 200.
Distribuição de outros códigos de status HTTP (como 404, 500, etc.).
*/

type Results struct {
	TotalTime     time.Duration
	TotalRequests uint64
	TotalSuccess  uint64
	TotalFailure  map[uint64]uint64
}
