package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*
	Problem Statement :-
		-> An auto stand had around 5 autos waiting for the passengers
		-> As soon as the passengers comes, the auto leading the pack will pick the passenger
			and travel to the destination and return back and wait in the same queue.
		-> By the time the auto comes back, it can be the leading auto in the group again or
			it may be at the last depending on the incoming customers.
		-> If all the auto's are busy, the customer keep waiting
		-> Once the time comes, all the autos had to be popped from the stand.
 */

/*
	Assumptions -> Customer comes at random time interval from anywhere b/w 1-5 seconds.
	Drop -> May take any time b/w 1-10 seconds.
*/


type autoChan struct {
	driverName string
	autoNumber string
}

type customerChan struct {
	customerName string
}

func main() {
	autoStand := make(chan autoChan, 5) // Create a auto stand with 5 auto's
	customers := make(chan customerChan, 5)
	doneForDay := make(chan bool)
	for i:=int64(0) ; i< 5; i++ {
		driver := "driver-"+strconv.FormatInt(i+1,10)
		autoNumber := "auto-"+strconv.FormatInt(i+1,10)
		log.Println("Driver "+ driver+ " started waiting for passengers")
		autoStand <- autoChan{
			driverName: driver,
			autoNumber: autoNumber,
		}
	}
	go func() {
		customerCount := int64(0)
		for {
			customerCount += 1
			newCustomerInterval := rand.Int63n(5)
			customerName := "customer-"+strconv.FormatInt(customerCount,10)
			time.Sleep(time.Duration(newCustomerInterval)*time.Second)
			log.Println("New customer "+ customerName +" arrived ")
			customers <- customerChan{
				customerName: customerName,
			}
		}
	}()
	go func(){
		<- doneForDay
		os.Exit(0)
	}()
	go func(){
		time.Sleep(30*time.Millisecond)
		doneForDay <- true
	}()
	for customerData := range customers{
		auto := <-autoStand
		log.Println("customer "+ customerData.customerName + " received by auto "+ auto.autoNumber)
		go func(auto autoChan) {
			dropTime := rand.Int63n(10)
			time.Sleep(time.Duration(dropTime)*time.Second)
			log.Println(fmt.Sprintf("Auto %v completed its ride in %v seconds and started waiting in queue",auto.autoNumber, dropTime))
			autoStand <- auto
		}(auto)
	}
}


