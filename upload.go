package goutils
import (
	"net/http"
	"fmt"
	"encoding/base64"
	"os"
	"strings"
	"bufio"
	"time"
)

type UploadHandler struct {}

func (m *UploadHandler)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	uri:=req.RequestURI
	fmt.Println("URI:",uri)
	if req.Method=="POST" {

		if strings.HasSuffix(uri,"/upload") {
			m.Upload(w,req)
		}else if strings.HasSuffix(uri,"/streamUpload") {
			m.StreamUpload(w,req)
		}
	}
}

func (m *UploadHandler) Upload(w http.ResponseWriter, req * http.Request)  {
	start:=time.Now()
	req.ParseForm()
	data:=req.FormValue("imgdata")
	b64data := data[strings.IndexByte(data, ',')+1:]
	enc:=base64.StdEncoding
	bs,derr:=enc.DecodeString(b64data)
	if CheckErr(derr) {
		fmt.Fprint(w, derr)
		return
	}
	gerr:=WriteFile("MDFs/page.png",bs)
	if CheckErr(gerr) {
		fmt.Fprint(w,gerr)
		return
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}

func (m *UploadHandler) StreamUpload(w http.ResponseWriter, req * http.Request)  {
	enc:=base64.StdEncoding
	file,err:=os.OpenFile("MDFs/page.png",os.O_CREATE|os.O_WRONLY,0644)
	defer file.Close()
	if CheckErr(err) {
		fmt.Fprint(w, err)
		return
	}

	try:=make([]byte,120)
	peekR:=bufio.NewReader(req.Body)
	peekR.ReadBytes(',')

	start:=time.Now()
	allbytes:=make([]byte,0,req.ContentLength)
	for{
		n1,err1:=peekR.Read(try)
		if n1>0 {
			allbytes = append(allbytes,try[:n1]...)
		}
		if CheckErr(err1) {
			break
		}
	}

	dst:=make([]byte,len(allbytes)*2)
	n3,err:=enc.Decode(dst,allbytes)
	if !CheckErr(err) {
		file.Write(dst[:n3])
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}