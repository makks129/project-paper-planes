package mock

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func StartMysqlContainer() {
	res, error := execCmd("docker run --rm --name go-test-mysql-mock -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=default -d -P mysql:8.0 --sql-mode=NO_ENGINE_SUBSTITUTION")
	if error != nil {
		fmt.Println("Failed to start MySQL mock container", string(res))
		panic(error)
	}
	fmt.Println("\n🚀 Started MySQL mock container: ", string(res))
}

func WaitForDB() {
	startTime := time.Now()
	fmt.Println("⏳ Waiting for MySQL to get ready...")
	for {
		res, error := execCmd("docker exec go-test-mysql-mock env MYSQL_PWD=root mysqladmin ping")
		if error != nil || !strings.Contains(res, "mysqld is alive") {
			if time.Since(startTime) > 20*time.Second {
				fmt.Println("Failed to wait for MySQL to get ready", string(res))
				panic(error)
			}
			time.Sleep(1 * time.Second)
		} else {
			fmt.Println("🛢  MySQL is ready")
			return
		}
	}
}

func CreateDB() {
	command := exec.Command("docker", "exec", "go-test-mysql-mock", "env", "MYSQL_PWD=root", "mysql", "-u", "root", "-e", "CREATE DATABASE IF NOT EXISTS ppp;")
	res, error := command.CombinedOutput()
	if error != nil {
		fmt.Println("Failed to create DB", string(res))
		panic(error)
	}
	fmt.Println("✅ Created DB ppp\n ")
}

func DropDB() {
	command := exec.Command("docker", "exec", "go-test-mysql-mock", "env", "MYSQL_PWD=root", "mysql", "-u", "root", "-e", "DROP DATABASE IF EXISTS ppp;")
	res, error := command.CombinedOutput()
	if error != nil {
		fmt.Println("Failed to drop DB", string(res))
		panic(error)
	}
	fmt.Println("\n💣 Dropped DB ppp")
}

func StopMysqlContainer() {
	res, error := execCmd("docker rm -f go-test-mysql-mock")
	if error != nil {
		fmt.Println("Failed to stop MySQL mock container", string(res))
		panic(error)
	}
	fmt.Println("🏁 Stopped MySQL mock container: ", string(res))
}

func execCmd(cmd string) (string, error) {
	cmdSplit := strings.Split(cmd, " ")
	command := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	res, error := command.CombinedOutput()
	return string(res), error
}
