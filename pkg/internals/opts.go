package internals

import "github.com/hellflame/argparse"


type Options struct{
	Cpu float64
	Mem float64

	
}

func AddOptions()(*Options,error){

	 parser:=argparse.NewParser("sysinfo", "Get system information",nil)


	 cpu:=parser.Float("c", "cpu", &argparse.Option{Help: "CPU usage",
	  Default: "0.0",
	},)

	 mem:=parser.Float("m", "mem", &argparse.Option{Help: "Memory usage",
	  Default: "0.0",},

	)

	err:=parser.Parse(nil)

	if err!=nil{
		return nil,err
	}

	return &Options{
		Cpu: *cpu,
		Mem: *mem,

	},nil
	
}