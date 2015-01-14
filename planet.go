package main

import (
	"fmt"
	"math/rand"
	"time"
)

// // defaults

// land: 1000, // available area
// agr: 0,
// agr_max: 0,
// pop: 0,
// birthrate: 0.002, //( 50 + random0to(50)) * 0.0005,
// deathrate: 0.001,
// popup: 500 + random0to(200) * 0.0030,
// indup: 500 + random0to(200) * 0.0030,
// agrup: 250 + random0to(200) * 0.15,
// shipcost: 5000,
// //
// pol: 0,
// ind: 0,
// cr: 1000,

// var land = 10000 * random.from1to(5);
// var agr = (15 + random.from1to(25)) * land/100;

// this.set({
//   id: uuid.v4(),
//   r: r,
//   a: a,
//   v: v,
//   pop: (15 + random.from1to(25)) * land/100,
//   agr: agr,
//   agr: agr_max,
//   ind: (5 + random.from1to(5)) * land/100,
//   pol: 0,
//   size: 2 + random.from0to(3),
//   land: land,
//   cr: 1000,
//   shipcost: 2500 + random0to(2500)
// });

// var sum = data.pop + data.pol + data.agr + data.ind;
//console.log('shrink', data.land, sum);
// if(sum > data.land){
//   data.pop *= 0.95;
//   data.ind *= 0.95;
//   data.agr *= 0.95;
// }

// // pollution recovery
// data.pol = data.pol * 0.90;

// data.pop += data.popup;

// // pop consumes agr
// data.agr -= data.pop * 0.0005;

// // pol kills agr
// data.agr -= data.pol;

// if(data.agr > data.pop){
//   data.agr = data.pop;
// }

// // agr up
// data.agr += data.agrup;

// //ind up
// data.ind += data.indup;
// data.ind += data.pop * 0.01;

// data.pol += ((data.ind + data.pop) / data.land)  * 1;

// if(data.ind > data.pop * 0.95){
//   data.ind *= 0.90;
// }

//   var earnings = (data.ind/100 +data.pop/100) * ((50 + random0to(50))/100);

//   data.d_cr = earnings;
//   data.cr += earnings;

func population(report, pol chan float64) {
	val := rand.Float64() * 1000.0
	fmt.Printf("population starting at %v \n", val)

	for {
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

		select {
		case p := <-pol:
			pollutionDelta := p * 0.1
			growthDelta := val * 0.1

			val = val + growthDelta - pollutionDelta
			fmt.Printf("population %v \n", val)
		case report <- val:
			// fmt.Println("reporting population")
		}
	}

}

func pollution(report, pop chan float64) {
	val := rand.Float64() * 1000.0
	fmt.Printf("pollution starting at %v \n", val)

	for {
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		select {
		case populationCount := <-pop:

			val = populationCount
			fmt.Printf("pollution %v \n", val)
		case report <- val:
			// fmt.Println("reporting pollution")
		}
	}

}

func main() {
	pop := make(chan float64, 1)
	pol := make(chan float64, 1)

	go population(pop, pol)
	go pollution(pol, pop)
	time.Sleep(2 * time.Second)
}
