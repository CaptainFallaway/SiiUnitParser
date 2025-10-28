package main

// import (
// 	"fmt"
// 	"os"
// 	"slices"
// 	"strings"
// )

// func BuildDescriptorString(d Unit) string {
// 	var b strings.Builder
// 	fmt.Fprintf(&b, "%s : %s {\n", d.utype, d.uid)
// 	for _, line := range d.attrs {
// 		fmt.Fprintf(&b, "\t%s\n", line)
// 	}
// 	fmt.Fprint(&b, "}")
// 	return b.String()
// }

// func PrintDescriptor(d Unit) {
// 	fmt.Println(BuildDescriptorString(d))
// }

// func FilterByDtype(ds []Unit, dtypes ...string) []Unit {
// 	newDs := make([]Unit, 0)

// 	for _, dtype := range dtypes {
// 		for _, d := range ds {
// 			if d.utype == dtype {
// 				newDs = append(newDs, d)
// 			}
// 		}
// 	}

// 	return newDs
// }

// func FilterByBodyContains(ds []Unit, substr string) []Unit {
// 	newDs := make([]Unit, 0)

// 	for _, d := range ds {
// 		if d.BodyContains(substr) {
// 			newDs = append(newDs, d)
// 		}
// 	}

// 	return newDs
// }

// func CountVehicleAccessories(vehicle Unit, refCountMap map[string]int) {
// 	for _, line := range vehicle.attrs {
// 		if strings.Contains(line, "accessories[") {
// 			splitLine := strings.Split(line, ": ")
// 			uid := splitLine[1]
// 			refCountMap[uid]++
// 		}
// 	}
// }

// func UidInUidSlice(uidSlice []string, uid string) bool {
// 	return slices.Contains(uidSlice, uid)
// }

// func GetDescriptors() []Unit {
// 	file, err := os.OpenFile("game-decoded.sii", os.O_RDONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	return ReadAllDescriptors(file)
// }

// func GetEachAccessoryOfVehicle(vehicle Unit, accessories []Unit) []Unit {
// 	newSlice := make([]Unit, 0)

// 	for _, line := range vehicle.attrs {
// 		if strings.Contains(line, "accessories[") {
// 			splitLine := strings.Split(line, ": ")
// 			uid := splitLine[1]
// 			for _, accessory := range accessories {
// 				if accessory.uid == uid {
// 					newSlice = append(newSlice, accessory)
// 				}
// 			}
// 		}
// 	}

// 	return newSlice
// }
