package monitor

import (
	"github.com/kristenfelch/go-present/demo"
	"log"
	"time"
)

func Monitor(name string) demo.Decorator {
	return func(r demo.Runner) demo.Runner {
		return demo.RunnerFunc(func() {
			defer timeTrack(time.Now(), name)
			r.Run()
			return
		})
	}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
