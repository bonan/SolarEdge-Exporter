/*

MIT License

Copyright (c) 2019 David Suarez

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/
package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	SunSpec_DID = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "SunSpec_DID",
		Help: "101 = single phase 102 = split phase1 103 = three phase",
	})

	SunSpec_Length = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "SunSpec_Length",
		Help: "Registers 50 = Length of model block",
	})

	acCurrent = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Current",
		Help: "Amps AC Total Current value",
	}, []string{"phase"})

	AC_Current  = acCurrent.WithLabelValues("total")
	AC_CurrentA = acCurrent.WithLabelValues("A")
	AC_CurrentB = acCurrent.WithLabelValues("B")
	AC_CurrentC = acCurrent.WithLabelValues("C")

	acVoltage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "AC_Voltage",
		Help: "Volts AC Voltage Phase AB value",
	}, []string{"phase", "type"})

	AC_VoltageAB = acVoltage.WithLabelValues("AB", "l2l")
	AC_VoltageBC = acVoltage.WithLabelValues("BC", "l2l")
	AC_VoltageCA = acVoltage.WithLabelValues("CA", "l2l")
	AC_VoltageAN = acVoltage.WithLabelValues("A", "l2n")
	AC_VoltageBN = acVoltage.WithLabelValues("B", "l2n")
	AC_VoltageCN = acVoltage.WithLabelValues("C", "l2n")

	AC_Power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Power",
		Help: "Watts AC Power value",
	})
	AC_Frequency = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Frequency",
		Help: "Hertz AC Frequency value",
	})

	AC_VA = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_VA",
		Help: "VA Apparent Power",
	})

	AC_VAR = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_VAR",
		Help: "VAR Reactive Power",
	})

	AC_PF = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_PF",
		Help: "% Power Factor",
	})

	AC_Energy_WH = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "AC_Energy_WH",
		Help: "WattHours AC Lifetime Energy production",
	})

	DC_Current = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Current",
		Help: "Amps DC Current value",
	})

	DC_Voltage = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Voltage",
		Help: "Volts DC Voltage value",
	})

	DC_Power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "DC_Power",
		Help: "Watts DC Power value",
	})

	Temp_Sink = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Temp_Sink",
		Help: "Degrees C Heat Sink Temperature",
	})

	Status = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Status",
		Help: "Operating State",
	})

	Status_Vendor = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "Status_Vendor",
		Help: "Vendor-defined operating state and error codes. For error description, meaning and troubleshooting, refer to the SolarEdge Installation Guide.",
	})

	// Meter

	M_SunSpec_DID = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_SunSpec_DID",
		Help: "",
	})

	M_SunSpec_Length = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_SunSpec_Length",
		Help: "",
	})

	M_AC_Current = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Current",
		Help: "Amps AC Total Current value",
	})

	M_AC_CurrentA = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_CurrentA",
		Help: "Amps AC Phase A Current value",
	})

	M_AC_CurrentB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_CurrentB",
		Help: "Amps AC Phase B Current value",
	})

	M_AC_CurrentC = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_CurrentC",
		Help: "Amps AC Phase C Current value",
	})

	M_AC_VoltageLN = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageLN",
		Help: "Volts AC Voltage Phase AB value",
	})

	M_AC_VoltageAN = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageAN",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageBN = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageBN",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageCN = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageCN",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageLL = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageLL",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageAB = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageAB",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageBC = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageBC",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_VoltageCA = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VoltageCA",
		Help: "Volts AC Voltage Phase BC value",
	})

	M_AC_Frequency = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Frequency",
		Help: "Hertz AC Frequency value",
	})

	M_AC_Power = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Power",
		Help: "Watts AC Power value",
	})

	M_AC_Power_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Power_A",
		Help: "Watts AC Power value",
	})

	M_AC_Power_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Power_B",
		Help: "Watts AC Power value",
	})

	M_AC_Power_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_Power_C",
		Help: "Watts AC Power value",
	})

	M_AC_VA = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VA",
		Help: "VA Apparent Power",
	})

	M_AC_VA_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VA_A",
		Help: "VA Apparent Power",
	})

	M_AC_VA_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VA_B",
		Help: "VA Apparent Power",
	})

	M_AC_VA_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VA_C",
		Help: "VA Apparent Power",
	})

	M_AC_VAR = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VAR",
		Help: "VAR Reactive Power",
	})

	M_AC_VAR_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VAR_A",
		Help: "VAR Reactive Power",
	})

	M_AC_VAR_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VAR_B",
		Help: "VAR Reactive Power",
	})

	M_AC_VAR_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_VAR_C",
		Help: "VAR Reactive Power",
	})

	M_AC_PF = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_PF",
		Help: "% Power Factor",
	})

	M_AC_PF_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_PF_A",
		Help: "% Power Factor",
	})

	M_AC_PF_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_PF_B",
		Help: "% Power Factor",
	})

	M_AC_PF_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_AC_PF_C",
		Help: "% Power Factor",
	})

	M_Exported = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Exported",
		Help: "WattHours AC Exported",
	})

	M_Exported_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Exported_A",
		Help: "WattHours AC Exported",
	})

	M_Exported_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Exported_B",
		Help: "WattHours AC Exported",
	})

	M_Exported_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Exported_C",
		Help: "WattHours AC Exported",
	})

	M_Imported = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Imported",
		Help: "WattHours AC Imported",
	})

	M_Imported_A = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Imported_A",
		Help: "WattHours AC Imported",
	})

	M_Imported_B = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Imported_B",
		Help: "WattHours AC Imported",
	})

	M_Imported_C = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "M_Imported_C",
		Help: "WattHours AC Imported",
	})
)
