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

//#TODO
func PrintJson() {

	//u, err := json.Marshal(x)
	//
	//if err != nil {
	//	return nil, err
	//}
	//
	//fmt.Println(string(u))

}
