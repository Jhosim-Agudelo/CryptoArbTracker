package main

import ("net/http"
)

func main(){
	http.HandleFunc("/cryptoarbtracker",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("CryptoArbTracker"))
	})
	http.ListenAndServe(":8080",nil)
}