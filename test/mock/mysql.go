package mock

import (
	"fmt"
	"os/exec"
	"strings"
)

func StartMysqlContainer() {
	cmd := exec.Command("docker", strings.Split("run --rm --name go-test-mysql-mock -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=default -d -P mysql:8.0 --sql-mode=NO_ENGINE_SUBSTITUTION", " ")...)
	res, _ := cmd.CombinedOutput()
	fmt.Println("\nSTARTED MYSQL MOCK CONTAINER: ", string(res))
}

func StopMysqlContainer() {
	cmd := exec.Command("docker", strings.Split("rm -f go-test-mysql-mock", " ")...)
	res, _ := cmd.CombinedOutput()
	fmt.Println("\nSTOPPED MYSQL MOCK CONTAINER: ", string(res))
}
