// Part6: task2
// Написать многопоточную программу, в которой будет использоваться явный вызов
// планировщика. Выполните трассировку программы

package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type Job struct {
	id           int
	randomNumber int
}

type Result struct {
	job Job
	sum int
}

//Writing both channels here, in order all func have acess to them
var jobs = make(chan Job, 10) // 10 - is the buffer cherez kotoriy budem tolkat resultat
var results = make(chan Result, 10)

func Digits(number int) int {
	sum := 0
	num := number
	for num != 0 {
		digit := num % 10
		sum += digit
		num /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func Worker(wg *sync.WaitGroup) { // creating worker here
	for job := range jobs {
		output := Result{job, Digits(job.randomNumber)}
		results <- output
	}
	wg.Done()
}

func CreatingWorkerPool(NumberOfWorkers int) { // creating pool of goroutines (Workers)
	var wg sync.WaitGroup
	for i := 0; i < NumberOfWorkers; i++ {
		wg.Add(1)
		go Worker(&wg)
	}
	wg.Wait()
	close(results)
}

func Allocator(NumberOfJobs int) { // eta funksiya generit raboti i sopostavlayet kajdomu wWrker-u
	for i := 0; i < NumberOfJobs; i++ {
		randNum := rand.Intn(999)
		job := Job{i, randNum}
		jobs <- job
	}
	close(jobs)
}

func ResultFunc(done chan bool) { // schitivayem serultati
	for res := range results { // slushayem iz kanala results danniye poka close() ne vipadet
		fmt.Printf("Job id is %d, value is %d, result is %d\n", res.job.id, res.job.randomNumber, res.sum)
	}
	done <- true // yesli true (buleva zaglushka), to done
}

func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

	startTimer := time.Now()
	JobsNumber := 150
	go Allocator(JobsNumber) //Naplodit Workov skolko nujno, v dannom sluchaye 150
	done := make(chan bool)
	go ResultFunc(done)
	NumOfWorkers := 10 // skolko Workerov budet soderjatsa v Worker Pool-e 	// NOTE!!! Workerov ne doljno bit slishkom mnogo, nujno nayti balans, inache budut kak lishniye rabochiye na stroyke - jrat resursi
	CreatingWorkerPool(NumOfWorkers)
	<-done                            // dojidayemsa poka vsa rabota ne budet vipolnena (vivedeno v konsolku)
	endTime := time.Now()             // Vrema kontsa vipolneniya
	deltaT := endTime.Sub(startTimer) //deltaT is the delta time
	fmt.Println("Total time: ", deltaT.Seconds(), "sec.")
}


// ramin@ramin:~/go/src/GB/еуыеы$ GOMAXPROCS=1 go run main_goroutines.go 2>trace-goroutines.out
// ramin@ramin:~/go/src/GB/еуыеы$ go tool trace trace-goroutines.out