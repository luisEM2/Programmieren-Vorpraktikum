package structs

// Struct zur Darstellung eines Bankkontos
type BankAccount struct {
	Owner   string  // Kontoinhaber
	Number  string  // Kontonummer
	Balance float64 // Kontostand
}

// Deposit erhöht den Kontostand um den angegebenen Betrag
func (b *BankAccount) Deposit(amount float64) {
	// TODO: Betrag zum Kontostand addieren
}

// Withdraw verringert den Kontostand um den angegebenen Betrag,
// wenn genug Guthaben vorhanden ist. Gibt true zurück, wenn erfolgreich,
// sonst false.
func (b *BankAccount) Withdraw(amount float64) bool {
	// TODO: Nur abheben, wenn genug Guthaben
	return false // Platzhalter
}
