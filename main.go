package main

import (
	"os"
)

func main() {
	descriptors := GetDescriptors()

	vehicles := FilterByDtype(descriptors, "vehicle")
	volvoDescriptors := FilterByBodyContains(descriptors, "volvo.fh16_2012")
	volvoAccessories := FilterByDtype(volvoDescriptors, "vehicle_accessory", "vehicle_addon_accessory")

	// for _, accessory := range volvoAccessories {
	// 	PrintDescriptor(accessory)
	// }

	for _, vehicle := range vehicles {
		accessories := GetEachAccessoryOfVehicle(vehicle, volvoAccessories)

		if len(accessories) == 0 {
			continue
		}

		file, err := os.OpenFile(vehicle.uid, os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		for _, accessory := range accessories {
			text := BuildDescriptorString(accessory)
			file.WriteString(text)
			file.WriteString("\n")
		}
	}
}
