# IS-105-1-22V-Datakommunikasjon-og-operativsystem
Kode for IS-105-1 22V Datakommunikasjon og operativsystem
## [Module 1](https://github.com/Eiriksb/IS-105-1-22V-Datakommunikasjon-og-operativsystem/tree/main/Module1)
- Enkelt program som kjører quotes.go fra main.go med rsc.io/quote. Printer ut quotes men returnerer ingen ting.

## [Rivercrossing](https://github.com/Eiriksb/IS-105-1-22V-Datakommunikasjon-og-operativsystem/tree/main/Rivercrossing)
- Veldig likt et rivercrossing spill men ikke helt da den ikke egentlig lagrer noe.
- Når endringer er gjort får man en melding tilbake som bekrefter kommandoen og man kan se ved **info** kommandoen at det er blitt endringer i info panelet.
- Spillet reseter seg til standard verdier når du går ut av spillet med Ctrl+C 
- Spiller kan bruke kommandoer som: 

    | Command | Description |
    | --- | --- |
    | info | for å får å kjøre en StateView funksjon som kommer tilbake med info om hvor ting er (Båt, Land Vest eller Land Øst) |
    | Put | Flytter enhet til destinasjon. Feks. **PutKyB** Som Flytter Kylling til Båt |
    
    | Enheter | Description | Destinasjon | Description |
    | --- | --- | --- | --- |
    | Ky | Kylling | B | Flytter enhet til Båt |
    | Rev | Rev | V | Flytter enhet til Land Vest |
    | Korn | Korn | E | Flytter enhet til Land Øst |
    | Hs | Menneske | --- | --- |
    | Boat | Båt | --- | --- |
