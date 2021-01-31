# ueb03

- Programmiersprache: Go
- Netzwerkprotokoll: TCP
- Nachrichtenformat: Protocol Buffers (protobuf)
- Programmdateien: [baccount](cmd/baccount) und graphgen (siehe ueb01)

## Idee

### Election

Analog zu ueb02. Die Election ist mittels *Echo-based Election* realisiert mit dem Unterschied, dass sie bei Start automatisch stattfindet und nicht per Kontrollnachricht gestartet wird. Der gewählte Koordinator versendet nach Feststellen des eigenen Wahlsiegs Startnachrichten an alle anderen Knoten, die daraufhin Schritte 1-11 durchlaufen.

### Mutual Exclusion

Für den wechselseitigen Ausschluss wird der Algorithmus von Ricart & Agrawala verwendet, da dieser im Vergleich zu dem Lamport Algorithmus Nachrichten einspart. Die Nachrichtenkomplexität ist aufgrund von Flooding nämlich ohnehin schlecht. Für den Algorithmus kommt eine Thread-sichere Lamport Clock, sowie eine Queue in Form einer doppelt verketteten Liste zum Einsatz.

Die Implementierung weicht dabei von der aus der Vorlesung bekannten etwas ab, sodass in einer Abfrage geprüft wird, ob es notwendig ist, einen Request zu queuen (vgl. [Distributed Mutual Exclusion, S.36](http://www2.imm.dtu.dk/courses/02222/Spring_2011/W9L2/Chapter_12a.pdf)).

### Snapshots

Hier gab es keine Wahl zwischen mehreren Algorithmen, weswegen logischerweise der Chandy-Lamport-Algorithmus verwendet wird. Dabei werden nicht nur Kontostände aufgezeichnet, sondern auch die einzelnen Veränderungen der Kontostände. Die Implementierung ist analog zu der in der Vorlesung vorgestellten Vorgehensweise in Pseudocode:

```go
if h.dir.Snapshot.IsFirstMarker() {
    h.dir.Snapshot.RecordState(h.conf.Params.Balance)
    h.dir.Snapshot.MarkSenderAsEmpty(sender)
    h.sentSnapshotMarkerToNeighbors()
    h.dir.Snapshot.StartRecording(sender, h.conf.FindAllNeighbors())
} else {
    h.dir.Snapshot.StopRecording(sender)
}
```

Der Koordinator wiederholt die Snapshots alle 5 Sekunden und gibt eine entsprechende Warnung bei sich ändernden Summen aus. Diese werden anhand des Kontostands erkannt, könnten aber auch anhand der aufgezeichneten Kontostand-Änderungen ermittelt werden (etwa dann, wenn ein Betrag x addiert aber ein anderer Betrag y subtrahiert wird oder zwei Mal Addiert wird statt ein Mal Subtrahiert).

## Nachrichtenformat

Analog zu ueb02. Hier ist es bei Protocol Buffers geblieben und die Nachrichten wurden um entsprechende neue Nachrichten erweitert, z.B. Nachrichten für Mutex oder Snapshots.

## Softwarestruktur

Analog zu ueb02. Auch die Softwarestruktur hat sich kaum verändert. Neben den neuen notwendigen Modulen für die `directory` und `handler` Pakete ist im `util` Paket ein neues Paket für benötigte Datenstrukturen hinzugekommen. Das `collection` Paket enthält die Strukturen für die Lamport Clock und Queue, die zur Realisierung des wechselseitigen Ausschluss verwendet werden.

## Stärken & Schwächen

Die bereits in vorherigen Übungen öfter angesprochenen Punkte "Performance" und "Resilienz" werden an dieser Stelle nicht wiederholt, gelten aber auch für diese Übung. Dazu kommen folgende Punkte:

### Korrektheit

Die Korrektheit der Anwendung ist nicht vollständig gewährleistet, da nur einseitig gelockt wird, statt beidseitig wie in Schritt 3 beschrieben. Dementsprechend kann es bei folgender Situation zu einem veränderten globalen Kontostand kommen:

1 lockt 2 und 2 lockt 3. Dann kann es passieren, dass 1 mit einem alten Kontostand von 2 rechnet und entsprechend mit zwei verschiedenen Werten x und y addiert/subtrahiert wird.

Außerdem kann es zu einem Problem kommen, wenn Zeitstempel identisch sind. Hier könnte man zur ID zurückfallen und diese vergleichen, jedoch wurde das für die Abgabe wieder ausgebaut, da dies Deadlocks verursachen kann. Entsprechend ist Ricart-Agrawala wie im Lehrbuch implementiert, ohne auf diese Dinge Acht zu nehmen.

### Nachrichtenkomplexität

Dadurch, dass Flooding verwendet wird, werden Unmengen an Nachrichten versendet, obwohl diese eingespart werden könnten. Bei der Election wird nämlich ein minimaler Spannbaum aufgespannt, welcher zur Kommunikation verwendet werden könnte. Darauf wurde aber verzichtet, da in der Aufgabenstellung explizit von Flooding die Rede war. Im Log der Anwendung werden jedoch Nachrichten, die das Flooding dokumentieren, zwecks Einfachheit bewusst weggelassen. Nachrichten, die mit Flooding versendet wurden, sind mit einem Asterisk (*) gekennzeichnet.

## Beispiele

```sh
$ ./scripts/a1.sh 5
n,m
5,9
[account-003] 17:36:29.715849 Listening on port :5002
[account-001] 17:36:29.715849 Listening on port :5000
[account-005] 17:36:29.715849 Listening on port :5004
[account-004] 17:36:29.715849 Listening on port :5003
[account-003] 17:36:29.716849 Neighbors: 1, 2, 5
[account-002] 17:36:29.715849 Listening on port :5001
[account-001] 17:36:29.716849 Neighbors: 2, 3, 4, 5
[account-005] 17:36:29.716849 Neighbors: 2, 4, 1, 3
[account-004] 17:36:29.716849 Neighbors: 1, 5, 2
[account-003] 17:36:29.716849 Balance: 27000
[account-002] 17:36:29.716849 Neighbors: 1, 5, 3, 4
[account-001] 17:36:29.716849 Balance: 17000
[account-005] 17:36:29.716849 Balance: 95000
[account-004] 17:36:29.716849 Balance: 32000
[account-002] 17:36:29.716849 Balance: 52000
[account-005] 17:36:29.723849 Sent explorer to node 2
[account-005] 17:36:29.723849 Sent explorer to node 4
[account-002] 17:36:29.723849 Received explorer message with ID 5 from 5
[account-004] 17:36:29.723849 Received explorer message with ID 5 from 5
[account-005] 17:36:29.723849 Sent explorer to node 1
[account-001] 17:36:29.724851 Received explorer message with ID 5 from 5
[account-005] 17:36:29.724851 Sent explorer to node 3
[account-003] 17:36:29.724851 Received explorer message with ID 5 from 5
[account-004] 17:36:29.726850 Sent explorer to node 1
[account-002] 17:36:29.727851 Sent explorer to node 1
[account-001] 17:36:29.727851 Received explorer message with ID 5 from 4
[account-002] 17:36:29.727851 Received explorer message with ID 5 from 1
[account-001] 17:36:29.727851 Sent explorer to node 2
[account-003] 17:36:29.727851 Sent explorer to node 1
[account-001] 17:36:29.727851 Received explorer message with ID 5 from 2
[account-004] 17:36:29.727851 Sent explorer to node 2
[account-001] 17:36:29.727851 Received explorer message with ID 5 from 3
[account-002] 17:36:29.727851 Received explorer message with ID 5 from 4
[account-003] 17:36:29.727851 Received explorer message with ID 5 from 2
[account-002] 17:36:29.727851 Sent explorer to node 3
[account-003] 17:36:29.727851 Received explorer message with ID 5 from 1
[account-001] 17:36:29.727851 Sent explorer to node 3
[account-002] 17:36:29.727851 Received explorer message with ID 5 from 3
[account-003] 17:36:29.727851 Sent explorer to node 2
[account-004] 17:36:29.728852 Received explorer message with ID 5 from 2
[account-002] 17:36:29.728852 Sent explorer to node 4
[account-001] 17:36:29.728852 Sent explorer to node 4
[account-004] 17:36:29.728852 Received explorer message with ID 5 from 1
[account-003] 17:36:29.728852 Sent echo to node 5
[account-005] 17:36:29.728852 Received echo message with ID 5 from 3
[account-002] 17:36:29.728852 Sent echo to node 5
[account-005] 17:36:29.728852 Received echo message with ID 5 from 2
[account-005] 17:36:29.728852 Received echo message with ID 5 from 4
[account-004] 17:36:29.728852 Sent echo to node 5
[account-001] 17:36:29.728852 Sent echo to node 5
[account-005] 17:36:29.728852 Received echo message with ID 5 from 1
[account-005] 17:36:29.728852 INITIATOR IS GREEN
[account-005] 17:36:30.728862 ------- ELECTION VICTORY -------
[account-002] 17:36:30.738390 Received control message: START
[account-004] 17:36:30.738390 Received control message: START
[account-001] 17:36:30.739390 Received control message: START
[account-001] 17:36:30.739390 Received control message: START
[account-001] 17:36:30.739390 Received control message: START
[account-003] 17:36:30.739390 Received control message: START
[account-005] 17:36:30.739390 Received control message: START
[account-005] 17:36:30.739390 Received control message: START
[account-002] 17:36:30.739390 Received control message: START
[account-003] 17:36:30.740391 Received control message: START
[account-004] 17:36:30.740391 Starting...
[account-002] 17:36:30.740391 Received control message: START
[account-001] 17:36:30.740391 Received control message: START
[account-003] 17:36:30.740391 Received control message: START
[account-002] 17:36:30.740391 Received control message: START
[account-004] 17:36:30.740391 Received control message: START
[account-002] 17:36:30.740391 Starting...
[account-004] 17:36:30.740391 Received control message: START
[account-002] 17:36:30.740391 Received control message: START
[account-003] 17:36:30.740391 Starting...
[account-005] 17:36:30.741392 Received control message: START
[account-004] 17:36:30.740391 Received control message: START
[account-001] 17:36:30.741392 Received control message: START
[account-005] 17:36:30.741392 Received control message: START
[account-003] 17:36:30.741392 Received control message: START
[account-001] 17:36:30.741392 Starting...
[account-005] 17:36:30.741392 Starting...
[account-001] 17:36:30.741392 Broadcasting mutex request '1-1' with resource = 5, timestamp = 1
[account-002] 17:36:30.742391 Received mutex request from node 1
[account-003] 17:36:30.742391 Received mutex request from node 1
[account-004] 17:36:30.743390 Received mutex request from node 1
[account-005] 17:36:30.743390 Received mutex request from node 1
[account-001] 17:36:30.744391 Received mutex response from node 2
[account-002] 17:36:30.744391 Sent mutex response to node 1
[account-003] 17:36:30.744391 Sent mutex response to node 1
[account-001] 17:36:30.744391 Received mutex response from node 3
[account-001] 17:36:30.745392 Received mutex response from node 4
[account-004] 17:36:30.745392 Sent mutex response to node 1
[account-001] 17:36:30.745392 Received mutex response from node 5
[account-005] 17:36:30.745392 Sent mutex response to node 1
[account-001] 17:36:30.745392 --- LOCKING RESOURCE 5 ---
[account-001] 17:36:30.746391 Sent balance request to node 5
[account-005] 17:36:30.746391 Received balance request from node 1
[account-005] 17:36:30.747390 Sent balance response to node 1: B = 95000
[account-001] 17:36:30.747390 Received balance response from node 5
[account-005] 17:36:30.747390 Received application message from node 1
[account-001] 17:36:30.747390 Sent application message to node 5: B = 17000, p = 40
[account-005] 17:36:30.747390 Decreasing balance by 38000: Old = 95000, New = 57000 (8B)
[account-001] 17:36:30.748390 Increasing balance by 38000: Old = 17000, New = 55000 (7A)
[account-001] 17:36:30.748390 Received acknowledgment from node 5
[account-005] 17:36:30.748390 Sent acknowledgment to node 1
[account-001] 17:36:30.748390 --- UNLOCKING RESOURCE 5 ---
[account-002] 17:36:31.746738 Broadcasting mutex request '2-2' with resource = 4, timestamp = 3
[account-003] 17:36:31.746738 Broadcasting mutex request '3-2' with resource = 2, timestamp = 3
# ...
5 20520
2 34676
1 17410
3 147185
4 3209
Previous system balance = 0
Current system balance = 223000
# ...
```

## Fazit

Eine sehr zeitaufwendige und auch frustrierende Übung, da man erst einmal herausfinden muss, wieso sich die globale Geldmenge überhaupt verändert und anschließend beim Versuch, diesen "Bug" zu beheben, Deadlocks verursacht.
