package airhumidity

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	placeofreceiptinfo "app/placeOfReceiptInfo"
)

func AirHumiditySensor(ctx context.Context, transferPoint chan<- string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Air humidity sensor has finished working")
			close(transferPoint)
			return
		default:
			receivedData := placeofreceiptinfo.RandomPlace()
			n := rand.Intn(100)
			time.Sleep(time.Second)
			transferPoint <- fmt.Sprintf("%s Humidity: %d%%", receivedData, n)
		}
	}
}
