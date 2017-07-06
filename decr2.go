package main

func decr(n uint64 , c chan bool) {
    for n > 0 {
        n--
    }
    c <- true //koniec wykonywania goroutyny
}

func main() {
    c := make (chan bool, 4)
    go decr(10000000000, c) // wszystkie
    go decr(10000000000, c) // goroutyny piszą
    go decr(10000000000, c) // do jednego
    go decr(10000000000, c) // kanału
    <-c //z kanału
    <-c //wyciągnijmy
    <-c //cztery
    <-c // wartości
} //i zakończmy program (co jest równoznaczne z zakończeniem wykonywania funkcji main)
