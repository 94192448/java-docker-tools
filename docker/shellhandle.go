package docker

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func ExecShell(writer http.ResponseWriter, request *http.Request) {
	var s = "uname"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}

func BuildDockerImages(writer http.ResponseWriter, request *http.Request) bool {
	command := `./jar-to-docker.sh .`

	cmd := exec.Command("/bin/bash", "-c", command)

	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()

	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	cmd.Wait()
	return true

	/*content, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(content)) //输出ls命令查看到的内容*/

	/*output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))*/

}
