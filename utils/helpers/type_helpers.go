package helpers

import "database/sql"

/*
	Recebe um sql.NullString e valida se a string é válida
	Se for válida, devolve o próprio valor, senão dovolve vazio
*/
func String(value sql.NullString) string {
	if value.Valid {
		return value.String
	}
	return ""
}

/*
	Recebe um sql.NullInt64 e valida se o int é válida
	Se for válida, devolve o próprio valor, senão dovolve 0
*/
func Int(value sql.NullInt64) int64 {
	if value.Valid {
		return value.Int64
	}
	return 0
}

/*
	Função para casos de metadados booleano no SQL
	Recebe um sql.NullString e valida se o string é válida
	Se for válida, devolve o próprio valor, senão dovolve "f'
*/
func StringBoolean(value sql.NullString) string {
	if value.Valid || value.String == "t" || value.String == "f" {
		return value.String
	}
	return "f"
}

/*
	Função para casos de metadados float64 no SQL
	Recebe um sql.NullFloat64 e valida se o número é válido
	Se for válida, devolve o próprio valor, senão dovolve 0.0
*/
func Float64(value sql.NullFloat64) float64 {
	if value.Valid {
		return value.Float64
	}
	return 0.0
}
