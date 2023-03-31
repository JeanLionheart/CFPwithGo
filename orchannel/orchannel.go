package orchannel

var closeChannel chan any = make(chan any)

func init() {
	close(closeChannel)
}

/* wait until any one channel closed */
func Wait(channelsToWait ...<-chan any) <-chan any {
	var orDone chan any
	switch len(channelsToWait) {
	case 0:
		return closeChannel
	case 1:
		return channelsToWait[0]
	case 2:
		select {
		case <-channelsToWait[0]:
		case <-channelsToWait[1]:
		}
		return closeChannel
	default: //any of ch0 or ch1 closed, the new done in that goroutine will be closed too
		orDone = make(chan any)
		go func() {
			defer close(orDone)
			select {
			case <-channelsToWait[0]:
			case <-channelsToWait[1]:
			case <-Wait(append(channelsToWait[2:], orDone)...):
			}
		}()
	}
	return orDone
}
