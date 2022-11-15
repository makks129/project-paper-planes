package mock

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/makks129/project-paper-planes/src/utils"
)

const IMAGE = "mysql:8.0"
const MYSQL_CONTAINER_NAME = "go-test-mysql-mock"
const MYSQL_USER = "root"
const MYSQL_PASSWORD = "root"
const DB_NAME = "ppp"

func StartMysqlContainer() {
	utils.Log("Starting MySQL mock container...")
	cmd := fmt.Sprintf("docker run --rm --name %s -p 3306:3306 -e MYSQL_ROOT_PASSWORD=%s -e MYSQL_DATABASE=default -d -P %s --sql-mode=NO_ENGINE_SUBSTITUTION",
		MYSQL_CONTAINER_NAME, MYSQL_PASSWORD, IMAGE)
	res, error := execCmd(cmd)
	if error != nil {
		utils.Log("Failed to start MySQL mock container", string(res))
		panic(error)
	}
	utils.Log("Started MySQL mock container: ", string(res))
}

func WaitForDB() {
	utils.Logt("Waiting for MySQL server to get ready")
	startTime := time.Now()
	for {
		cmd := fmt.Sprintf("docker exec %s env MYSQL_PWD=%s mysqladmin ping", MYSQL_CONTAINER_NAME, MYSQL_PASSWORD)
		res, error := execCmd(cmd)
		if error != nil || !strings.Contains(res, "mysqld is alive") {
			if time.Since(startTime) > 20*time.Second {
				utils.Log("Failed to wait for MySQL server to get ready", string(res))
				panic(error)
			}
			utils.Logm(".")
			time.Sleep(1 * time.Second)
		} else {
			utils.Logm("\n")
			utils.Log("MySQL server is ready")
			return
		}
	}
}

func CreateDB() {
	command := exec.Command("docker", "exec", MYSQL_CONTAINER_NAME, "env", fmt.Sprintf("MYSQL_PWD=%s", MYSQL_PASSWORD), "mysql", "-u", "root", "-e", fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", DB_NAME))
	res, error := command.CombinedOutput()
	if error != nil {
		utils.Log("Failed to create DB", string(res))
		panic(error)
	}
	utils.Log("Created DB", DB_NAME)
}

func DropDB() {
	command := exec.Command("docker", "exec", MYSQL_CONTAINER_NAME, "env", fmt.Sprintf("MYSQL_PWD=%s", MYSQL_PASSWORD), "mysql", "-u", "root", "-e", fmt.Sprintf("DROP DATABASE IF EXISTS %s;", DB_NAME))
	res, error := command.CombinedOutput()
	if error != nil {
		utils.Log("Failed to drop DB", string(res))
		panic(error)
	}
	utils.Log("Dropped DB", DB_NAME)
}

func StopMysqlContainer() {
	utils.Log("Stopping MySQL mock container...")
	cmd := fmt.Sprintf("docker rm -f %s", MYSQL_CONTAINER_NAME)
	res, error := execCmd(cmd)
	if error != nil {
		utils.Log("Failed to stop MySQL mock container", string(res))
		panic(error)
	}
	utils.Log("Stopped MySQL mock container: ", string(res))
}

func execCmd(cmd string) (string, error) {
	cmdSplit := strings.Split(cmd, " ")
	command := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	res, error := command.CombinedOutput()
	return string(res), error
}
