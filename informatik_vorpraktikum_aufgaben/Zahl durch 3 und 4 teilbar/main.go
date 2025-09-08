package main

import "fmt"

// schreibe den code Zahl durch zwei teilbar so um, dass die funktion prüft, ob eine Zahl durch 3 und / oder durch 4 teilbar ist.
// Zusatz: Schaffst du es, die funktion so zu schreiben, dass sie bei der zahl 12 nur zahl durch 3 teilbar ausgibt?
// Zusatz: Schaffst du es, die funktion so zu schreiben, dass sie bei der zahl 12 nur zahl durch 4 teilbar ausgibt?
// Zusatz: Schaffst du es, die funktion so zu schreiben, dass sie bei der zahl 12 zahl durch 3 teilbar ausgibt und Zahl durch 4 teilbar ausgibt?

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
