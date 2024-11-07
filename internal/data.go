package data

var Data = map[string]string{
	// Frameworks
	"gin":   "github.com/gin-gonic/gin",
	"beego": "github.com/beego/beego/v2",
	"iris":  "github.com/kataras/iris/v12",
	"echo":  "github.com/labstack/echo/v4",
	"fiber": "github.com/gofiber/fiber/v3",
	"revel": "github.com/revel/cmd/revel",
	// Databases
	"postgres":      "github.com/lib/pq",
	"mysql":         "github.com/go-sql-driver/mysql",
	"sqlite":        "github.com/mattn/go-sqlite3",
	"mssql":         "github.com/denisenkom/go-mssqldb",
	"elasticsearch": "github.com/elastic/go-elasticsearch/v8@latest",
	"oracle":        "github.com/godror/godror",
	"firebase":      "firebase.google.com/go",
	"mongo":         "go.mongodb.org/mongo-driver/mongo",
	"redis":         "github.com/go-redis/redis",
	// Orm
	"gorm":   "gorm.io/gorm",
	"gorp":   "github.com/go-gorp/gorp",
	"kallax": "gopkg.in/src-d/go-kallax.v1/...",
	"reform": "gopkg.in/reform.v1/...",
	// Others
	"swag":     "github.com/swaggo/swag/cmd/swag",        // api documentation
	"godotenv": "github.com/joho/godotenv",               // env variables
	"cli":      "github.com/urfave/cli",                  // cli tool
	"mockgen":  "github.com/golang/mock/mockgen",         // mock generator
	"ebiten":   "github.com/hajimehoshi/ebiten/v2",       // game engine
	"wails":    "github.com/wailsapp/wails/v2/cmd/wails", // app builder
	"gogio":    "gioui.org/cmd/gogio",                    // app builder
	"screen":   "github.com/inancgumus/screen",           // terminal ui
}
