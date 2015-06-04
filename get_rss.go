package gs

/*
#cgo LDFLAGS: -L/usr/local/lib -lget_rss -lstdc++

#include <stdlib.h>
#include <get_rss.h>
*/
import "C"

import (
	"fmt"
)

func GetPeakRSS() int {
	var res int
	res = int(C.getPeakRSS())
	return res
}

func GetCurrentRSS() int {
	var res int
	res = int(C.getCurrentRSS())
	return res
}

func PrintRss() {
	// fmt.Println("PeakRSS: %dK, CurrentRSS: %dK\n", GetPeakRSS()/1024, GetCurrentRSS()/1024)
	// fmt.Printf("PeakRSS: %vK, CurrentRSS: %vK\n", C.getPeakRSS(), C.getCurrentRSS())
	fmt.Printf("PeakRSS: %vKB, CurrentRSS: %vKB\n", GetPeakRSS()/1024, GetCurrentRSS()/1024)
}

func PrintRssB() {
	// fmt.Println("PeakRSS: %dK, CurrentRSS: %dK\n", GetPeakRSS()/1024, GetCurrentRSS()/1024)
	// fmt.Printf("PeakRSS: %vK, CurrentRSS: %vK\n", C.getPeakRSS(), C.getCurrentRSS())
	fmt.Printf("PeakRSS: %vB, CurrentRSS: %vB\n", GetPeakRSS(), GetCurrentRSS())
}

func StrRssB() string {
	return fmt.Sprintf("PeakRSS: %vB, CurrentRSS: %vB\n", GetPeakRSS(), GetCurrentRSS())
}
