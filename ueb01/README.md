# ueb01

- Programmiersprache: Go
- Netzwerkprotokoll: TCP
- Nachrichtenformat: Protocol Buffers (protobuf)
- Programmdateien: [mynode](./cmd/mynode) (A1, A2, A4) und [graphgen](./cmd/graphgen) (A3)

## Idee

### A1/A2

Jeder Knoten eines Netzwerks führt Buch über die empfangenen/gesendeten Anwendungsnachrichten. Empfängt ein Knoten eine Nachricht und hat bereits von allen anderen Knoten eine Nachricht empfangen sowie eine Nachricht an alle anderen Knoten gesendet, weiß dieser, dass es "von vorne losgeht". Dementsprechend setzt der Knoten seine Tabelle der empfangenen/gesendeten Nachrichten zurück und beginnt erneut damit, Nachrichten zu senden.

### A3

Zur Generierung eines Graphen wird eine Matrix der Größe n\*n erstellt. Diese stellt alle verfügbaren Kanten dar und hat initial bereits die Diagonale markiert, da Schleifen nicht zulässig sind. Zu Beginn wird jeder Knoten eingefügt und entsprechend dem Algorithmus der Aufgabenstellung mit dem restlichen Teilgraph per Zufall verbunden (verbinde je mit einem beliebigen Vorgängerknoten). Wenn dem Graph eine Kante hinzugefügt wird, wird dies in der Matrix registriert. Da es sich um ungerichtete Kanten handelt, werden dabei immer zwei Einträge gemacht. Für die verbleibenden m-n Kanten wird sich zufällig der Matrix bedient, sofern diese überhaupt noch Kanten hergibt (maximal n\*(n-1)/2, wird bei Programmstart geprüft).

### A4

Die Gerüchte verhalten sich anders als dich Anwendungsnachrichten aus A1 und A2. Trotzdem wird die gleiche Programmdatei verwendet. Die Unterscheidung findet anhand des Nachrichtenformats statt. Das Prinzip ist dabei jedoch ähnlich, es wird nämlich in einer Tabelle Buch darüber geführt, wie oft welches Gerücht bereits empfangen wurde. Auf das Zählen von gesendeten Gerüchten wird hierbei verzichtet. Hört ein Knoten mehrfach von einem Gerücht, so gibt dieser die Anzahl aus. Gibt ein Knoten keine Auskunft darüber, wie oft er ein Gerücht empfangen hat, so hat dieser das Gerücht nur ein einziges Mal gehört. Das ist dadurch sichergestellt, dass der Graph zusammenhängend ist.

## Nachrichtenformat

Zu Beginn der Übung wurde noch freier Text verwendet. Das stellte sich jedoch schnell als großer Nachteil heraus, da Nachrichten lediglich über einen Prefix unterschieden wurden und mehre Informationen in einem String schwierig aufzulösen waren. Deswegen dient nun Google's [Protocol Buffers](https://developers.google.com/protocol-buffers) als Datenformat. Alle Nachrichtendefinitionen befinden sich in einem zentralen Ordner (siehe [Softwarestruktur](#softwarestruktur)) und haben die Dateiendung `.proto`. Neben diesen Dateien sind auch die mit `protoc` kompilierten Nachrichtendefinitionen mit der Dateiendung `.pb.go` Teil der Abgabe. Anders als bei Nachrichtenformaten wie JSON oder XML sind Nachrichten in protobuf nicht für Menschen lesbar und lediglich eine Folge von Bytes.

Für die ersten zwei Aufgaben gibt es Nachrichtendefinitionen für Kontrollnachrichten und Anwendungsnachrichten und für die letzte Aufgabe eine für Gerüchte:

![Diagramm mit den verschiedenen Nachrichtendefinitionen](./docs/messages.svg)

## Softwarestruktur

Soweit wie Möglich wird das "[Standard Go Project Layout](https://github.com/golang-standards/project-layout)" eingehalten. Das heißt konkret:

- `/api` enthält die Nachrichtendefinitionen im Protocol Buffers Format
- `/cmd` enthält den ausführbaren Programmcode mit main-Funktionen
- `/configs` enthält die Konfigurationsdateien inklusive Graphivz Dateien
- `/docs` enthält Dokumentation und Messergebnisse
- `/internal` enthält den wesentlichen Programmcode
- `/scripts` enthält Shellskripte zum Ausführen der verschiedenen Aufgaben

Die main-Funktion des lokalen Knotens ist möglichst kompakt gehalten, sodass sich die wesentliche Anwendungslogik in `/internal` befindet. Hierbei wurden für verschiedene Aufgabenbereiche passende Packages erstellt. Neben einigen "util" Packages gibt es noch folgende Packages:

