package seismicactivity

import (
	placeofreceiptinfo "app/placeOfReceiptInfo"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func SeismicActivitySensor(ctx context.Context, transferPoint chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Seismic Activity sensor has finished working")
			close(transferPoint)
			return
		default:
			receivedData := placeofreceiptinfo.RandomPlace()
			n := rand.Float64() * 10
			time.Sleep(time.Second)
			transferPoint <- fmt.Sprintf("%s Seismic event: %.1f Mw", receivedData, n)
		}
	}
}
