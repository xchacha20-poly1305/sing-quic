package udphop

// MinimumHopInterval limits the littlest seconds for hop interval.
const MinimumHopInterval = 5

type UDPHopOption struct {
	HopPorts    string
	HopInterval int
}
