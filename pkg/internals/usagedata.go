package internals

type UseageData struct{
	  CPUUsage    float64 `json:"cpu_usage"`
    MemoryUsage float64 `json:"memory_usage"`
    Timestamp   string  `json:"timestamp"`

}

func CreateUsageData(cpuUsage,memoryUsage float64,timeStamp string)(UseageData){

	return UseageData{
		 CPUUsage: cpuUsage,
		 MemoryUsage: memoryUsage,
		 Timestamp: timeStamp,
	}


}