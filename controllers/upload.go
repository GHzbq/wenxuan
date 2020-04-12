package controllers

//const (
//	ImgBasePath = "/static/www/"
//)
//
// HandleNcGetUploadPicture 获取上传图片页面
func (c *MainController) HandleNcGetUploadPicture() {
	c.TplName = "upload_picture.html"
}
//
//// HandleNcPostUploadPicture 处理上传图片
//func (c *MainController) HandleNcPostUploadPicture() {
//	f, h, e := c.GetFile("filename")
//	if e != nil {
//		logs.Error("c.GetFile failed, error = %v", e.Error())
//		c.Ctx.WriteString("get file failed")
//		return
//	}
//	defer f.Close()
//
//	// 1. 要限定格式
//	fileExt := path.Ext(h.Filename)
//	if fileExt != ".jpg" && fileExt != ".png" {
//		logs.Error("format not supported")
//		c.Ctx.WriteString("format not supported")
//		return
//	}
//	// 2. 限制大小
//	if h.Size > fileSize {
//		logs.Error("file is too big")
//		c.Ctx.WriteString("file is too big")
//		return
//	}
//	// 3. 对文件重命名 避免重名覆盖
//	count := atomic.AddInt64(&atomicCount, 1)
//	fileName := strconv.FormatInt(time.Now().Unix(), 10)
//	fileName += "_" + strconv.FormatInt(count, 10) + fileExt
//	filePath := "." + ImgBasePath + fileName
//	//fileURL := "/get_picture?id=" + fileName
//	fileURL := ImgBasePath + "?id=" + fileName
//	logs.Debug("filePath = %v, fileURL = %v", filePath, fileURL)
//	e = c.SaveToFile("filename", filePath)
//	if e != nil {
//		logs.Error("c.SaveToFile failed, error = %v", e.Error())
//		c.Ctx.WriteString("save file failed")
//		return
//	} else {
//		logs.Debug("save file succeed, filePath = %v, fileURL = %v", filePath, fileURL)
//		c.Ctx.WriteString(fileURL)
//		return
//	}
//}
//
////func (c *MainController) HandleNcPostUploadPictureBase64() {
////	pictures := c.GetStrings("picture_name")
////	if pictures == nil {
////		c.Redirect("/", 302)
////		return
////	}
////	pictureBase64 := pictures[0]
////	for i := 1; i < len(pictures); i++ {
////		pictureBase64 += " "
////		pictureBase64 += pictures[i]
////	}
////	if pictureBase64 == "" {
////		c.Redirect("/", 302)
////		return
////	}
////	logs.Debug("pictureBase64=%v", pictureBase64)
////	count := atomic.AddInt64(&atomicCount, 1)
////	fileName := strconv.FormatInt(time.Now().Unix(), 10)
////	fileName += "_" + strconv.FormatInt(count, 10)
////	picturePath := "/home/work/project/wenxuan/img/" + fileName
////	pictureURL := "/get_picture_base64?id=" + fileName
////
////	// file, e := os.OpenFile(picturePath, os.O_CREATE|os.O_WRONLY , 0755)
////	file, e := os.Create(picturePath)
////	if e != nil {
////		logs.Error("os.Create failed, error = %v", e.Error())
////		return
////	}
////	defer file.Close()    //关闭文件
////
////	_, e = file.WriteString(pictureBase64)
////	if e != nil {
////		logs.Error("file.WriteString failed, error = %v", e.Error())
////		return
////	}
////
////	c.Ctx.ResponseWriter.Write([]byte(pictureURL))
////}
//
//// HandleNcGetGetPicture 获取图片
//func (c *MainController) HandleNcGetGetPicture() {
//	fileName := c.GetString("id")
//	if fileName == "" {
//		logs.Error("id is unset")
//		c.Ctx.WriteString("id is unset")
//		return
//	}
//	filePath := "." + ImgBasePath + fileName
//	logs.Debug("filePath = %v", filePath)
//	c.Ctx.
//
//	//file, e := os.OpenFile(filePath, os.O_RDONLY, 0666)
//	//if e != nil {
//	//	logs.Error("os.OpenFile failed, fileName = %v, error = %v", fileName, e.Error())
//	//	c.Ctx.WriteString("open file failed")
//	//	return
//	//}
//	//defer file.Close()
//
//
//
//	//buf, err := ioutil.ReadAll(file)
//	//if err != nil {
//	//	logs.Error("ioutil.ReadAll failed, error = %v", e.Error())
//	//	return
//	//}
//	//
//	//logs.Debug("buf = %v", string(buf))
//	//c.Ctx.ResponseWriter.Write(buf)
//}