package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	airhumidity "app/airHumidity"
	airpressure "app/airPressure"
	seismicactivity "app/seismicActivity"
)

var (
	dataStorage []string
	mtx         sync.Mutex
)

func main() {
	wg := &sync.WaitGroup{}
	pressureCtx, pressureCtxCancel := context.WithCancel(context.Background())
	humidityCtx, humidityCtxCancel := context.WithCancel(context.Background())
	seismicactivityCtx, seismicactivityCtxCancel := context.WithCancel(context.Background())
	pressureCh := make(chan string)
	humidityCh := make(chan string)
	seismicactivityCh := make(chan string)

	go airhumidity.AirHumiditySensor(humidityCtx, humidityCh)
	go airpressure.AirPressureSencor(pressureCtx, pressureCh)
	go seismicactivity.SeismicActivitySensor(seismicactivityCtx, seismicactivityCh)

	wg.Add(3)
	go manageGoroutine(pressureCtxCancel, pressureCh, wg, 5)
	go manageGoroutine(humidityCtxCancel, humidityCh, wg, 10)
	go manageGoroutine(seismicactivityCtxCancel, seismicactivityCh, wg, 15)
	wg.Wait()
}

func manageGoroutine(CtxCancel context.CancelFunc, ch chan string, wg *sync.WaitGroup, workTime time.Duration) {
	defer wg.Done()
	go func() {
		time.Sleep(workTime * time.Second)
		CtxCancel()
	}()
	for data := range ch {
		fmt.Println(data)
		mtx.Lock()
		dataStorage = append(dataStorage, data)
		mtx.Unlock()
	}
}
