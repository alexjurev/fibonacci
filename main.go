package main
 
import (
	"fibo/internal/app/apiserver"
	"fibo/pkg"
	"github.com/gin-gonic/gin"
)



func main() {
	router := gin.Default()
	router.POST("/employee", handler.FiboCounter)
	pkg.FibSlice(0, 11)

}