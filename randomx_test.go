package randomx

import (
	"encoding/hex"
	"fmt"
	"runtime"
	"testing"
)

//var testPairs = [][][]byte{
//	// randomX
//	{
//		[]byte("test key 000"),
//		[]byte("This is a test"),
//		[]byte("639183aae1bf4c9a35884cb46b09cad9175f04efd7684e7262a0ac1c2f0b4e3f"),
//	},
//	// randomXL
//	{
//		[]byte("test key 000"),
//		[]byte("This is a test"),
//		[]byte("b291ec8a532bc4f78bd75b43d211e1169bb65b1a8f66d4250376ba1d6fcff1bd"),
//	},
//}
//
//func TestAllocCache(t *testing.T) {
//	cache, _ := AllocCache(FlagDefault)
//	InitCache(cache, []byte("123"))
//	ReleaseCache(cache)
//}
//
//func TestAllocDataset(t *testing.T) {
//	ds, _ := AllocDataset(FlagDefault)
//	cache, _ := AllocCache(FlagDefault)
//
//	seed := make([]byte, 32)
//	InitCache(cache, seed)
//	log.Println("rxCache initialization finished")
//
//	count := DatasetItemCount()
//	log.Println("dataset count:", count/1024/1024, "mb")
//	InitDataset(ds, cache, 0, count)
//	log.Println(GetDatasetMemory(ds))
//
//	ReleaseDataset(ds)
//	ReleaseCache(cache)
//}
//
//func TestCreateVM(t *testing.T) {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	var tp = testPairs[1]
//	cache, _ := AllocCache(FlagDefault)
//	log.Println("alloc cache mem finished")
//	seed := tp[0]
//	InitCache(cache, seed)
//	log.Println("cache initialization finished")
//
//	ds, _ := AllocDataset(FlagDefault)
//	log.Println("alloc dataset mem finished")
//	count := DatasetItemCount()
//	log.Println("dataset count:", count)
//	var wg sync.WaitGroup
//	var workerNum = uint32(runtime.NumCPU())
//	for i := uint32(0); i < workerNum; i++ {
//		wg.Add(1)
//		a := (count * i) / workerNum
//		b := (count * (i + 1)) / workerNum
//		go func() {
//			defer wg.Done()
//			InitDataset(ds, cache, a, b-a)
//		}()
//	}
//	wg.Wait()
//	log.Println("dataset initialization finished") // too slow when one thread
//	vm, _ := CreateVM(cache, ds, FlagJIT, FlagHardAES, FlagFullMEM)
//
//	var hashCorrect = make([]byte, hex.DecodedLen(len(tp[2])))
//	_, err := hex.Decode(hashCorrect, tp[2])
//	if err != nil {
//		log.Println(err)
//	}
//
//	if bytes.Compare(CalculateHash(vm, tp[1]), hashCorrect) != 0 {
//		t.Fail()
//	}
//}
//
//func TestNewRxVM(t *testing.T) {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	start := time.Now()
//	pair := testPairs[1]
//	workerNum := uint32(runtime.NumCPU())
//
//	seed := pair[0]
//	dataset, _ := NewRxDataset(FlagJIT)
//	if dataset.GoInit(seed, workerNum) == false {
//		log.Fatal("failed to init dataset")
//	}
//	//defer dataset.Close()
//	fmt.Println("Finished generating dataset in", time.Since(start).Seconds(), "sec")
//
//	vm, _ := NewRxVM(dataset, FlagFullMEM, FlagHardAES, FlagJIT, FlagSecure)
//	//defer vm.Close()
//
//	blob := pair[1]
//	hash := vm.CalcHash(blob)
//
//	var hashCorrect = make([]byte, hex.DecodedLen(len(pair[2])))
//	_, err := hex.Decode(hashCorrect, pair[2])
//	if err != nil {
//		log.Println(err)
//	}
//
//	if bytes.Compare(hash, hashCorrect) != 0 {
//		log.Println(hash)
//		t.Fail()
//	}
//}
//
//func TestCalculateHashFirst(t *testing.T) {
//	runtime.GOMAXPROCS(runtime.NumCPU())
//	start := time.Now()
//	pair := testPairs[1]
//	workerNum := uint32(runtime.NumCPU())
//
//	seed := pair[0]
//	dataset, _ := NewRxDataset(FlagJIT)
//	if dataset.GoInit(seed, workerNum) == false {
//		log.Fatal("failed to init dataset")
//	}
//	//defer dataset.Close()
//	fmt.Println("Finished generating dataset in", time.Since(start).Seconds(), "sec")
//	vm, _ := NewRxVM(dataset, FlagFullMEM, FlagHardAES, FlagJIT, FlagSecure)
//	//defer vm.Close()
//
//	targetBlob := make([]byte, 76)
//	targetNonce := make([]byte, 4)
//	binary.LittleEndian.PutUint32(targetNonce, 2333)
//	copy(targetBlob[39:43], targetNonce)
//
//	targetResult := vm.CalcHash(targetBlob)
//
//	var wg sync.WaitGroup
//	for i := 0; i < runtime.NumCPU(); i++ {
//		vm, _ := NewRxVM(dataset, FlagFullMEM, FlagHardAES, FlagJIT, FlagSecure)
//
//		wg.Add(1)
//		blob := make([]byte, 76)
//		vm.CalcHashFirst(blob)
//
//		n := uint32(0)
//		go func() {
//			for {
//				n++
//				nonce := make([]byte, 4)
//				binary.LittleEndian.PutUint32(nonce, n)
//				copy(blob[39:43], nonce)
//				result := vm.CalcHashNext(blob)
//				if bytes.Compare(result, targetResult) == 0 {
//					fmt.Println(n, "found")
//					wg.Done()
//				} else {
//					//fmt.Println(n, "failed")
//				}
//			}
//		}()
//	}
//	wg.Wait()
//
//}
//
//// go test -v -bench "." -benchtime=30m
//func BenchmarkCalculateHash(b *testing.B) {
//	cache, _ := AllocCache(FlagDefault)
//	ds, _ := AllocDataset(FlagDefault)
//	InitCache(cache, []byte("123"))
//	FastInitFullDataset(ds, cache, uint32(runtime.NumCPU()))
//	vm, _ := CreateVM(cache, ds, FlagDefault)
//	for i := 0; i < b.N; i++ {
//		nonce := strconv.FormatInt(rand.Int63(), 10) // just test
//		CalculateHash(vm, []byte("123"+nonce))
//	}
//
//	DestroyVM(vm)
//}

