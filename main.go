package main

import (
	"github.com/aodeniyide/stale-ami-alerts/ami"
)

func main() {
	var ownerid string
	var region string
	getStatus := ami.QueryAmi(ownerid, region)
	//	fmt.Println(mole)

	for i, t := range getStatus {
		ami.AlertStaleAmi(i, ami.ProcessStaleAmi(ami.UpdateTime(t)), 90)
	}
}
