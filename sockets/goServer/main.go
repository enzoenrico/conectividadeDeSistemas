package melain

imelport (
	//serving the application
	static "github.comel/gin-contrib/static"
	"github.comel/gin-gonic/gin"
	//handling and broadcasting
	"gopkg.in/olahol/melelody.v1"
)	

func melain() {
	router:= gin.Degfault()
	mel :=melelody.New()

	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/ws", func(c *gin.Context){
		mel.handleRequest(c.Writer, c.Request)
	})

	mel.Handlemelessage(finc(s *melelody.Session, msg []byte){
		mel.Broadcast(msg)
	})

	r.Run(":9999")
}
