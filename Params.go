package main

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"sync"
)

type ParamsType struct {
	ProductID      string  // PID
	Volts          float64 // V
	Amps           float64 // I
	Power          float64 // P
	ConsumedEnergy float64 // CE
	StateOfCharge  float64 // SOC
	TimeToGo       float64 // TTG
	Alarm          string  // Alarm
	AlarmReason    string  // AR
	Model          string  // BMV
	Firmware       string  // FW
	MonitorMode    uint16  // MON
	//Checksum        s
	DeepestDischarge float64 // H1
	LastDischarge    float64 // H2
	AvgDischarge     float64 // H3
	Cycles           uint16  // H4
	FullDischarges   uint16  // H5
	CumulativeAHr    float64 // H6
	MinVolts         float64 // H7
	MaxVolts         float64 // H8
	SecondsSinceFull uint32  // H9
	NumAutoSync      uint32  // H10
	NumLowVoltAlarm  uint32  // H11
	NumHighVoltAlarm uint32  // H12
	MinAuxVolts      float64 // H15
	MaxAuxVolts      float64 // H16
	DischargedEnergy float64 // H17
	ChargedEnergy    float64 // H18
	//Checksum        s

	mu sync.Mutex
}

func DecodeProductID(id string) string {
	switch id {
	case "0x203":
		return "BMV-700"
		break
	case "0x204":
		return "BMV-702"
		break
	case "0x205":
		return "BMV-700H"
		break
	case "0x0300":
		return "BlueSolar MPPT 70|15"
		break
	case "0xA040":
		return "BlueSolar MPPT 75|50"
		break
	case "0xA041":
		return "BlueSolar MPPT 150|35"
		break
	case "0xA042":
		return "BlueSolar MPPT 75|15"
		break
	case "0xA043":
		return "BlueSolar MPPT 100|15"
		break
	case "0xA044":
		return "BlueSolar MPPT 100|30"
		break
	case "0xA045":
		return "BlueSolar MPPT 100|50"
		break
	case "0xA046":
		return "BlueSolar MPPT 150|70"
		break
	case "0xA047":
		return "BlueSolar MPPT 150|100"
		break
	case "0xA049":
		return "BlueSolar MPPT 100|50 rev2"
		break
	case "0xA04A":
		return "BlueSolar MPPT 100|30 rev2"
		break
	case "0xA04B":
		return "BlueSolar MPPT 150|35 rev2"
		break
	case "0xA04C":
		return "BlueSolar MPPT 75|10"
		break
	case "0xA04D":
		return "BlueSolar MPPT 150|45"
		break
	case "0xA04E":
		return "BlueSolar MPPT 150|60"
		break
	case "0xA04F":
		return "BlueSolar MPPT 150|85"
		break
	case "0xA050":
		return "SmartSolar MPPT 250|100"
		break
	case "0xA051":
		return "SmartSolar MPPT 150|100"
		break
	case "0xA052":
		return "SmartSolar MPPT 150|85"
		break
	case "0xA053":
		return "SmartSolar MPPT 75|15"
		break
	case "0xA054":
		return "SmartSolar MPPT 75|10"
		break
	case "0xA055":
		return "SmartSolar MPPT 100|15"
		break
	case "0xA056":
		return "SmartSolar MPPT 100|30"
		break
	case "0xA057":
		return "SmartSolar MPPT 100|50"
		break
	case "0xA058":
		return "SmartSolar MPPT 150|35"
		break
	case "0xA059":
		return "SmartSolar MPPT 150|100 rev2"
		break
	case "0xA05A":
		return "SmartSolar MPPT 150|85 rev2"
		break
	case "0xA05B":
		return "SmartSolar MPPT 250|70"
		break
	case "0xA05C":
		return "SmartSolar MPPT 250|85"
		break
	case "0xA05D":
		return "SmartSolar MPPT 250|60"
		break
	case "0xA05E":
		return "SmartSolar MPPT 250|45"
		break
	case "0xA05F":
		return "SmartSolar MPPT 100|20"
		break
	case "0xA060":
		return "SmartSolar MPPT 100|20"
		break
	case "0xA061":
		return "SmartSolar MPPT 150|45"
		break
	case "0xA062":
		return "SmartSolar MPPT 150|60"
		break
	case "0xA063":
		return "SmartSolar MPPT 150|70"
		break
	case "0xA064":
		return "SmartSolar MPPT 250|85 rev2"
		break
	case "0xA065":
		return "SmartSolar MPPT 250|100 rev2"
		break
	case "0xA066":
		return "BlueSolar MPPT 100|20"
		break
	case "0xA067":
		return "BlueSolar MPPT 100|20 48V"
		break
	case "0xA068":
		return "SmartSolar MPPT 250|60 rev2"
		break
	case "0xA069":
		return "SmartSolar MPPT 250|70 rev2"
		break
	case "0xA06A":
		return "SmartSolar MPPT 150|45 rev2"
		break
	case "0xA06B":
		return "SmartSolar MPPT 150|60 rev2"
		break
	case "0xA06C":
		return "SmartSolar MPPT 150|70 rev2"
		break
	case "0xA06D":
		return "SmartSolar MPPT 150|85 rev3"
		break
	case "0xA06E":
		return "SmartSolar MPPT 150|100 rev3"
		break
	case "0xA06F":
		return "BlueSolar MPPT 150|45 rev2"
		break
	case "0xA070":
		return "BlueSolar MPPT 150|60 rev2"
		break
	case "0xA071":
		return "BlueSolar MPPT 150|70 rev2"
		break
	case "0xA072":
		return "BlueSolar MPPT 150/45 rev3"
		break
	case "0xA073":
		return "SmartSolar MPPT 150/45 rev3"
		break
	case "0xA074":
		return "SmartSolar MPPT 75/10 rev2"
		break
	case "0xA075":
		return "SmartSolar MPPT 75/15 rev2"
		break
	case "0xA076":
		return "BlueSolar MPPT 100/30 rev3"
		break
	case "0xA077":
		return "BlueSolar MPPT 100/50 rev3"
		break
	case "0xA078":
		return "BlueSolar MPPT 150/35 rev3"
		break
	case "0xA079":
		return "BlueSolar MPPT 75/10 rev2"
		break
	case "0xA07A":
		return "BlueSolar MPPT 75/15 rev2"
		break
	case "0xA07B":
		return "BlueSolar MPPT 100/15 rev2"
		break
	case "0xA07C":
		return "BlueSolar MPPT 75/10 rev3"
		break
	case "0xA07D":
		return "BlueSolar MPPT 75/15 rev3"
		break
	case "0xA07E":
		return "SmartSolar MPPT 100/30 12V"
		break
	case "0xA07F":
		return "All-In-1 SmartSolar MPPT 75/15 12V"
		break
	case "0xA102":
		return "SmartSolar MPPT VE.Can 150/70"
		break
	case "0xA103":
		return "SmartSolar MPPT VE.Can 150/45"
		break
	case "0xA104":
		return "SmartSolar MPPT VE.Can 150/60"
		break
	case "0xA105":
		return "SmartSolar MPPT VE.Can 150/85"
		break
	case "0xA106":
		return "SmartSolar MPPT VE.Can 150/100"
		break
	case "0xA107":
		return "SmartSolar MPPT VE.Can 250/45"
		break
	case "0xA108":
		return "SmartSolar MPPT VE.Can 250/60"
		break
	case "0xA109":
		return "SmartSolar MPPT VE.Can 250/70"
		break
	case "0xA10A":
		return "SmartSolar MPPT VE.Can 250/85"
		break
	case "0xA10B":
		return "SmartSolar MPPT VE.Can 250/100"
		break
	case "0xA10C":
		return "SmartSolar MPPT VE.Can 150/70 rev2"
		break
	case "0xA10D":
		return "SmartSolar MPPT VE.Can 150/85 rev2"
		break
	case "0xA10E":
		return "SmartSolar MPPT VE.Can 150/100 rev2"
		break
	case "0xA10F":
		return "BlueSolar MPPT VE.Can 150/100"
		break
	case "0xA112":
		return "BlueSolar MPPT VE.Can 250/70"
		break
	case "0xA113":
		return "BlueSolar MPPT VE.Can 250/100"
		break
	case "0xA114":
		return "SmartSolar MPPT VE.Can 250/70 rev2"
		break
	case "0xA115":
		return "SmartSolar MPPT VE.Can 250/100 rev2"
		break
	case "0xA116":
		return "SmartSolar MPPT VE.Can 250/85 rev2"
		break
	case "0xA117":
		return "BlueSolar MPPT VE.Can 150/100 rev2"
		break
	case "0xA201":
		return "Phoenix  Inverter 12V 250VA 230V"
		break
	case "0xA202":
		return "Phoenix  Inverter 24V 250VA 230V"
		break
	case "0xA204":
		return "Phoenix  Inverter 48V 250VA 230V"
		break
	case "0xA211":
		return "Phoenix  Inverter 12V 375VA 230V"
		break
	case "0xA212":
		return "Phoenix  Inverter 24V 375VA 230V"
		break
	case "0xA214":
		return "Phoenix  Inverter 48V 375VA 230V"
		break
	case "0xA221":
		return "Phoenix  Inverter 12V 500VA 230V"
		break
	case "0xA222":
		return "Phoenix  Inverter 24V 500VA 230V"
		break
	case "0xA224":
		return "Phoenix  Inverter 48V 500VA 230V"
		break
	case "0xA231":
		return "Phoenix  Inverter 12V 250VA 230V"
		break
	case "0xA232":
		return "Phoenix  Inverter 24V 250VA 230V"
		break
	case "0xA234":
		return "Phoenix  Inverter 48V 250VA 230V"
		break
	case "0xA239":
		return "Phoenix  Inverter 12V 250VA 120V"
		break
	case "0xA23A":
		return "Phoenix  Inverter 24V 250VA 120V"
		break
	case "0xA23C":
		return "Phoenix  Inverter 48V 250VA 120V"
		break
	case "0xA241":
		return "Phoenix  Inverter 12V 375VA 230V"
		break
	case "0xA242":
		return "Phoenix  Inverter 24V 375VA 230V"
		break
	case "0xA244":
		return "Phoenix  Inverter 48V 375VA 230V"
		break
	case "0xA249":
		return "Phoenix  Inverter 12V 375VA 120V"
		break
	case "0xA24A":
		return "Phoenix  Inverter 24V 375VA 120V"
		break
	case "0xA24C":
		return "Phoenix  Inverter 48V 375VA 120V"
		break
	case "0xA251":
		return "Phoenix  Inverter 12V 500VA 230V"
		break
	case "0xA252":
		return "Phoenix  Inverter 24V 500VA 230V"
		break
	case "0xA254":
		return "Phoenix  Inverter 48V 500VA 230V"
		break
	case "0xA259":
		return "Phoenix  Inverter 12V 500VA 120V"
		break
	case "0xA25A":
		return "Phoenix  Inverter 24V 500VA 120V"
		break
	case "0xA25C":
		return "Phoenix  Inverter 48V 500VA 120V"
		break
	case "0xA261":
		return "Phoenix  Inverter 12V 800VA 230V"
		break
	case "0xA262":
		return "Phoenix  Inverter 24V 800VA 230V"
		break
	case "0xA264":
		return "Phoenix  Inverter 48V 800VA 230V"
		break
	case "0xA269":
		return "Phoenix  Inverter 12V 800VA 120V"
		break
	case "0xA26A":
		return "Phoenix  Inverter 24V 800VA 120V"
		break
	case "0xA26C":
		return "Phoenix  Inverter 48V 800VA 120V"
		break
	case "0xA271":
		return "Phoenix  Inverter 12V 1200VA 230V"
		break
	case "0xA272":
		return "Phoenix  Inverter 24V 1200VA 230V"
		break
	case "0xA274":
		return "Phoenix  Inverter 48V 1200VA 230V"
		break
	case "0xA279":
		return "Phoenix  Inverter 12V 1200VA 120V"
		break
	case "0xA27A":
		return "Phoenix  Inverter 24V 1200VA 120V"
		break
	case "0xA27C":
		return "Phoenix  Inverter 48V 1200VA 120V"
		break
	case "0xA281":
		return "Phoenix  Inverter 12V 1600VA 230V"
		break
	case "0xA282":
		return "Phoenix  Inverter 24V 1600VA 230V"
		break
	case "0xA284":
		return "Phoenix  Inverter 48V 1600VA 230V"
		break
	case "0xA291":
		return "Phoenix  Inverter 12V 2000VA 230V"
		break
	case "0xA292":
		return "Phoenix  Inverter 24V 2000VA 230V"
		break
	case "0xA294":
		return "Phoenix  Inverter 48V 2000VA 230V"
		break
	case "0xA2A1":
		return "Phoenix  Inverter 12V 3000VA 230V"
		break
	case "0xA2A2":
		return "Phoenix  Inverter 24V 3000VA 230V"
		break
	case "0xA2A4":
		return "Phoenix  Inverter 48V 3000VA 230V"
		break
	case "0xA340":
		return "Phoenix  Smart IP43 Charger 12|50 (1+1)"
		break
	case "0xA341":
		return "Phoenix  Smart IP43 Charger 12|50 (3)"
		break
	case "0xA342":
		return "Phoenix  Smart IP43 Charger 24|25 (1+1)"
		break
	case "0xA343":
		return "Phoenix  Smart IP43 Charger 24|25 (3)"
		break
	case "0xA344":
		return "Phoenix  Smart IP43 Charger 12|30 (1+1)"
		break
	case "0xA345":
		return "Phoenix  Smart IP43 Charger 12|30 (3)"
		break
	case "0xA346":
		return "Phoenix  Smart IP43 Charger 24|16 (1+1)"
		break
	case "0xA347":
		return "Phoenix  Smart IP43 Charger 24|16 (3)"
		break
	case "0xA381":
		return "BMV-712 Smart"
		break
	case "0xA382":
		return "BMV-710H Smart"
		break
	case "0xA383":
		return "BMV-712 Smart Rev2"
		break
	case "0xA389":
		return "SmartShunt 500A/50mV"
		break
	case "0xA38A":
		return "SmartShunt 1000A/50mV"
		break
	case "0xA38B":
		return "SmartShunt 2000A/50mV"
		break
	case "0xA3F0":
		return "Smart BuckBoost 12V/12V-50A"
		break
	case "0xC030":
		return "SmartShunt 500A/50mv IP65"
	}
	return "Unknown"
}

