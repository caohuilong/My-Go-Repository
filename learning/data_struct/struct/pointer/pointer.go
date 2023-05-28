package main

import (
	"fmt"
	"sync"
)

type RayCluster struct {
	lock sync.RWMutex
	worker2GamecoreMap map[int32]*WorkerInfo
}

type WorkerInfo struct {
	workerId int32
	connectedGamecores map[int32]struct{}
}

func main() {
	rc := &RayCluster{
		worker2GamecoreMap: make(map[int32]*WorkerInfo),
	}

	rc.worker2GamecoreMap[1] = &WorkerInfo{
		workerId: 1,
		connectedGamecores: map[int32]struct{}{
			1: struct{}{},
		},
	}

	fmt.Printf("%+v\n", rc)

	test1(rc)
	fmt.Printf("%+v\n", rc)

	test2(rc)
	fmt.Printf("%+v\n", rc)

	test3(rc)
	for _, workerInfo := range rc.worker2GamecoreMap {
		fmt.Printf("%+v\n", workerInfo)
	}

}

func test1(r *RayCluster) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.worker2GamecoreMap[2] = &WorkerInfo{
		workerId: 2,
		connectedGamecores: map[int32]struct{}{
			2: struct{}{},
		},
	}
	for _, workerInfo := range r.worker2GamecoreMap {
		fmt.Printf("%+v\n", workerInfo)
	}
}

func test2(r *RayCluster) {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, workerInfo := range r.worker2GamecoreMap {
		workerInfo.connectedGamecores[5] = struct{}{}
	}
	for _, workerInfo := range r.worker2GamecoreMap {
		fmt.Printf("%+v\n", workerInfo)
	}
}

func test3(r *RayCluster) {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, workerInfo := range r.worker2GamecoreMap {
		for gamecoreId, _ := range workerInfo.connectedGamecores {
			delete(workerInfo.connectedGamecores, gamecoreId)
		}
	}
	for _, workerInfo := range r.worker2GamecoreMap {
		fmt.Printf("%+v\n", workerInfo)
	}
}