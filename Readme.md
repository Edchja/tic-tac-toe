# X-X-O

Erstelle ein XXO Spiel in der Kommandozeile.

Das Spielbrett wird permanent folgendermaßen angezeigt:

|   |   |   |
–––––––––––––
|   |   |   |
–––––––––––––
|   |   |   |

Der Spieler hat das X und startet als erster.
Er kann über den Nummernblock ein Feld auswählen, das er markieren möchte.
Das markierte Feld wird schließlich entsprechend angezeigt.

3 Zustände: 0, 1 (player), 2 (computer)

Beispiel:

Spieler gibt ein: Num 7

| X |   |   |
–––––––––––––
|   |   |   |
–––––––––––––
|   |   |   |

Als nächstes wählt der Computer zufällig ein freies Feld:

| X |   | O |
–––––––––––––
|   |   |   |
–––––––––––––
|   |   |   |

Felder können nie doppelt gewählt werden. In dem Fall wird eine Meldung angezeigt und die Eingabe wiederholt.

Hat einer der Spieler zuerst 3 in einer Reihe (auch diagonal), endet das Spiel mit einer kleinen Animation, die dem Gewinner gratuliert.

Vorgaben:

Das Spielbrett wird intern als Mehrdimensionales Array abgebildet.

Bonus:

Der Computer verhält sich intelligent!


Zum wiederholten anzeigen des Spielbretts können die Funktionen aus der vorherigen Übung verwendet werden:

// Bewegt den cursor eine Zeile hoch.
func up() {
	fmt.Print("\033[A")
}

// Bewegt den cursor eine Zeile runter.
func down() {
	fmt.Print("\033[B")
}

// Löscht die aktuelle Zeile.
func clearLine() {
	fmt.Print("\033[G\033[K")
}
