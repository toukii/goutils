package goutils


type UploadHandler struct {}

func (m *UploadHandler)ServeHTTP(w http.ResponseWriter, req *http.Request)  {
	uri:=req.RequestURI
	if req.Method=="POST" {
		if "/upload"==uri {
			m.Upload(w,req)
		}else if "/streamUpload"==uri{
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
	if goutils.CheckErr(derr) {
		fmt.Fprint(w, derr)
		return
	}
	gerr:=goutils.WriteFile("upload.png",bs)
	if goutils.CheckErr(gerr) {
		fmt.Fprint(w,gerr)
		return
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}

func (m *UploadHandler) StreamUpload(w http.ResponseWriter, req * http.Request)  {
	enc:=base64.StdEncoding
	file,err:=os.OpenFile("page.png",os.O_CREATE|os.O_WRONLY,0644)
	defer file.Close()
	if goutils.CheckErr(err) {
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
		if goutils.CheckErr(err1) {
			break
		}
	}

	dst:=make([]byte,len(allbytes)*2)
	n3,err:=enc.Decode(dst,allbytes)
	if !goutils.CheckErr(err) {
		file.Write(dst[:n3])
	}
	fmt.Fprint(w,"success")
	fmt.Println("cost:",time.Now().Sub(start))
}