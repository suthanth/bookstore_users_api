package app

func StartApplication() {
	router := NewRouter()
	router.Run(":8000")
}
