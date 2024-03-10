package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

		// The file might contain multiple lines with clock states, e.g.:
		// 0: 300Mhz
		// 1: 600Mhz *
		// The line with the '*' indicates the current clock state.
		lines := strings.Split(clockInfo, "\n")
		currentGPUClockState := ""
		for _, line := range lines {
			if strings.Contains(line, "*") {
				currentGPUClockState = line
				// fmt.Println("Current GPU Clock State:", line)
				break
			}
		}

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
		fmt.Print("\033[H\033[2J")

		// Print the table headers with proper spacing
		fmt.Printf("%-25s %-25s %-25s %-20s %-15s\n",
			"GPU Temperature (°C)",
			"GPU Power Consumption (W)",
			"Current GPU Clock State",
			"GPU VRAM Usage (%)",
			"GPU Utilization (%)")

		// Print the values in the same order as the headers
		fmt.Printf("%-25.2f %-25.2f %-25s %-20.2f %-15d\n",
			float64(tempDegrees)/1000.0,        // Convert millidegrees to degrees
			float64(powerMilliWatts)/1000000.0, // Convert microwatts to milliwatts
			currentGPUClockState,
			vramUsedPercentage,
			gpuUtilization)

		// Populate the GPUData struct with your data
		data := GPUData{
			Temperature:          float64(tempDegrees) / 1000.0,        // Convert millidegrees to degrees
			PowerConsumption:     float64(powerMilliWatts) / 1000000.0, // Convert microwatts to watts
			CurrentGPUClockState: currentGPUClockState,
			VramUsedPercentage:   vramUsedPercentage,
			GPUUtilization:       gpuUtilization,
		}

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