- `config` enthält Structs und Funktionen für den Umgang mit der eingelesenen Konfigurationsdatei, dazugehörigen Kommandozeilenparametern sowie Graphiz Dateien (Bestimmung von Nachbarn etc.)
- `directory` enthält die zu Beginn erwähnte Logik zum Tracken von Anwendungsnachrichten und Gerüchten
- `handler` enthält den Code, der die Anfragen an einen Knoten behandelt (Unterscheiden der Nachrichtentypen, Senden von Nachrichten, etc.)

Die Schritte 1-9 der Aufgabenstellung sind analog im Programmcode kommentiert, sodass ein Nachvollziehen hoffentlich einfach ist. Im Falle der `graphgen` Skripts ist ebenfalls jeder wesentliche Schritt der Anwendung kommentiert und beschrieben.

## Stärken & Schwächen

### Performance und Nebenläufigkeit

Go hat sich in diesem Fall als äußerst gut geeignete Programmiersprache herausgestellt. Im Vergleich zu Java hat sie den Vorteil, dass sie extrem leichtgewichtig ist (keine JVM) und im Vergleich zu C hat sie den den Vorteil, dass sie etwas weniger low-level ist aber dafür um sinnvolle Konzepte wie etwa Garbage Collection oder Goroutinen erweitert wurde. Gerade bei der Implementierung des TCP-Servers hat sich das bemerkt gemacht. So kann jeder Knoten ohne den Thread zu blocken eine Anfrage bearbeiten und währenddessen eine weitere annehmen. Zusätzlich enthält die Standardbibliothek bereits passende Konstrukte zur Implementierung von wechselseitigem Ausschuss. Auch das Generieren von sehr großen Beispielgraphen hat sich als kein Problem herausgestellt.

### Distribution makes things worse

Ein Problem, das die Implementierung sicher hat, ist, dass jeder Knoten die Information darüber, von wem empfangen und an wen gesendet wurde, selbst hält. Fällt nämlich ein Knoten aus, kann es passieren, dass ein Zurücksetzen der Tabelle mit diesen Informationen bei den anderen Knoten nicht mehr funktioniert. D.h. dass man mit einer Kontrollnachricht nicht von vorne beginnen könnte. Das wurde aber bewusst in Kauf genommen, da man sonst eine Reset-Nachricht propagieren müsste, was wiederum das Nachrichtenaufkommen mit steigendem n und m in die Höhe schießen lässt. Alternativ könnte man nicht erreichbare Knoten wieder aus der Tabelle entfernen, dann wären sie aber aus der Konfiguration verschwunden, auch wenn sie wieder erreichbar wären.

### Graceful shutdown

Eine Herausforderung war es, alle Knoten im Netzwerk zu beenden. Der Einsatz von Goroutinen hat sich nämlich nicht nur als Vorteil erwiesen, sondern kommt aufgrund der Nebenläufigkeit auch mit dazugewonnener Komplexität daher. Einfachheitshalber wurde deswegen ein "close" statt einem "graceful shutdown" implementiert, bei dem ein Knoten beim Propagieren einer solchen Kontrollnachricht unterbrochen werden kann, weil die Gegenstelle die Anfrage nicht erst fertig bearbeitet, sondern einfach schließt. Für die Implementierung wurden Channels verwendet, welche es ermöglichen, Signale zwischen den Goroutinen eines einzelnen Knotens zu senden, sodass es bei simultanem Eingang einer entsprechenden Kontrollnachricht nicht zu einem Fehler kommt.

## Beispiele

```sh
# a2.sh analog
$ ./scripts/a1.sh 4
[node-001] 13:06:13.566609 Listening on port :5000
[node-003] 13:06:13.566609 Listening on port :5002
[node-002] 13:06:13.566609 Listening on port :5001
[node-001] 13:06:13.567609 Neighbors: 2, 4, 3
[node-003] 13:06:13.567609 Neighbors: 4, 2, 1
[node-002] 13:06:13.567609 Neighbors: 4, 1, 3
[node-004] 13:06:13.567609 Listening on port :5003
[node-004] 13:06:13.567609 Neighbors: 2, 1, 3
[node-001] 13:06:38.298816 Received control message: START
[node-001] 13:06:38.301817 Sent message to node 2
[node-002] 13:06:38.301817 Received application message: 1
[node-001] 13:06:38.302818 Sent message to node 4
[node-004] 13:06:38.302818 Received application message: 1
[node-001] 13:06:38.302818 Sent message to node 3
[node-003] 13:06:38.302818 Received application message: 1
[node-002] 13:06:38.304816 Sent message to node 4
[node-004] 13:06:38.304816 Received application message: 2
[node-004] 13:06:38.305817 Sent message to node 2
[node-002] 13:06:38.305817 Received application message: 4
[node-001] 13:06:38.305817 Received application message: 2
[node-002] 13:06:38.305817 Sent message to node 1
[node-004] 13:06:38.305817 Received application message: 3
[node-003] 13:06:38.305817 Sent message to node 4
[node-004] 13:06:38.305817 Sent message to node 1
[node-001] 13:06:38.305817 Received application message: 4
[node-002] 13:06:38.305817 Sent message to node 3
[node-003] 13:06:38.305817 Received application message: 2
[node-004] 13:06:38.306818 Sent message to node 3
[node-003] 13:06:38.306818 Sent message to node 2
[node-002] 13:06:38.306818 Received application message: 3
[node-003] 13:06:38.306818 Received application message: 4
[node-001] 13:06:38.306818 Received application message: 3
[node-003] 13:06:38.306818 Sent message to node 1

# In einem zweiten Terminal, verursacht die erste Nachricht um 13:06:38
$ printf 'control_message:{command:START}' |
protoc.exe --encode=ueb01.Message Message.proto |
nc localhost 5000
```

