package web

import (
	"fmt"
	"log"

	"encoding/json"
	"go/build"
	//	"io/ioutil"

	"net/http"
	//"net/http/httputil"

	"os/exec"

	"github.com/gorilla/mux"

	"github.com/haad/worktracker/model/customer"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
	router.HandleFunc("/customers", CustomerIndex)

	//	router.HandleFunc("/projects", index)
	//	router.HandleFunc("/entries/{id}", update)
	router.Handle("/", router.NotFoundHandler)
}

func StartServer(addr string) error {
	url := "http://" + addr + "/a/index.html"

	wtPkg, err := build.Import("github.com/haad/worktracker", "", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}

	http.Handle("/", router)
	http.Handle("/a/",
		http.StripPrefix(
			"/a/", http.FileServer(http.Dir(wtPkg.Dir+"/public")),
		),
	)

	exec.Command("open", url).Run()

	fmt.Printf("starting %s\n", url)

	return http.ListenAndServe(addr, nil)
}

func CustomerIndex(w http.ResponseWriter, req *http.Request) {
	var customers []customer.CustomerInt

	w.Header().Set("Content-Type", "application/json")

	customers = customer.CustomerList()

	b, err := json.Marshal(customers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, string(b))
}
