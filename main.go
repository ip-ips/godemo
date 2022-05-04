package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"

	//"godemo/services"
	"time"
	//"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		// time.Sleep(time.Millisecond * 100)
	}

}

//通道 在协程之间传递数据
var values = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	time.Sleep(time.Second * 5)
	values <- value
}

var wg sync.WaitGroup

func showmsg(i int) {
	defer wg.Done()
	//协程结束
	fmt.Printf("i: %v\n", i)
}

func show1() {
	for i := 0; i < 7; i++ {

		if i > 5 {
			runtime.Goexit()
		}
		fmt.Println(i)
	}
}

var c = make(chan int)

var chanInt = make(chan int, 0)
var chanStr = make(chan string)

func createfile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f: %v\n", f)
	}
}

func makeDir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err2 != nil {
		fmt.Printf("err: %v\n", err2)
	}
}

func remove() {
	err := os.Remove("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	err2 := os.RemoveAll("a")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

func getWd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("dir: %v\n", dir)
	}
	os.Chdir("d:/") //修改工作目录

	s := os.TempDir()
	fmt.Printf("s: %v\n", s) //临时目录

}

func renume() {
	err := os.Rename("text.txt", "text2.txt") //重命名
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readfile() {
	b, _ := os.ReadFile("a.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

func writefile() {
	//写会覆盖原来的文件内容
	os.WriteFile("a.txt", []byte("helo"), os.ModePerm)
	//追加
	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
	f.Write([]byte(" hr golong"))
	f.WriteString("tgtftytfjt")
	f.Close()

}

func testCopy() {
	r := strings.NewReader("hello world")

	_, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
}

func check(s string) (string, error) {
	if s == "" {
		err := errors.New("字符串不能为空")
		return "", err
	} else {
		return s, nil
	}
}
func test1() {
	/* t:=time.Now()
	fmt.Printf("t: %v\n", t)
	y:=t.Year()
	m1:=t.Month()
	d:=t.Day()
	h:=t.Hour()
	m:=t.Minute()
	s:=t.Second() */

}

type Person struct {
	Name string
	Age  int
}

// var db *sql.DB

// func initDB() (err error) {
// 	dbn := "root:@tcp(127.0.0.1:3306)/cart"
// 	//?charset=utf8&parseTime=True

// 	db, err := sql.Open("mysql", dbn)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	return nil

// }

// func queryOne(id int) {

// 	var u1 User
// 	// 1.写查询单条记录的sql语句
// 	sqlStr := "select id,username,password,phone,address from t_user where id =?;"

// 	//2.执行并拿到结果
// 	//必须对row对象调用scan方法，该方法会释放数据库连接
// 	fmt.Println("ssss")
// 	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.username, &u1.password, &u1.phone, &u1.address) //从连接池拿一个连接出去数据库查询单条记录
// 	fmt.Println("www")
// 	//打印结果
// 	fmt.Println(u1)
// }

/* func insertData() {

	s := "insert into t_user (username,password,phone,address) values (?,?,?,?)"
	strTmt, err := utils.Db.Prepare(s)
	if err != nil {
		return err
	}
	//3.执行
	//Exec执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。参数args表示query中的占位参数。
	_, err = strTmt.Exec("zhangsan", "123121", "15376689901", "zjsuz") //id自动增长
	if err != nil {
		return err
	}

		db = new(sql.DB)
	   	r, err := db.Exec(s, "zhangsan", "123121", "15376689901", "zjsuz")
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   		return
	   	}
	   	i, err2 := r.LastInsertId()
	   	if err2 != nil {
	   		fmt.Printf("err: %v\n", err)
	   		return
	   	}
	   	fmt.Printf("i: %v\n", i)
} */

type User struct {
	id       int
	username string
	password string
	phone    string
	address  string
}

// func queryData() {
// 	s := "select * from t_user where id = ?"
// 	var u User

// 	err := db.QueryRow(s, 1).Scan(&u.id, &u.username, &u.password, &u.phone, &u.address)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("u: %v\n", u)
// 	}

// }

// func main() {

// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Println("连接成功！")
// 	}
// 	queryOne(1)
/* fmt.Printf("db: %v\n", db)
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10) */

/* 	b1 := []byte(`{"Name":"lili","Age":20}`)
   	var p map[string]interface{}
   	json.Unmarshal(b1, &p) //把json对象转换成结构体
   	fmt.Printf("p: %v\n", p)

   	json.NewDecoder()
   	math.Sqrt()
   	for k, v := range p {
   		fmt.Printf("k: %v\n", k)
   		fmt.Printf("v: %v\n", v)
   	} */

/* p := Person{
	Name: "lili",
	Age:  20,
}
b, _ := json.Marshal(p) //把结构体转换成json对象，返回值为比特字节
fmt.Printf("string(b): %v\n", string(b))
fmt.Printf("string(b): %v\n", b)
fmt.Printf("string(b): %T\n", b) */

/* s := []int{2, 4, 2, 1, 3}
s1 := []float64{1.1, 3.2, 3.1, 2.1}
sort.Ints(s) //对切片进行排序
sort.Float64s(s1)
ok := sort.IntsAreSorted(s) //判断切片是否排序过，返回一个布尔值
fmt.Printf("s: %v\n", s)
fmt.Printf("ok: %v\n", ok) */

/* s, err := check("")
if err != nil {
	fmt.Printf("err.Error(): %v\n", err.Error())
} else {
	fmt.Printf("s: %v\n", s)
} */

/* var s string = "hello world"
b1:=[]byte{1,2,2}
s =string(b1)
b1 = []byte(s) */

// readfile()
// writefile()
//testCopy()
/* name := "lili"
age := 29
log.Print("my log")
log.Panic("my log")
log.Printf("my log %d", 100)
log.Println(name, "", age) */
//fmt.Printf("os.Getgid(): %v\n", os.Getgid())
/* fmt.Printf("os.Environ(): %v\n", os.Environ())
fmt.Printf("os.Getenv(\"GOPATH\"): %v\n", os.Getenv("GOPATH"))
fmt.Printf("os.Getenv(\"JAVA_HOME\"): %v\n", os.Getenv("JAVA_HOME")) */

/* ticker := time.NewTicker(time.Second)
counter := 1
for _ = range ticker.C {
	fmt.Println("ticker----------")
	counter++
	if counter >= 5 {
		ticker.Stop()
		break
	}
}
chanInt1 := make(chan int)

go func() {
	for _ = range ticker.C {
		select {

		//在通道里写
		case chanInt1 <- 1:
		case chanInt1 <- 2:
		case chanInt1 <- 3:
		}

	}
}()

sum := 0
//读通道
for v := range chanInt1 {
	sum += v
	fmt.Println("shoudao:", v)
	if sum >= 10 {
		fmt.Printf("sum: %v\n", sum)
		break
	}
}*/

/* timer1 := time.NewTimer(time.Second * 3)
fmt.Printf("time.Now(): %v\n", time.Now())
timer2 := <-timer1.C //阻塞的，指定的时间到了
fmt.Printf("timer2: %v\n", timer2) */

/* t1 := time.NewTimer(time.Second)
//timer只能执行一次
go func() {
	<-t1.C
	fmt.Println("func--------")
}()
s := t1.Stop()
if s {
	fmt.Println("Stop---------")
}

time.Sleep(time.Second * 3)
fmt.Println("main end ----------") */

/* go func() {
	chanInt <- 100
	chanStr <- "hello"

	defer close(chanInt)
	defer close(chanStr)
}()
for {
	select {
	case r := <-chanInt:
		fmt.Println("r:", r)
	case r := <-chanStr:
		fmt.Println("r:", r)
	default:
		fmt.Println("default-----")
	}
	time.Sleep(time.Second)
} */

/* go func() {
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
}()

r := <-c
fmt.Printf("r: %v\n", r)
r = <-c
fmt.Printf("r: %v\n", r)
// r = <-c
// fmt.Printf("r: %v\n", r)
// r = <-c
// fmt.Printf("r: %v\n", r)
for v := range c {
	fmt.Printf("v: %v\n", v)

} */

/* //runtime.GOMAXPROCS(1)
go show1()
go show("sss")
time.Sleep(time.Second)
fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU()) */

/* go show1("xiecheng")
for i := 0; i < 3; i++ {
	//runtime.Gosched()
	fmt.Println("golang")
} */

/* for i := 0; i < 10; i++ {
	//启动一个协程
	go showmsg(i)
	wg.Add(1)
}
wg.Wait()
//主协程等待group里的任务执行完后执行
fmt.Println("end.....") */

/* defer close(values)
go send()
fmt.Println("wait...........")
value := <-values
fmt.Printf("receive: %v\n", value)
fmt.Println("end..............") */

/* go show("Java")
go show("gogogo")
time.Sleep(time.Millisecond * 2000)
fmt.Println("main end-----------") */

/* fmt.Println("hello word!")
	services.TestUserServic()
	services.TestCustomer()
	//Default返回一个默认的路由引擎
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        //输出json结果给调用方
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080 */

// }