```sh
# a3.sh pipet das Ergebnis in Dateien
$ graphgen -n 5 -m 6
graph  {
        2--1;
        3--1;
        4--1;
        5--2;
        2--4;
        1--5;
        1;
        2;
        3;
        4;
        5;

}
```

```sh
$ ./scripts/a4.sh 08
[node-001] 14:41:12.887723 Listening on port :5000
[node-002] 14:41:12.887723 Listening on port :5001
[node-001] 14:41:12.887723 Neighbors: 2, 3, 5, 4
[node-002] 14:41:12.887723 Neighbors: 1, 4, 5
[node-005] 14:41:12.889721 Listening on port :5004
[node-003] 14:41:12.889721 Listening on port :5002
[node-005] 14:41:12.889721 Neighbors: 4, 2, 1
[node-003] 14:41:12.890722 Neighbors: 1
[node-004] 14:41:12.889721 Listening on port :5003
[node-004] 14:41:12.890722 Neighbors: 2, 5, 1
[node-002] 14:42:05.976412 Received rumor: HTW > UdS
[node-002] 14:42:05.979412 Told node 1 about rumor 'HTW > UdS'
[node-001] 14:42:05.979412 Received rumor: HTW > UdS
[node-002] 14:42:05.979412 Told node 4 about rumor 'HTW > UdS'
[node-004] 14:42:05.980410 Received rumor: HTW > UdS
[node-002] 14:42:05.980410 Told node 5 about rumor 'HTW > UdS'
[node-005] 14:42:05.980410 Received rumor: HTW > UdS
[node-001] 14:42:05.982414 Told node 3 about rumor 'HTW > UdS'
[node-003] 14:42:05.982414 Received rumor: HTW > UdS
[node-005] 14:42:05.983413 Received rumor: HTW > UdS
[node-004] 14:42:05.983413 Told node 5 about rumor 'HTW > UdS'
[node-001] 14:42:05.983413 Told node 5 about rumor 'HTW > UdS'
[node-005] 14:42:05.983413 Received rumor: HTW > UdS
[node-004] 14:42:05.983413 Received rumor: HTW > UdS
[node-005] 14:42:05.983413 Told node 4 about rumor 'HTW > UdS'
[node-001] 14:42:05.983413 Received rumor: HTW > UdS
[node-004] 14:42:05.983413 Told node 1 about rumor 'HTW > UdS'
[node-001] 14:42:05.983413 Told node 4 about rumor 'HTW > UdS'
[node-004] 14:42:05.983413 Received rumor: HTW > UdS
[node-005] 14:42:05.983413 Told node 1 about rumor 'HTW > UdS'
[node-001] 14:42:05.983413 Heard about rumor 'HTW > UdS' 2 times
[node-004] 14:42:05.983413 Heard about rumor 'HTW > UdS' 2 times
[node-005] 14:42:05.984411 Heard about rumor 'HTW > UdS' 2 times
[node-001] 14:42:05.983413 Received rumor: HTW > UdS
[node-004] 14:42:05.984411 Heard about rumor 'HTW > UdS' 3 times
[node-005] 14:42:05.984411 Heard about rumor 'HTW > UdS' 3 times
[node-001] 14:42:05.984411 Heard about rumor 'HTW > UdS' 3 times

# In einem zweiten Terminal, verursacht die erste Nachricht um 14:42:05
$ printf 'rumor:{content:"HTW > UdS"}' |
protoc.exe --encode=ueb01.Message Message.proto |
nc localhost 5000
```

## Fazit

Die Übung war eine gute Gelegenheit sich Go und Protocol Buffers das erste Mal anzuschauen. Es gab viele Probleme zu lösen, wie kritische Bereiche des Request-Handlings durch Mutex zu schützen oder das Beenden aller Knoten zu koordinieren, ohne dass Knoten abstürzten.