func TestCalculateHash(t *testing.T) {
	cache, _ := AllocCache(FlagDefault)
	ds, _ := AllocDataset(FlagDefault)
	seedBytes, _ := hex.DecodeString("e2f3cee2d752d17997b1a000feacc0f21f0edae62cfd189b1568786b28a2a167")
	InitCache(cache, []byte(seedBytes))

	FastInitFullDataset(ds, cache, uint32(runtime.NumCPU()))
	vm, _ := CreateVM(cache, ds, FlagDefault)

	tp := "0e0eceb6d9830694b14d501222f7da7490d0b2d326c3750f08e60f101c285331d67e7ffc0be17c00000000161708ce1b2846cbe8bea4593c1a56cc8e3c1222ed0a7898961cee642d93af042e"

	m1 := string(tp[0:78]) + "a52e0000" + string(tp[86:])
	m, _ := hex.DecodeString(m1)
	h := CalculateHash(vm, []byte(m))
	fmt.Println("TestCalculateHash1:", hex.EncodeToString(h))

	m1 = string(tp[0:78]) + "0cb90000" + string(tp[86:])
	m, _ = hex.DecodeString(m1)
	h = CalculateHash(vm, []byte(m))
	fmt.Println("TestCalculateHash2:", hex.EncodeToString(h))

	m1 = string(tp[0:78]) + "836d0100" + string(tp[86:])
	m, _ = hex.DecodeString(m1)
	h = CalculateHash(vm, []byte(m))
	fmt.Println("TestCalculateHash3:", hex.EncodeToString(h))

	DestroyVM(vm)
}
