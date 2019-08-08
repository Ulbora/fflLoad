package main




import()

func main(){
	router := mux.NewRouter()
	port := "3000"	

	fmt.Println("Starting server Oauth2 Server on " + port)
	http.ListenAndServe(":"+port, router)
}