func (p *ParamsType) setValues(line string) {
	p.mu.Lock()
	defer p.mu.Unlock()

	vals := strings.Split(line, "\t")
	switch vals[0] {
	case "PID":
		p.ProductID = DecodeProductID(vals[1])
		break
	case "V":
		if v, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.Volts = v / 1000 // mV
		}
		break
	case "I":
		if i, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.Amps = i / 1000 // mA
		}
		break
	case "P":
		if pwr, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.Power = pwr // Watts
		}
		break
	case "CE":
		if ahr, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.ConsumedEnergy = ahr / 1000 // Ahr
		}
		break
	case "SOC":
		if soc, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.StateOfCharge = soc / 10 // %
		}
		break
	case "TTG":
		if ttg, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.TimeToGo = ttg // Minutes
		}
		break
	case "Alarm":
		p.Alarm = vals[1]
		break
	case "AR":
		p.AlarmReason = vals[1]
		break
	case "BMV":
		p.Model = vals[1]
		break
	case "FW":
		p.Firmware = vals[1]
		break
	case "MON":
		if v, err := strconv.ParseUint(vals[1], 10, 16); err != nil {
			log.Println(err)
		} else {
			p.MonitorMode = uint16(v)
		}
		break
	case "H1":
		if discharge, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.DeepestDischarge = discharge / 1000 // Ahr
		}
		break
	case "H2":
		if discharge, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.LastDischarge = discharge // %
		}
		break
	case "H3":
		if discharge, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.AvgDischarge = discharge // %
		}
		break
	case "H4":
		if cycles, err := strconv.ParseInt(vals[1], 10, 16); err != nil {
			log.Print(err)
		} else {
			p.Cycles = uint16(cycles) // %
		}
		break
	case "H5":
		if discharge, err := strconv.ParseInt(vals[1], 10, 16); err != nil {
			log.Print(err)
		} else {
			p.FullDischarges = uint16(discharge) // %
		}
		break
	case "H6":
		if ahr, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.CumulativeAHr = ahr // %
		}
		break
	case "H7":
		if bv, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.MinVolts = bv // %
		}
		break
	case "H8":
		if bv, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.MaxVolts = bv // %
		}
		break
	case "H9":
		if secs, err := strconv.ParseInt(vals[1], 10, 16); err != nil {
			log.Print(err)
		} else {
			p.SecondsSinceFull = uint32(secs) // %
		}
		break
	case "H10":
		if num, err := strconv.ParseInt(vals[1], 10, 32); err != nil {
			log.Print(err)
		} else {
			p.NumAutoSync = uint32(num) // %
		}
		break
	case "H11":
		if num, err := strconv.ParseInt(vals[1], 10, 32); err != nil {
			log.Print(err)
		} else {
			p.NumLowVoltAlarm = uint32(num) // %
		}
		break
	case "H12":
		if num, err := strconv.ParseInt(vals[1], 10, 32); err != nil {
			log.Print(err)
		} else {
			p.NumHighVoltAlarm = uint32(num) // %
		}
		break
	case "H15":
		if mv, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.MinAuxVolts = mv // %
		}
		break
	case "H16":
		if mv, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.MaxAuxVolts = mv // %
		}
		break
	case "H17":
		if energy, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.DischargedEnergy = energy // %
		}
		break
	case "H18":
		if energy, err := strconv.ParseFloat(vals[1], 64); err != nil {
			log.Print(err)
		} else {
			p.ChargedEnergy = energy // %
		}
		break
	}
}

func (p *ParamsType) getJSON() ([]byte, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	return json.MarshalIndent(p, "", "    ")
}
