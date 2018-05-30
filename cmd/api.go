package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lostsnow/ipd/ip"
	"github.com/spf13/cobra"
)

var host string
var port int
var route string

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "ipd api",
	Long:  `ipd api`,
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.HandleFunc("/"+route, IpHandler)
		addr := fmt.Sprintf("%s:%d", host, port)

		log.Printf("listen on %s/%s", addr, route)
		log.Fatal(http.ListenAndServe(addr, r))
	},
}

func init() {
	RootCmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "bind address")
	apiCmd.Flags().IntVarP(&port, "port", "p", 22033, "bind port")
	apiCmd.Flags().StringVarP(&route, "route", "r", "ip", "route")
}

func IpHandler(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("p")
	if param == "" {
		param = "g_locinfo"
	}

	format := r.URL.Query().Get("f")
	if format == "" {
		format = "js"
	}

	jsonpParam := r.URL.Query().Get("pp")

	ipv := r.URL.Query().Get("_ip_")
	if ipv == "" {
		ipv = r.Header.Get("X-Real-IP")
		if ipv == "" {
			v, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				fmt.Printf("%s\n", err)
				return
			}
			ipv = v
		}
	}

	l, err := cfg.Db.Find(ipv)
	if err != nil {
		fmt.Printf("%s %s\n", ipv, err)
		return
	}

	type LocationResponse ip.Location
	jsonByte, err := json.Marshal(l)
	if err != nil {
		fmt.Printf("%s %s\n", ipv, err)
		return
	}

	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	jsonStr := string(jsonByte)
	if format == "js" {
		jsonStr = fmt.Sprintf("var %s=%s;", param, jsonStr)
	} else if format == "jsonp" && jsonpParam != "" {
		jsonStr = fmt.Sprintf("%s(%s);", jsonpParam, jsonStr)
	} else {
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonStr)
}
