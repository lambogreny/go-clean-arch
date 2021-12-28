package helpers

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
)

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
	//if value.Valid || value.String == "t" || value.String == "f" {

	if value.String == "t" || value.String == "f" {
		return value.String
	}
	if value.Valid {
		fmt.Println("Cai aqui no switch case")
		if value.String == "true" {
			return "t"
		}
		if value.String == "false" {
			return "f"
		}
	}
	return "f"
}

/*
	Função para casos de metadados string com a intenção de boolean no SQL
	Recebe um sql.NullString e valida se é válido
	Se for válida, transforma o false em 0 e true em 1
	Senão retorna 0
*/
func StringBooleanInt(value sql.NullString) int64 {
	if value.Valid {
		if value.String == "t" || value.String == "true" {
			return 1
		}

		if value.String == "f" || value.String == "false" {
			return 0
		}
	}
	return 0

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

/*
	Função para casos de metadados datatime no SQL
	Recebe um sql.NullString e valida se é válido
	Se for válida, faz o parse da data no formato correto do Mysql
	Senão faz o parse de uma data genérica no formato correto
*/
func StringDatetime(value sql.NullString) string {
	fmt.Println(value.String)
	layout := "2006-01-02 15:04:05"

	if value.Valid {
		v, _ := time.Parse(layout, value.String)
		v2 := v.Format("2006-01-02 15:04:05")
		return v2
	}

	genericDate, _ := time.Parse(layout, "2016-04-05 00:00:00")
	genericDate2 := genericDate.Format("2006-01-02 15:04:05")

	return genericDate2
}

/*
	Função para casos de metadados datatime no SQL
	Recebe um sql.NullString e valida se é válido
	Se for válida, faz o parse da data no formato correto do Mysql
	Senão faz o parse de uma data genérica no formato correto
*/
func StringInt(value sql.NullString) int64 {

	if value.Valid || value.String != "" {
		conv, _ := strconv.ParseInt(value.String, 10, 64)
		return conv
	}

	return 0
}
