package main

import (
	"log"
	"net/http"
	"os"
)

const base = "https://github.com/zerbazaar/releases/releases"

var routes = map[string]string{
	"/sunucu/windows": base + "/latest/download/zerbazaar-sunucu-kurulum.exe",
	"/sunucu/linux":   base + "/latest/download/zerbazaar-server_amd64.deb",
	"/sunucu/macos":   base + "/latest/download/zerbazaar-sunucu-kurulum.pkg",
	// The update check on every installed server asks for these two; they
	// are pinned to the rolling release by tag, immune to the Latest flag.
	"/sunucu/manifest.json":     base + "/download/indir/manifest.json",
	"/sunucu/manifest.json.sig": base + "/download/indir/manifest.json.sig",
	"/tanistirici/linux":        base + "/latest/download/zerbazaar-introducer_amd64.deb",
	"/tanistirici/windows":      base + "/latest/download/zerbazaar-introducer_windows.zip",
	"/tanistirici/macos":        base + "/latest/download/zerbazaar-introducer_macos.tar.gz",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if target, ok := routes[r.URL.Path]; ok {
			http.Redirect(w, r, target, http.StatusFound)
			return
		}
		http.Redirect(w, r, base, http.StatusFound)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	log.Println("indir redirect listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
