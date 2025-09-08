package main

import "fmt"

// schreibe den code Zahl durch zwei teilbar so um, dass die funktion prüft, ob eine Zahl positiv, negativ oder null ist

// bevor du das programm ausführst, musst du geschriebene änderungen mit strg + s speichern
// Zum ausführen vom code musst du oben new terminal auswählen
// cd dein ordner (in aktuelles verzeichnis wechseln)
// go run main.go

// in diese Funktion (func) muss der gesamte code geschrieben werden, damit er ausgeführt wird
func main() {
	// variable Zahl definieren
	var zahl int = 2
	if zahl%2 == 0 {
		fmt.Println("Zahl durch 2 teilbar")
	} else {
		fmt.Println("Zahl nicht durch 2 teilbar")
	}
}
