package airpressure

import (
	placeofreceiptinfo "app/placeOfReceiptInfo"
	"context"
	"fmt"
	"math/rand"
	"time"
)

func AirPressureSencor(ctx context.Context, transferPoint chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Air pressure sensor has finished working")
			close(transferPoint)
			return
		default:
			receivedData := placeofreceiptinfo.RandomPlace()
			n := rand.Intn(601) + 900
			time.Sleep(time.Second)
			transferPoint <- fmt.Sprintf("%s Air pressure: %d hPa", receivedData, n)
		}
	}
}
