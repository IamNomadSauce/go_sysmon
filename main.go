package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

// GPUData struct to hold all the GPU information
type GPUData struct {
	Temperature          float64 `json:"temperature"`
	PowerConsumption     float64 `json:"powerConsumption"`
	CurrentGPUClockState string  `json:"currentGPUClockState"`
	VramUsedPercentage   float64 `json:"vramUsedPercentage"`
	GPUUtilization       int     `json:"gpuUtilization"`
	CurrentGPUClockSpeed int     `json:"currentGPUClockSpeed"`
	Sclk                 int     `json:"sclk"`
	Mclk                 int     `json:"mclk"`
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()
	for {

		// GPU Temperature
		tempFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/temp1_input"
		tempStr, err := ioutil.ReadFile(tempFilePath)
		if err != nil {
			fmt.Println("Error reading temperature:", err)
			return
		}

		// The temperature is usually reported in millidegrees Celsius; convert to degrees
		tempMilliDegrees := strings.TrimSpace(string(tempStr))
		tempDegrees, err := strconv.Atoi(tempMilliDegrees)
		if err != nil {
			fmt.Println("Error converting temperature:", err)
			return
		}
		// fmt.Printf("GPU Temperature: %.2f°C\n", float64(tempDegrees)/1000.0)

		// Average Power Consumption
		powerFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/power1_average"
		powerStr, err := ioutil.ReadFile(powerFilePath)
		if err != nil {
			fmt.Println("Error reading power consumption:", err)
			return
		}

		// The power consumption is usually reported in microwatts; convert to milliwatts
		powerMicroWatts := strings.TrimSpace(string(powerStr))
		powerMilliWatts, err := strconv.Atoi(powerMicroWatts)
		if err != nil {
			fmt.Println("Error converting power consumption:", err)
			return
		}

		// fmt.Printf("GPU Power Consumption: %.2f mW\n", float64(powerMilliWatts)/1000.0)

		// GPU Sys_Clock
		clockFilePath := "/sys/class/drm/card0/device/pp_dpm_sclk" // Adjust as needed
		clockInfoBytes, err := ioutil.ReadFile(clockFilePath)
		if err != nil {
			fmt.Println("Error reading clock info:", err)
			return
		}
		clockInfo := strings.TrimSpace(string(clockInfoBytes))

		lines := strings.Split(clockInfo, "\n")
		currentGPUClockState := ""
		currentGPUClockSpeed := 0 // New variable to store the clock speed as an integer
		for _, line := range lines {
			// fmt.Println("Lines", line)
			if strings.Contains(line, "*") {
				currentGPUClockState = line
				fmt.Println("Line", line)

				// Extract the numeric value using regex
				re := regexp.MustCompile(`\d+`)
				speedStr := re.FindString(line)
				fmt.Println("Regexed", speedStr)
				if speedStr != "" {
					// Convert the clean speed string to an integer
					currentGPUClockSpeed, err = strconv.Atoi(speedStr)
					fmt.Println("CGCS", currentGPUClockSpeed)
					if err != nil {
						fmt.Println("Error converting clock speed to integer:", err)
						return
					}
				}
				break
			}
		}

		// ////////////////////////

		// sclk label
		fmt.Println("\n")
		// sclk_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq1_label" // Adjust as needed
		// sclkBytes, err := ioutil.ReadFile(sclk_fp)
		// if err != nil {
		// 	fmt.Println("Error reading clock info:", err)
		// 	return
		// }

		// mclk label
		// mclk_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq2_label" // Adjust as needed
		// mclkBytes, err := ioutil.ReadFile(mclk_fp)
		// if err != nil {
		// 	fmt.Println("Error reading clock info:", err)
		// 	return
		// }

		// //////////////
		// sclk info
		sclkinfo_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq1_input" // Adjust as needed
		sclkinfoBytes, err := ioutil.ReadFile(sclkinfo_fp)
		if err != nil {
			fmt.Println("Error reading clock info:", err)
			return
		}

		// mclck info
		mclkinfo_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq2_input" // Adjust as needed
		mclkinfoBytes, err := ioutil.ReadFile(mclkinfo_fp)
		if err != nil {
			fmt.Println("Error reading clock info:", err)
			return
		}

		// Print the frequencies
		// sclk label
		// sclk_lbl := strings.TrimSpace(string(sclkBytes))

		// mclk label
		// mclk_lbl := strings.TrimSpace(string(mclkBytes))
		// sclk freq
		sclk_info := strings.TrimSpace(string(sclkinfoBytes))
		sclk := 0
		sclk, err = strconv.Atoi(sclk_info)
		if err != nil {
			// Handle the error
			fmt.Println("Error converting string to int:", err)
		} else {
			// fmt.Println("Converted integer:", sclk)
		}

		mclk_info := strings.TrimSpace(string(mclkinfoBytes))
		mclk := 0
		mclk, err = strconv.Atoi(mclk_info)
		if err != nil {
			// Handle the error
			fmt.Println("Error converting string to int:", err)
		} else {
			// fmt.Println("Converted integer:", mclk)
		}
		// fmt.Println(mclk_lbl, mclk)

		// mclk freq
		// fmt.Println(mclk_lbl, mclk_info)

		// /////////////////////////////////////////////////////

		// VRam Usage
		// Adjust these file paths as needed for your specific GPU and system
		vramTotalFilePath := "/sys/class/drm/card0/device/mem_info_vram_total"
		vramUsedFilePath := "/sys/class/drm/card0/device/mem_info_vram_used"

		vramTotalBytes, err := ioutil.ReadFile(vramTotalFilePath)
		if err != nil {
			fmt.Println("Error reading total VRAM:", err)
			return
		}

		vramUsedBytes, err := ioutil.ReadFile(vramUsedFilePath)
		if err != nil {
			fmt.Println("Error reading used VRAM:", err)
			return
		}

		// Convert the read values from bytes to integers
		vramTotal, err := strconv.Atoi(strings.TrimSpace(string(vramTotalBytes)))
		if err != nil {
			fmt.Println("Error converting total VRAM:", err)
			return
		}

		vramUsed, err := strconv.Atoi(strings.TrimSpace(string(vramUsedBytes)))
		if err != nil {
			fmt.Println("Error converting used VRAM:", err)
			return
		}

		// Calculate the percentage of VRAM used
		vramUsedPercentage := (float64(vramUsed) / float64(vramTotal)) * 100

		// fmt.Printf("GPU VRAM Usage: %.2f%%\n", vramUsedPercentage)

		// GPU Utilization %
		gpuUtilizationFilePath := "/sys/class/drm/card0/device/gpu_busy_percent" // Adjust as needed
		gpuUtilizationStr, err := ioutil.ReadFile(gpuUtilizationFilePath)
		if err != nil {
			fmt.Println("Error reading GPU utilization:", err)
			return
		}

		// Convert the utilization string to an integer
		gpuUtilization, err := strconv.Atoi(strings.TrimSpace(string(gpuUtilizationStr)))
		if err != nil {
			fmt.Println("Error converting GPU utilization:", err)
			return
		}

		// Clear the screen and move the cursor to the top left corner
		// fmt.Print("\033[H\033[2J")

		// // Print the table headers with proper spacing
		// fmt.Printf("%-25s %-25s %-25s %-20s %-15s\n",
		// 	"GPU Temperature (°C)",
		// 	"GPU Power Consumption (W)",
		// 	"Current GPU Clock State",
		// 	"GPU VRAM Usage (%)",
		// 	"GPU Utilization (%)")

		// Print the values in the same order as the headers
		// fmt.Printf("%-25.2f %-25.2f %-25s %-20.2f %-15d\n",
		// 	float64(tempDegrees)/1000.0,        // Convert millidegrees to degrees
		// 	float64(powerMilliWatts)/1000000.0, // Convert microwatts to milliwatts
		// 	currentGPUClockState,
		// 	vramUsedPercentage,
		// 	gpuUtilization)

		// Populate the GPUData struct with your data
		data := GPUData{
			Temperature:          float64(tempDegrees) / 1000.0,        // Convert millidegrees to degrees
			PowerConsumption:     float64(powerMilliWatts) / 1000000.0, // Convert microwatts to watts
			CurrentGPUClockState: currentGPUClockState,
			VramUsedPercentage:   vramUsedPercentage,
			GPUUtilization:       gpuUtilization,
			CurrentGPUClockSpeed: currentGPUClockSpeed,
			Sclk:                 sclk,
			Mclk:                 mclk,
		}

		fmt.Printf("Temp: %f\nPower: %f\nSCLK: %d\nMCLK: %d\nGPU_Clock: %d", data.Temperature, data.PowerConsumption, data.Sclk, data.Mclk, data.CurrentGPUClockSpeed)

		// Send the data over the WebSocket connection
		if err := conn.WriteJSON(data); err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
		time.Sleep(1 * time.Second)
	}

}

func main() {
	http.HandleFunc("/ws", serveWs)
	fmt.Println("Starting Server on port:8069 ")
	http.ListenAndServe(":8069", nil)
}
