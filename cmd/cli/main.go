package main

import (
	"encoding/json"
	"fmt"
	"os"
	"syshealthcli/pkg/internals"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)


func checkUsage(cpuUsage float64,memUsage float64,logFile *os.File){
  cpuPercent, err := cpu.Percent(0, false)
        if err != nil {
            fmt.Println("Error getting CPU usage:", err)
            return
        }

        // Get memory usage
        memInfo, err := mem.VirtualMemory()
        if err != nil {
            fmt.Println("Error getting memory info:", err)
            return
        }

				usageData:=internals.CreateUsageData(cpuPercent[0],memInfo.UsedPercent,
					time.Now().Format(time.RFC3339),

				
				)


				        // Log data in JSON format
        logUsage(logFile, usageData)


				        if cpuPercent[0] > cpuUsage {
									fmt.Printf("cpu usage %v",cpuUsage)
            notify("CPU usage exceeded", fmt.Sprintf("Current CPU usage: %.2f%%", cpuPercent[0]))
        }

        if memInfo.UsedPercent > memUsage {
									fmt.Printf("mem usage %v",memUsage)

            notify("Memory usage exceeded", fmt.Sprintf("Current Memory usage: %.2f%%", memInfo.UsedPercent))
        }

        time.Sleep(5 * time.Second) // Check every 5 seconds
    


	  

}

func notify(title, message string) {
    // uses `osascript` for macOS notifications
    // cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, message, title))
    // err := cmd.Run()
    // if err != nil {
    //     fmt.Println("Error sending notification:", err)
    // }

		 err := beeep.Alert(title, message, "")
    if err != nil {
        fmt.Println("Error sending alert:", err)
    }
}




func logUsage(logFile *os.File,data internals.UseageData){
	      logEntry, err := json.Marshal(data)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }
    logFile.WriteString(string(logEntry) + "\n")

}


func main(){

	opt,err:=internals.AddOptions()

	if err!=nil{
		fmt.Printf("Error: %s\n",err)
		return
	}

	// creating a file and save the usage in this file

	logFile, err := os.OpenFile("usage_log.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %s\n",err)
		return
	}
	defer logFile.Close()





for {
		checkUsage(opt.Cpu, opt.Mem, logFile)
	}



}