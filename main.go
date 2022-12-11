package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"luxamrown/go-pcinfo-fetch/structs"
	"luxamrown/go-pcinfo-fetch/utils"
	"os/exec"
	"strings"
)

func GetCMDOutput(command string, arg ...string) string {
	var out bytes.Buffer
	var nextCommand []string
	var cmd *exec.Cmd

	if arg[0] == "|" {
		nextCommand = append(nextCommand, arg[1:]...)
		cmd = exec.Command(command)
		findCmd := exec.Command(nextCommand[0], nextCommand[1:]...)
		reader, writer := io.Pipe()
		cmd.Stdout = writer
		findCmd.Stdin = reader
		findCmd.Stdout = &out
		cmd.Start()
		findCmd.Start()
		cmd.Wait()
		writer.Close()
		findCmd.Wait()
		reader.Close()
		return out.String()
	}
	cmd = exec.Command(command, arg...)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return out.String()
}

func GetGPUName() string {
	var outputArray []string
	output := GetCMDOutput("wmic", "path", "win32_VideoController", "get", "name")
	outputArray = strings.Split(output, " ")
	outputArray = utils.SimplifyOutput(outputArray)
	return strings.Join(outputArray[2:], " ")
}

func GetCPUName() string {
	var outputArray []string
	output := GetCMDOutput("wmic", "cpu", "get", "name")
	outputArray = strings.Split(output, " ")
	outputArray = utils.SimplifyOutput(outputArray)
	return strings.Join(outputArray[2:], " ")
}

func GetHostAndOsname() (string, string) {
	var outputHost []string
	var outputOs []string
	// var outputArray []string

	command := []string{"/C:Host Name", "/C:OS Name"}
	for idx, elem := range command {
		output := GetCMDOutput("systeminfo", "|", "findstr", elem)
		if idx == 0 {
			outputHost = strings.Split(output, " ")
			outputHost = utils.SimplifyOutput(outputHost)
		} else {
			outputOs = strings.Split(output, " ")
			outputOs = utils.SimplifyOutput(outputOs)
		}
	}

	// outputArray = append(outputHost, outputOs...)
	return strings.Join(outputHost[2:], " "), strings.Join(outputOs[2:], " ")
	// output := GetCMDOutput("systeminfo", "|", "findstr", "/C:Host Name")
	// outputArray = strings.Split(output, " ")
	// outputArray = SimplifyOutput(outputArray)
	// return outputArray
}

func GetPcInfo() structs.PcInfo {
	hostName, osName := GetHostAndOsname()
	cpuName := GetCPUName()
	gpuName := GetGPUName()
	return structs.NewPcInfo(structs.NewSysInfo(hostName, osName, cpuName), structs.NewGpuInfo(gpuName))
}

func main() {
	PcInfo := GetPcInfo()
	fmt.Println("========================================================")
	fmt.Printf("OS: %s\n", PcInfo.OsName)
	fmt.Printf("Host: %s\n", PcInfo.HostName)
	fmt.Printf("CPU: %s\n", PcInfo.CPU)
	fmt.Printf("GPU: %s\n", PcInfo.GPU)
	fmt.Println("========================================================")

}
