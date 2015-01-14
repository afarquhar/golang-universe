package main

import (
	"fmt"
	// "math"
	"math/rand"
	"time"
)

type Update struct {
	Name  string
	Value float64
}

type Grower interface {
	grow() float64
	tick()
}

type Stat struct {
	Name     string
	Value    float64
	Reporter chan Update
	Inputs   []chan Update
}

type Population Stat

func (pop *Population) grow() float64 {
	return 500 + (rand.Float64()*200)*0.0030
}

func (pop *Population) tick() {

	pop.Value += pop.grow()

	for _, input := range pop.Inputs {
		select {
		case update := <-input:
			fmt.Printf("%v received %v from %v\n", pop.Name, update.Value, update.Name)
		case <-time.After(time.Millisecond * 100):
			// fmt.Printf("fuck it move on in %v\n", pop.Name)
		}

	}
	// fmt.Printf("writing %v reporter\n", pop.Name)
	pop.Reporter <- Update{pop.Name, pop.Value}
}

func float64From1To(n int) float64 {
	i := rand.Intn(n) + 1
	return float64(i)
}

func start(stat Grower) {
	go func(stat Grower) {
		t := time.Tick(time.Duration(500) * time.Millisecond)

		for range t {
			// fmt.Println(now)
			stat.tick()
		}

	}(stat)
}

func main() {

	land := 10000 * float64From1To(5)
	fmt.Printf("land: %v \n", land)
	// // (15 + random.from1to(25)) * land/100)
	popInit := (15 + float64From1To(25)) * (land / 100)

	fmt.Printf("popInit: %v \n", popInit)

	polInit := popInit / land
	fmt.Printf("polInit: %v \n", polInit)

	popReporter := make(chan Update, 1)
	polReporter := make(chan Update, 1)

	popInputs := []chan Update{polReporter}
	// polInputs := []chan Update{popReporter}

	pop := &Population{"population", popInit, popReporter, popInputs}
	// pol := Stat{"pollution", popInit, popGrowth, polReporter, polInputs}

	start(pop)
	// start(pol)

	time.Sleep(time.Duration(5) * time.Second)

}
