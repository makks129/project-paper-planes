package task

import (
	"time"

	"github.com/makks129/project-paper-planes/src/controller"
	"github.com/makks129/project-paper-planes/src/utils"
)

// TODO cover

func RunRecurringMessageUnassignTask() {
	for {
		time.Sleep(10 * time.Minute)
		run()
	}
}

func run() {
	utils.Log("MessageUnassignTask running...")
	rows, error := controller.UnassignOldAssignedUnreadMessage()
	if error == nil {
		utils.Log("MessageUnassignTask done, rows affected: ", rows)
	} else {
		utils.Log("MessageUnassignTask error: ", error.Error())
	}
}
