package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/slack-go/slack"
)

const token = "xoxb-737504271157-3309190307425-mH7f60v7bzfaOcInR8Nx0PxZ"
const channelID = "CMC1L92R1"

type challageRequest struct {
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
	Type      string `json:"type"`
}

type response struct {
	Challenge string `json:"challenge"`
}

func echo(w http.ResponseWriter, req *http.Request) {
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(bodyBytes))
	st := challageRequest{}
	err = json.Unmarshal(bodyBytes, &st)
	if err != nil {
		fmt.Println(err.Error())
	}
	r := response{Challenge: st.Challenge}
	js, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(js))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	api := slack.New(token, slack.OptionDebug(true))
	_ = api

}

func main() {
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8090", nil)
}
