package gamelogic

type SafetyCar struct {
	IsDeployed      bool
	IsGoneInThisLap bool
}

func (s *SafetyCar) OnDriverStuckedInSession(eventData interface{}) {
	// change IsDeployed to true

}
