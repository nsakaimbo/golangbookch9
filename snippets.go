// Package bank provides a concurrency-safe bank with one account.
package bank

// START1 OMIT
import "sync"

var (
	// guard balance with a binary ("count only 1") semaphore
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema // release token
}

func Balance() int {
	sema <- struct{}{} // acquire token
	b := balance
	<-sema // release token
	return b
}

// END1 OMIT

// START2 OMIT
// convention: define the variables being guarded immediately after the mutex!

var (
	mu      sync.Mutex
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount // "critical section"
	mu.Unlock()
}

func Balance() int {
	mu.Lock()
	b := balance // "critical section"
	mu.Unlock()
	return b
}

// END2 OMIT

// START3 OMIT
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false // insufficient funds
	}
	return true
}

// Exported func acquires the lock
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// Unexported function requires the lock be held.
func deposit(amount int) { balance += amount }

// END3 OMIT

// START4 OMIT
var x, y int
go func() {
	x = 1 					// A1
	fmt.Print("y:", y, " ") // A2
}()

go func() {
	y = 1 					// B1
	fmt.Print("x", x, " ") 	// B2
}

// Possible outcomes
y: 0 x: 1
x: 0 y: 1
x: 1 y: 0
y: 1 x: 1

// Surprise! Also very possible.
x: 0 y: 0
y: 0 x: 0
// END4 OMIT

// START5 OMIT
var mu sync.RWMutex
var icons map[string]image.Image

// Concurrency-safe
func Icon(name string) image.Image {
	mu.RLock()
	if icons != nil {
		icon := icons[name]
		mu.RUnlock()
		return icon
	}
	mu.RUnlock()

	// acquire exclusive lock
	mu.Lock()
	if icons == nil { // NOTE: must recheck for nil!
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}
// "Correct", but clumsy and error-prone 👎🏼
// END5 OMIT

// START6 OMIT

WARNING: DATA RACE
Read by goroutine 185:
  net.(*pollServer).AddFD()
      src/net/fd_unix.go:89 +0x398
  // ...

Previous write by goroutine 184:
  net.setWriteDeadline()
      src/net/sockopt_posix.go:135 +0xdf
  // ...
// END6 OMIT