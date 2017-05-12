package main

import (
	"log"
	"github.com/figoxu/Figo"
	"strings"
	"fmt"
	"bytes"
	"path/filepath"
	"os"
)
var basepath="D:/figo/workspace/workspace_figo_shine/sdz-admin/src/main/java/"

func main(){
	scanpath:="D:/figo/workspace/workspace_praticse/workspace_go/sdz-mobile-app/"
	filepath.Walk(scanpath, func(path string, fileInfo os.FileInfo, err error) error {
		if ( fileInfo == nil ) {return err}
		if fileInfo.IsDir() {return nil}
		if strings.HasSuffix(path,"go") {
			targetPath:=Figo.FilePath(path).UnixPath()
			targetPath = targetPath[0:strings.LastIndex(targetPath,"/")]
			sPath:=Figo.FilePath(scanpath).UnixPath()
			packageName:=strings.Replace(targetPath,sPath,"",-1)
			if targetPath==packageName {
				packageName=""
			}
			log.Println("path:",path)
			log.Println("targetPath:",targetPath)
			log.Println("packageName:",packageName)
			genJava(path,packageName)
		}
		return nil
	})

}

func genJava(fpath,packageName string){
	fileUtee:=Figo.FileUtee{}
	content:=fileUtee.ReadAll(fpath)
	parser:=Figo.Parser{
		PrepareReg :[]string{"type.+?struct.+?\\{[\\s\\S]+?\\}"},
		ProcessReg :[]string{"//.+","/*"},
	}
	tpStruts:=parser.Exe(content)
	log.Println(tpStruts)
	parseClassName:=func(content string)string{
		parser:=Figo.Parser{
			PrepareReg:[]string{"type.+?struct.+?\\{"},
			ProcessReg:[]string{"type","struct","\\{"," "},
		}
		return parser.Exe(content)[0]
	}
	for _,tps := range tpStruts {
		propParser:=Figo.Parser{
			PrepareReg:[]string{"\t.+"},
			ProcessReg:[]string{},
		}
		className:=parseClassName(tps)
		if strings.HasSuffix(className,"Dao") {
			continue
		}
		codes:=bytes.NewBufferString("")
		if packageName!=""{
			codes.WriteString(fmt.Sprintln("package ",packageName,";"))
		}
		codes.WriteString(fmt.Sprintln("public class ",className," {"))
		for _,v:=range propParser.Exe(tps){
			v=strings.TrimSpace(v)
			props :=strings.Split(v," ")
			if len(props) <2{
				props =strings.Split(v,"\t")
			}
			if len(props) <2{
				log.Println("err at: ",props,"  @len:",len(props))
				continue
			}
			i,typeStr,nameStr:=0,"",""
			for _,val:=range props {
				val=strings.TrimSpace(val)
				if val==""{
					continue
				}else if i<1 {
					i++
					nameStr = val
				}else if i<2 {
					i++
					typeStr = val
				}
			}
			codes.WriteString(fmt.Sprintln("\t\t",mapTpGo2Java(typeStr),"  ",nameStr,";"))
		}
		codes.WriteString(fmt.Sprintln("}"))
		targetPath:=fmt.Sprint(basepath,"/",packageName,"/",className,".java")
		Figo.FilePath(targetPath).Open()
		if packageName!=""{
			fileUtee.FlushWrite(targetPath,codes.String())
		}else{
			fileUtee.FlushWrite(fmt.Sprint(basepath,className,".java"),codes.String())
		}
	}
}

func mapTpGo2Java(tp string)string{
	if strings.HasPrefix(tp,"*") {
		tp = strings.Replace(tp,"*","",-1)
	}
	castTp := func(tpStr string)string{
		if(tpStr=="int"){
			return "Integer"
		}else if(tpStr=="int64"){
			return "Long"
		}else if(tpStr=="float32"){
			return "Float"
		}else if(tpStr=="float64"){
			return "Double"
		}else if(tpStr=="string"){
			return "String"
		}else if(tpStr=="bool"){
			return "Boolean"
		}else if (tpStr=="time.Time"){
			return "java.util.Date"
		}else if(tpStr=="uint"){
			return "Integer"
		}else if(tpStr=="StrArray"){
			return "String[]"
		}
		return tpStr
	}
	if v:=castTp(tp);v!=tp{
		return v
	}else if strings.HasPrefix(tp,"[]") {
		return fmt.Sprint(castTp(tp[2:]),"[]")
	}
	return tp
}