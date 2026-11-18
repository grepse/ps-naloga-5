// Package vreme omogoca dodajanje vremenskih postaj, dodajenje meritev
// izracun statistike in izpis vremenskih porocil
package vreme

import (
	"fmt"
	"slices"
)

// Postaja je vremenska postaja z imenom, krejem in seznamom meritev
// ustarimo jo s funkcijo DodajPostajo
type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

var postaje = make(map[string]Postaja)

var (
	minTemp   = -50.0
	maxTemp   = 50.0
	stMeritev = 3
)

// Nastavi doloci min, max temeperaturo in pa najmanjse dovoljeno stevilo meritev
func Nastavi(min, max float64, st int) {
	minTemp = min
	maxTemp = max
	stMeritev = st
}

// DodajPostajo doda novo postajo s podano oznako, imenom in krajem
func DodajPostajo(oznaka string, ime string, kraj string) {
	postaje[oznaka] = Postaja{
		ime:  ime,
		kraj: kraj,
	}
}

// DodajMeritev doda novo meritev v seznam meritev podane postaje (preko oznake)
func DodajMeritev(oznaka string, temperatura float64) {
	// na to postajo doda novo meritev
	postaja, ok := postaje[oznaka]

	if !ok {
		fmt.Printf("Postaja %s ne obstaja\n", oznaka)
		return
	}

	if temperatura < minTemp || temperatura > maxTemp {
		fmt.Printf("Neveljavna meritev\n")
		return
	}

	postaja.meritve = append(postaja.meritve, temperatura)
	postaje[oznaka] = postaja
}

// IzpisMeritev izpise vse meritve in postaje kjer so te meritve zapisane
func IzpisMeritev() {
	st := len(postaje)

	if st < 1 {
		return
	}

	fmt.Printf("MERITVE:\n")

	for oznaka, p := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", oznaka, p.ime, p.kraj, p.meritve)
	}
}

func povprecje(m []float64) float64 {
	vsota := 0.0
	for _, x := range m {
		vsota += x
	}
	return vsota / float64(len(m))
}

func statistika(oznaka string) (float64, float64, float64, bool) {
	p, ok := postaje[oznaka]

	if !ok {
		return 0, 0, 0, false
	}

	t := p.meritve

	if len(t) < stMeritev {
		return 0, 0, 0, false
	}

	najnizja := slices.Min(t)
	najvisja := slices.Max(t)
	avg := povprecje(t)

	return najnizja, najvisja, avg, true
}

// IzpisPorocila za vsako postajo izpise mintemp, maxtemp in avg temp in pa oceno vremena
// za postaje s premalo meritvami napise opozoriloc
func IzpisPorocila() {
	st := len(postaje)
	if st < 1 {
		return
	}

	for oznaka, p := range postaje {
		najnizja, najvisja, avg, ok := statistika(oznaka)

		if !ok {
			fmt.Printf("%s: premalo meritev\n", p.ime)
			continue
		}

		fmt.Printf("%s: min %.2f °C, max %.2f °C, povprecje %.2f °C", p.ime, najnizja, najvisja, avg)

		switch {
		case avg >= 25:
			fmt.Printf(" -> Vroče\n")
		case avg >= 15:
			fmt.Printf(" -> Toplo\n")
		case avg >= 5:
			fmt.Printf(" -> Hladno\n")
		default:
			fmt.Printf(" -> Mrzlo\n")
		}
	}
}
