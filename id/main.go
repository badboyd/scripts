package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"
)

// id 64 bit
// [64-27] timestamp / timestamp will be sharding % 100
// [26-17] increase number
// [16-1] ip

const (
	maskSequence = uint16(1<<sequenceLen - 1)

	idLen        = 63
	timestampLen = 39
	sequenceLen  = 8
	machineIDlen = 16
)

func privateIPv4() (net.IP, error) {
	as, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, a := range as {
		ipnet, ok := a.(*net.IPNet)
		if !ok || ipnet.IP.IsLoopback() {
			continue
		}

		ip := ipnet.IP.To4()
		if isPrivateIPv4(ip) {
			return ip, nil
		}
	}
	return nil, errors.New("no private ip address")
}

func isPrivateIPv4(ip net.IP) bool {
	return ip != nil &&
		(ip[0] == 10 || ip[0] == 172 && (ip[1] >= 16 && ip[1] < 32) || ip[0] == 192 && ip[1] == 168)
}

func lower16BitPrivateIP() (uint16, error) {
	ip, err := privateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

type idGenerator struct {
	m         sync.Mutex
	sequence  uint8
	machineID uint16
}

func newIDGenerator() (*idGenerator, error) {
	machineID, err := lower16BitPrivateIP()
	if err != nil {
		return nil, err
	}

	return &idGenerator{
		sequence:  0,
		machineID: machineID,
	}, nil
}

func (i *idGenerator) next() uint64 {
	i.m.Lock()
	defer i.m.Unlock()

	i.sequence++
	fmt.Println("seq: ", i.sequence)

	t := time.Now()

	// fmt.Printf("nanosec %64b %d\n", t.UnixNano(), t.UnixNano())
	// fmt.Printf("milisec %64b %d\n", t.UnixNano()/int64(time.Millisecond), t.UnixNano()/int64(time.Millisecond))
	// fmt.Printf("10 milisec %64b %d\n", t.UnixNano()/int64(10*time.Millisecond), t.UnixNano()/int64(10*time.Millisecond))
	// fmt.Printf("1000 ms %64b %d\n", t.UnixNano()/int64(1000*time.Millisecond), t.UnixNano()/int64(1000*time.Millisecond))
	// fmt.Printf("sec %64b %d\n", t.Unix(), t.Unix())
	// fmt.Printf("%64b\n", t<<(sequenceLen+machineIDlen))

	return uint64(t.UnixNano()/int64(10*time.Millisecond))<<(sequenceLen+machineIDlen) |
		uint64(i.sequence)<<machineIDlen | uint64(i.machineID)
}

const (
	maskMachineID  = uint64(1<<machineIDlen - 1)
	maskIDSequence = uint64((1<<sequenceLen - 1) << machineIDlen)
)

func decompose(id uint64) map[string]uint64 {
	msb := id >> 63
	time := id >> (sequenceLen + machineIDlen)
	sequence := id & maskIDSequence >> machineIDlen
	machineID := id & maskMachineID

	return map[string]uint64{
		"id":         id,
		"msb":        msb,
		"time":       time,
		"sequence":   sequence,
		"machine_id": machineID,
	}
}

func main() {
	// fmt.Printf("%+v", decompose(260507599508602904))
	// 2605308826471104514
	for k, v := range decompose(2605310040588026278) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(2605331534349074441) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(2605331233936244745) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(2605331237409128457) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(customID(9, 155289860627, 1)) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(customID(9, 155289860627, 2)) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}
	fmt.Println()
	for k, v := range decompose(2605346043233632284) {
		fmt.Printf("%s %d %064b \n ", k, v, v)
	}

	// test()
	// 1110 0111 0110 0000 1001 0010 0000 0100 10
	// 1011 0100 1100 0011 1000 1000 0111 0011 0000 1001 0
	// 1001 0000 1001 1100 0110 1110 1111 0110 1001 11
}

func customID(machineID uint16, t int64, sequence uint8) uint64 {
	return uint64(t)<<(sequenceLen+machineIDlen) |
		uint64(sequence)<<machineIDlen | uint64(machineID)
}

func test() {
	idGen, err := newIDGenerator()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	defer func() {
		fmt.Println("run time: ", time.Since(start).Seconds()*1000)
	}()

	fmt.Println(maskSequence)
	fmt.Println(idGen.machineID)

	var wg sync.WaitGroup
	res := make(chan uint64, 10)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				id := idGen.next()
				res <- id
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	tmp := make(map[uint64]bool)

	for {
		r, ok := <-res
		if !ok {
			break
		}
		fmt.Printf("id: %064b\n", r)
		fmt.Printf("id: %d\n", r)
		if tmp[r] {
			panic("dup id")
		} else {
			tmp[r] = true
		}
	}
}
