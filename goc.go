package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

type syslist struct {
	GOOS   string
	GOARCH string
}

var syslists [5]syslist

func init() {
	syslists[0] = syslist{GOOS: "darwin", GOARCH: "amd64"}
	syslists[1] = syslist{GOOS: "linux", GOARCH: "amd64"}
	syslists[2] = syslist{GOOS: "linux", GOARCH: "arm"}
	syslists[3] = syslist{GOOS: "linux", GOARCH: "arm64"}
	syslists[4] = syslist{GOOS: "windows", GOARCH: "amd64"}

}

//syslists[4] = syslist{GOOS: "darwin", GOARCH: "386"}
//syslists[5] = syslist{GOOS: "dragonfly", GOARCH: "amd64"}
//syslists[6] = syslist{GOOS: "freebsd", GOARCH: "386"}
//syslists[7] = syslist{GOOS: "freebsd", GOARCH: "amd64"}
//syslists[8] = syslist{GOOS: "freebsd", GOARCH: "arm"}
//syslists[9] = syslist{GOOS: "linux", GOARCH: "386"}
//syslists[10] = syslist{GOOS: "linux", GOARCH: "ppc64"}
//syslists[11] = syslist{GOOS: "linux", GOARCH: "ppc64le"}
//syslists[12] = syslist{GOOS: "linux", GOARCH: "mips"}
//syslists[13] = syslist{GOOS: "linux", GOARCH: "mipsle"}
//syslists[14] = syslist{GOOS: "linux", GOARCH: "mips64"}
//syslists[15] = syslist{GOOS: "linux", GOARCH: "mips64le"}
//syslists[16] = syslist{GOOS: "linux", GOARCH: "s390x"}
//syslists[17] = syslist{GOOS: "nacl", GOARCH: "386"}
//syslists[18] = syslist{GOOS: "nacl", GOARCH: "amd64p32"}
//syslists[19] = syslist{GOOS: "nacl", GOARCH: "arm"}
//syslists[20] = syslist{GOOS: "netbsd", GOARCH: "386"}
//syslists[21] = syslist{GOOS: "netbsd", GOARCH: "amd64"}
//syslists[22] = syslist{GOOS: "netbsd", GOARCH: "arm"}
//syslists[23] = syslist{GOOS: "openbsd", GOARCH: "386"}
//syslists[24] = syslist{GOOS: "openbsd", GOARCH: "amd64"}
//syslists[25] = syslist{GOOS: "openbsd", GOARCH: "arm"}
//syslists[26] = syslist{GOOS: "plan9", GOARCH: "386"}
//syslists[27] = syslist{GOOS: "plan9", GOARCH: "amd64"}
//syslists[28] = syslist{GOOS: "plan9", GOARCH: "arm"}
//syslists[29] = syslist{GOOS: "solaris", GOARCH: "amd64"}
//syslists[30] = syslist{GOOS: "windows", GOARCH: "386"}

// 编译
func main() {
	// 文件存放目录
	var parentFolder string
	// 编译输出存放的子目录
	var subFolder string = "output/"
	// 文件名前缀
	var filePrefix string
	// 要编译的源文件列表
	var files string
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	//fmt.Println(scanner.Text())
	//fmt.Println("请输文件存放目录：")
	//// 当程序到此，会停止执行等待用户输入
	//fmt.Scanln(&parentFolder)

	flag.StringVar(&parentFolder, "p", "", "文件存放目录，默认：当前目录")
	flag.StringVar(&subFolder, "s", "", "编译输出存放子目录，默认：空")
	flag.StringVar(&filePrefix, "fp", "", "创建文件名前缀，默认：空")
	flag.StringVar(&files, "fs", "", "源文件列表，默认：空")
	flag.Parse()

	cmde := path.Join(parentFolder, subFolder, filePrefix)

	//编译输出存放"bin"子目录
	cmde = `bin/` + cmde

	if filePrefix != "" && len(filePrefix) > 0 {
		cmde = cmde + "-"
	}
	for _, v := range syslists {
		ext := ""
		if v.GOOS == "windows" {
			ext = ".exe"
		}
		thisCmde := fmt.Sprintf("%v%v-%v%v", cmde, v.GOOS, v.GOARCH, ext)
		fmt.Println(thisCmde)
		cmd := exec.Command("go", "build", "-ldflags=-w", "-i", "-o", thisCmde, files)
		cmd.Env = append(os.Environ(), "GOARCH="+v.GOARCH, "GOOS="+v.GOOS)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(fmt.Sprint(err) + ": " + string(output))
			return
		}
		fmt.Println(string(output))
	}
	fmt.Println("编译完成")

	//编译完成清理缓存
	cmdf := exec.Command("go", "clean", "-cache")
	_ = cmdf.Run()

}
