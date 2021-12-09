package api

func Init() {
	r := NewRouter()
	r.Run(":8080")
}
