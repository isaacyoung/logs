
// cd src/isaac
// go run logs.go

package main

import (
	"os"
	"bufio"
	"net/http"
	"log"
	"strings"
)

func readLog(w http.ResponseWriter, r *http.Request) {
	//fileName := "F:/log/cdsq-error.log"
	fileName := "/data/logs/cdsq-manage/cdsq-error.log"
	f, err := os.Open(fileName)
	if err != nil {
		return
	}
	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		line = strings.Replace(line,"\t","&emsp;&emsp;", -1)
		if strings.Contains(line, "ERROR") || strings.Contains(line, "com.cdsq.manage") {
			line = "<p style='color:red;'>" + line + "</p>"
		} else {
			line = "<p>" + line + "</p>"
		}

		w.Write([]byte(line))
		if err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/", readLog) //注册URI路径与相应的处理函数
	er := http.ListenAndServe(":8083", nil)  // 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
	if er != nil {
		log.Fatal("ListenAndServe: ", er)
	}
}