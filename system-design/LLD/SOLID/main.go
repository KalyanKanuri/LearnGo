package main

import (
	dip "solid/DIP"
	isp "solid/ISP"
	ocp "solid/OCP"
)

func main() {
	ocp.OCP()
	isp.ISP()
	dip.DIP()
}
