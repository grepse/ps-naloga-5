// program ps-naloga-5 vodi vremenske postaje in izpisuje poročila
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grepse/ps-naloga-5/vreme"
	"github.com/urfave/cli/v3"
)

func zazeni() {
	vreme.DodajPostajo("KP01", "Luka Koper", "Koper")
	vreme.DodajPostajo("KR01", "Kredarica", "Triglav")
	vreme.DodajPostajo("LJ01", "Ljubljana Bezigrad", "Ljubljana")
	vreme.DodajPostajo("MB01", "Maribor Tabor", "Maribor")
	vreme.DodajPostajo("CE01", "Celje Medlog", "Celje")
	vreme.DodajPostajo("NM01", "Novo mesto", "Novo Mesto")
	vreme.DodajPostajo("MS01", "Murska Sobota Rakican", "Murska Sobota")
	vreme.DodajPostajo("PO01", "Postojna", "Postojna")
	vreme.DodajPostajo("RA01", "Ratece", "Kranjska Gora")
	vreme.DodajPostajo("NG01", "Nova Gorica Bilje", "Nova Gorica")
	vreme.DodajPostajo("KK01", "Kocevje", "Kocevje")

	/*
		var kp01 Postaja
		kp01.ime = "Koper Luka"
		kp01.kraj = "Koper"
		kp01.meritve = []float64{24.5, 27, 26.1, 28.4}
		slovar["KP01"] = kp01

		var kr01 Postaja
		kr01.ime = "Kredarica"
		kr01.kraj = "Triglav"
		kr01.meritve = []float64{2, -1.5, 3.5}
		slovar["KR01"] = kr01

		var lj01 Postaja
		lj01.ime = "Ljubljana Bežigrad"
		lj01.kraj = "Ljubljana"
		lj01.meritve = []float64{18.2, 21.5, 19.8, 16.5}
		slovar["LJ01"] = lj01

		var mb01 Postaja
		mb01.ime = "Maribor Tabor"
		mb01.kraj = "Maribor"
		mb01.meritve = []float64{12.4}
		slovar["MB01"] = mb01

		var ce01 Postaja
		ce01.ime = "Celje Medlog"
		ce01.kraj = "Celje"
		ce01.meritve = []float64{8, 10.5, 9.1}
		slovar["CE01"] = ce01

		var nm01 Postaja
		nm01.ime = "Novo mesto"
		nm01.kraj = "Novo mesto"
		nm01.meritve = []float64{24, 25, 26}
		slovar["NM01"] = nm01

		var ms01 Postaja
		ms01.ime = "Murska Sobota Rakičan"
		ms01.kraj = "Murska Sobota"
		ms01.meritve = []float64{14, 15, 16}
		slovar["MS01"] = ms01

		var po01 Postaja
		po01.ime = "Postojna"
		po01.kraj = "Postojna"
		po01.meritve = []float64{4, 5, 6}
		slovar["PO01"] = po01

		var ra01 Postaja
		ra01.ime = "Rateče"
		ra01.kraj = "Kranjska Gora"
		ra01.meritve = []float64{-12.5, -8, -15.3, -10.2}
		slovar["RA01"] = ra01

		var ng01 Postaja
		ng01.ime = "Nova Gorica Bilje"
		ng01.kraj = "Nova Gorica"
		ng01.meritve = []float64{20.1, 22.3}
		slovar["NG01"] = ng01

		var kk01 Postaja
		kk01.ime = "Kočevje"
		kk01.kraj = "Kočevje"
		kk01.meritve = []float64{}
		slovar["KK01"] = kk01
	*/
	vreme.DodajMeritev("KP01", 24.5)
	vreme.DodajMeritev("KP01", 27)
	vreme.DodajMeritev("KP01", 26.1)
	vreme.DodajMeritev("KP01", 28.4)

	vreme.DodajMeritev("KR01", 2)
	vreme.DodajMeritev("KR01", -1.5)
	vreme.DodajMeritev("KR01", 3.5)

	vreme.DodajMeritev("LJ01", 18.2)
	vreme.DodajMeritev("LJ01", 21.5)
	vreme.DodajMeritev("LJ01", 19.8)
	vreme.DodajMeritev("LJ01", 16.5)

	vreme.DodajMeritev("MB01", 12.4)

	vreme.DodajMeritev("CE01", 8)
	vreme.DodajMeritev("CE01", 10.5)
	vreme.DodajMeritev("CE01", 9.1)

	vreme.DodajMeritev("NM01", 24)
	vreme.DodajMeritev("NM01", 25)
	vreme.DodajMeritev("NM01", 26)

	vreme.DodajMeritev("MS01", 14)
	vreme.DodajMeritev("MS01", 15)
	vreme.DodajMeritev("MS01", 16)

	vreme.DodajMeritev("PO01", 4)
	vreme.DodajMeritev("PO01", 5)
	vreme.DodajMeritev("PO01", 6)

	vreme.DodajMeritev("RA01", -12.5)
	vreme.DodajMeritev("RA01", -8)
	vreme.DodajMeritev("RA01", -15.3)
	vreme.DodajMeritev("RA01", -10.2)

	vreme.DodajMeritev("NG01", 20.1)
	vreme.DodajMeritev("NG01", 22.3)

	vreme.IzpisMeritev()
	fmt.Println()
	vreme.DodajMeritev("MB01", 14.0)
	vreme.DodajMeritev("MB01", 75.0)
	vreme.DodajMeritev("CE01", 10.0)
	vreme.DodajMeritev("XX01", 10.0)

	fmt.Println()

	fmt.Println()
	vreme.IzpisPorocila()

	fmt.Println()
}

func main() {
	cmd := &cli.Command{
		Name:  "vreme",
		Usage: "vodenje vremenskih postaj in izpis porocil",
		Flags: []cli.Flag{
			&cli.FloatFlag{Name: "minTemp", Value: -50, Usage: "najnizja dovoljena temperatura meritve"},
			&cli.FloatFlag{Name: "maxTemp", Value: 50, Usage: "najvisja dovoljena temperatura meritve"},
			&cli.IntFlag{Name: "stMeritev", Value: 3, Usage: "najmanjse dovoljeno stevilo meritev za statistiko"},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			vreme.Nastavi(cmd.Float("minTemp"), cmd.Float("maxTemp"), cmd.Int("stMeritev"))
			zazeni()
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
