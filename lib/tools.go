package kilib

import (
        "strings"
)


// Read tools switch information
func GetToolSwitch(currentDir string, logName string, mode string) []string {
        toolSwitch := []string{"on", "on", "on"}
        toolSwitchData, _ := ReadFile(currentDir + "/data/config/tools.txt")
        if toolSwitchData != "" {
                toolSwitch = strings.Split(toolSwitchData, ",")
        }
        return toolSwitch
}

// Switch tools status
func SetToolSwitch(sshTool string, installTool string, calendarTool string, currentDir string, logName string, mode string) error {
        toolswitch := sshTool + "," + installTool + "," + calendarTool
        err := DatabaseUpdate(currentDir+"/data/config/tools.txt", toolswitch, currentDir, logName, mode)
        return err
}




