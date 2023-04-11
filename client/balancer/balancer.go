package balancer

var BalanceCh = make(chan bool)

func Balancer() {

	go func() {
		for {
			cp := getCpuUsage()
			lat := GetLatency()

			if cp > 80 || lat == "" {
				BalanceCh <- true
				//
			}
		}
	}()

}
