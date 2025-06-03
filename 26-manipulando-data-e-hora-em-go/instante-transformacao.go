package main

import (
    "fmt"
    "time"
)

func main() {
	agoraLocal := time.Now()
	agoraUTC := agoraLocal.UTC()
	locSP, _ := time.LoadLocation("America/Sao_Paulo")
	agoraSP := agoraUTC.In(locSP)

	fmt.Println("Now (local):", agoraLocal)
	fmt.Println("Now (UTC):", agoraUTC)
	fmt.Println("Now (Sao Paulo):", agoraSP)
}