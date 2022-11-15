package mock

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

const IMAGE = "mysql:8.0"
const MYSQL_CONTAINER_NAME = "go-test-mysql-mock"
const MYSQL_USER = "root"
const MYSQL_PASSWORD = "root"
const DB_NAME = "ppp"

func StartMysqlContainer() {
	cmd := fmt.Sprintf("docker run --rm --name %s -e MYSQL_ROOT_PASSWORD=%s -e MYSQL_DATABASE=default -d -P %s --sql-mode=NO_ENGINE_SUBSTITUTION",
		MYSQL_CONTAINER_NAME, MYSQL_PASSWORD, IMAGE)
	res, error := execCmd(cmd)
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
		cmd := fmt.Sprintf("docker exec %s env MYSQL_PWD=%s mysqladmin ping", MYSQL_CONTAINER_NAME, MYSQL_PASSWORD)
		res, error := execCmd(cmd)
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
	command := exec.Command("docker", "exec", MYSQL_CONTAINER_NAME, "env", fmt.Sprintf("MYSQL_PWD=%s", MYSQL_PASSWORD), "mysql", "-u", "root", "-e", fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", DB_NAME))
	res, error := command.CombinedOutput()
	if error != nil {
		fmt.Println("Failed to create DB", string(res))
		panic(error)
	}
	fmt.Println("✅ Created DB ppp\n ")
}

func DropDB() {
	command := exec.Command("docker", "exec", MYSQL_CONTAINER_NAME, "env", fmt.Sprintf("MYSQL_PWD=%s", MYSQL_PASSWORD), "mysql", "-u", "root", "-e", fmt.Sprintf("DROP DATABASE IF EXISTS %s;", DB_NAME))
	res, error := command.CombinedOutput()
	if error != nil {
		fmt.Println("Failed to drop DB", string(res))
		panic(error)
	}
	fmt.Println("\n💣 Dropped DB ppp")
}

func StopMysqlContainer() {
	cmd := fmt.Sprintf("docker rm -f %s", MYSQL_CONTAINER_NAME)
	res, error := execCmd(cmd)
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
