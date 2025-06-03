package main

import (
    "fmt"
    "time"
)

func main() {
	t := time.Now()
	maisUmaHora := t.Add(time.Hour * 1)
	menosTrintaMin := t.Add(-30 * time.Minute)
	fmt.Println("Agora +1h:", maisUmaHora)
	fmt.Println("Agora -30m:", menosTrintaMin)
}
