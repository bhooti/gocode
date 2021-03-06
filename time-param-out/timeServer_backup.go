package main

import "fmt"
import "time"


func main()  {
    p := fmt.Println
    now := time.Now()
    p(now)
     then := time.Date(
        2019, 1, 1, 00, 00, 0, 0, time.UTC)
    var N  int = 365
    if(then.Year() % 4 == 0) {
     N = 366
    }
    p(N)
    diff := then.Sub(now)
    var L int = int(1+diff.Hours()/24)
    p(L)
    count :=0
    for temp :=now; temp<then; temp++ {     
     if(IsWeekend(temp)) {
       count ++
     }  
    }
   p(count)
}

func IsWeekend(date time.Time) bool {
	day := date.Weekday()
	return day == time.Saturday || day == time.Sunday
}
