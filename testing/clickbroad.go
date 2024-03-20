package test

import (
    "fmt"
    "time"
)

// // Periodically record some sample latencies for the three services.
// go func() {
//     for {
//         v := rand.Float64() * *uniformDomain
//         rpcDurations.WithLabelValues("uniform").Observe(v)
//         time.Sleep(time.Duration(100*oscillationFactor()) * time.Millisecond)
//     }
// }()
//
// go func() {
//     for {
//         v := (rand.NormFloat64() * *normDomain) + *normMean
//         rpcDurations.WithLabelValues("normal").Observe(v)
//         // Demonstrate exemplar support with a dummy ID. This
//         // would be something like a trace ID in a real
//         // application.  Note the necessary type assertion. We
//         // already know that rpcDurationsHistogram implements
//         // the ExemplarObserver interface and thus don't need to
//         // check the outcome of the type assertion.
//         rpcDurationsHistogram.(prometheus.ExemplarObserver).ObserveWithExemplar(
// , prometheus.Labels{"dummyID": fmt.Sprint(rand.Intn(100000))},
//         )
//         time.Sleep(time.Duration(75*oscillationFactor()) * time.Millisecond)
//     }
// }()
//
// go func() {
//     for {
//         v := rand.ExpFloat64() / 1e6
//         rpcDurations.WithLabelValues("exponential").Observe(v)
//         time.Sleep(time.Duration(50*oscillationFactor()) * time.Millisecond)
//     }
// }()
//
//
// ticker := time.NewTicker(5 * time.Second)
// quit := make(chan struct{})
// go func() {
//     for {
//         select {
//         case <- ticker.C:
//             // do stuff
//         case <- quit:
//             ticker.Stop()
//             return
//         }
//     }
// }()
//
//
//
//
// ch := make(chan int, 1)
// go func() { for { ch <- 1 } } ()
// L:
// for {
//     select {
//     case <-ch:
//         // do something
//     case <-time.After(timeoutNs):
//         // call timed out
//         break L
//     }
// }
