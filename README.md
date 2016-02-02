# substrmatchingsvc
A substring matching micro service
# example
```go
package main

import (
        "fmt"
        "github.com/blue-saber/fastmatching"
        "github.com/blue-saber/microsvc"
        "github.com/blue-saber/substrmatchingsvc"
        "github.com/blue-saber/summer"
        "github.com/gin-gonic/gin"
)

func main() {
        engine := gin.Default()

        applicationContext := summer.NewSummer()
        applicationContext.Add(engine)
        applicationContext.Add(new(microsvc.MicroService))
        applicationContext.Add(new(microsvc.MicroServiceStatus))
        applicationContext.Add(new(substrmatchingsvc.SubstrMatchingService))
        applicationContext.Add(fastmatching.NewFastMatching())

        done := applicationContext.Autowiring(func(err bool) {
                if err {
                        fmt.Println("Failed to autowiring.")
                } else {
                        fmt.Println("Autowired.")
                }
        })

        if err := <-done; !err {
                engine.Run("127.0.0.1:8088")
        }
}
```
