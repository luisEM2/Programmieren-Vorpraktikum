// immer dazuschreiben, wird später thematisiert
package main

// ermöglicht ausgabe auf die Konsole
import "fmt"

// Aufgabe: schau dir das programm an und versuch es zu verstehen
// Führe das programm aus und schau wie sich das ergebnis mit unterschiedlichen werten für Zahl ändert

// bevor du das programm ausführst, musst du geschriebene änderungen mit strg + s speichern
// Zum ausführen vom code musst du oben new terminal auswählen
// cd dein ordner (in aktuelles verzeichnis wechseln)
// go run main.go

// in diese Funktion (func) muss der gesamte code geschrieben werden, damit er ausgeführt wird
func main() {
	// variable Zahl definieren
	var zahl int = 2
	// zahl%2 ergibt den "rest" wenn man wie in der grundschule durch 2 teilt
	// wenn der rest == 1 ist, ist die Zahl ungerade
	// wenn der rest == 0 ist, ist die Zahl gerade
	// Im 2. Schritt wird geprüft, ob rest == 0 ist
	if zahl%2 == 0 {
		// "Zahl durch 2 teilbar ausgeben"
		fmt.Println("Zahl durch 2 teilbar")
		// else statement
	} else {
		// Falls if statement falsch ist, wird "Zahl nicht durch 2 teilbar" ausgegeben
		fmt.Println("Zahl nicht durch 2 teilbar")
	}
}
