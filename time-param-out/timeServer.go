package main

import "fmt"
import "time"


type Weekday int
const (
        Sunday Weekday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
)

func main()  {
    p := fmt.Println
   // now := time.Now()
    now := time.Date(
        2018, 12, 1, 00, 00, 0, 0, time.UTC)
    p("Time now:",now)
     then := time.Date(
        2019, 1, 1, 00, 00, 0, 0, time.UTC)
    var N  int = 365
    if(then.Year() % 4 == 0) {
     N = 366
    }
    p("Day today:",now.Weekday())
    p("No of days in 2018:",N)
    diff := then.Sub(now)
    var L int = int(1+diff.Hours()/24)
    p("No of days left in 2018:",L)
    day := now.Weekday()
    count := 0
    if (day == time.Monday) {
     count = 0
    } else if (day == time.Tuesday) {
     count = 1
    } else if (day ==time.Wednesday) {
     count = 2
    } else if (day == time.Thursday) {
     count = 3
    } else if (day == time. Friday) {
     count = 4
   } else if (day == time.Saturday) {
    count = 5
   } else {
    count = 6
   }
     p("count:",count)
    sat := (L-count+5)/7
    mod_sat := (L-count+5)%7
    sun := (L-count+6)/7   
    mod_sun := (L-count+6)%7
    p(mod_sat)
    p("Number of Saturdays:",sat)
    p(mod_sun)
    p("Number of Sundays:",sun) 
      
}
