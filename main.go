package main

/*
#cgo LDFLAGS: -lOpenCL
#define CL_TARGET_OPENCL_VERSION 300
#include <CL/cl.h>
#cgo CFLAGS: -DCGO
#include "main.c"
*/
import "C"

import (
	"fmt"
	"regexp"
	"strings"
	"unsafe"
	_ "unsafe"
	// "github.com/gorilla/websocket"
	// _ "github.com/jgillich/go-opencl/cl"
)

type GPU struct {
	Vendor string   `json:vendor`
	Name   []string `json:name`
}

type GPU_Devices struct {
	Devices []GPU `json:devices`
}

func listIntelGPUs() {
	C.listIntelGPUs()
}
func main() {
	fmt.Println("Getting GPU Devices")
	resultC := C.listIntelGPUs()
	defer C.free(unsafe.Pointer(resultC))
	resultGo := C.GoString(resultC)
	if strings.Contains(resultGo, "Intel") {
		dev := FindAllStringsBetweenBrackets(resultGo)
		// fmt.Println("Intel GPU", dev)
		device := GPU{
			Vendor: "Intel",
			Name:   dev,
		}
		fmt.Println("Device", device)

	}

	// listIntelGPUs()
	// http.HandleFunc("/ws", serveWs)
	// fmt.Println("Starting Server on port:8069 ")
	// http.ListenAndServe(":8069", nil)

	// getDevices()
}

// FindAllStringsBetweenBrackets finds and returns all occurrences of strings between square brackets.
func FindAllStringsBetweenBrackets(str string) []string {
	re := regexp.MustCompile(`\[([^\[\]]+)\]`)
	matches := re.FindAllStringSubmatch(str, -1)
	var results []string
	for _, match := range matches {
		if len(match) > 1 {
			results = append(results, match[1]) // Append the first capturing group.
		}
	}
	return results
}

// const kernelSource = `
// __kernel void add(__global const float* a, __global const float* b, __global float* c) {
//     int gid = get_global_id(0);
//     c[gid] = a[gid] + b[gid];
// }
// `

// GPUData struct to hold all the GPU information
// type  AMD_GPUData struct {
// 	Temperature        float64 `json:"temperature"`
// 	PowerConsumption   float64 `json:"powerConsumption"`
// 	VramUsedPercentage float64 `json:"vramUsedPercentage"`
// 	GPUUtilization     int     `json:"gpuUtilization"`
// 	Sclk               int     `json:"sclk"`
// 	Mclk               int     `json:"mclk"`
// }

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true // Allow connections from any origin
// 	},
// }

// Uncomment to enable server again
// func serveWs(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("Error upgrading to WebSocket:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for {
//
// 		// GPU Temperature
// 		tempFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/temp1_input"
// 		tempStr, err := ioutil.ReadFile(tempFilePath)
// 		if err != nil {
// 			fmt.Println("Error reading temperature:", err)
// 			return
// 		}
//
// 		// The temperature is usually reported in millidegrees Celsius; convert to degrees
// 		tempMilliDegrees := strings.TrimSpace(string(tempStr))
// 		tempDegrees, err := strconv.Atoi(tempMilliDegrees)
// 		if err != nil {
// 			fmt.Println("Error converting temperature:", err)
// 			return
// 		}
// 		// fmt.Printf("GPU Temperature: %.2f°C\n", float64(tempDegrees)/1000.0)
//
// 		// Average Power Consumption
// 		powerFilePath := "/sys/class/drm/card0/device/hwmon/hwmon3/power1_average"
// 		powerStr, err := ioutil.ReadFile(powerFilePath)
// 		if err != nil {
// 			fmt.Println("Error reading power consumption:", err)
// 			return
// 		}
//
// 		// The power consumption is usually reported in microwatts; convert to milliwatts
// 		powerMicroWatts := strings.TrimSpace(string(powerStr))
// 		powerMilliWatts, err := strconv.Atoi(powerMicroWatts)
// 		if err != nil {
// 			fmt.Println("Error converting power consumption:", err)
// 			return
// 		}
//
// 		// //////////////
// 		// sclk info
// 		sclkinfo_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq1_input" // Adjust as needed
// 		sclkinfoBytes, err := ioutil.ReadFile(sclkinfo_fp)
// 		if err != nil {
// 			fmt.Println("Error reading clock info:", err)
// 			return
// 		}
//
// 		// mclck info
// 		mclkinfo_fp := "/sys/class/drm/card0/device/hwmon/hwmon3/freq2_input" // Adjust as needed
// 		mclkinfoBytes, err := ioutil.ReadFile(mclkinfo_fp)
// 		if err != nil {
// 			fmt.Println("Error reading clock info:", err)
// 			return
// 		}
//
// 		sclk_info := strings.TrimSpace(string(sclkinfoBytes))
// 		sclk := 0
// 		sclk, err = strconv.Atoi(sclk_info)
// 		if err != nil {
// 			// Handle the error
// 			fmt.Println("Error converting string to int:", err)
// 		} else {
// 			// fmt.Println("Converted integer:", sclk)
// 		}
//
// 		mclk_info := strings.TrimSpace(string(mclkinfoBytes))
// 		mclk := 0
// 		mclk, err = strconv.Atoi(mclk_info)
// 		if err != nil {
// 			// Handle the error
// 			fmt.Println("Error converting string to int:", err)
// 		} else {
// 			// fmt.Println("Converted integer:", mclk)
// 		}
//
// 		// /////////////////////////////////////////////////////
//
// 		// VRam Usage
// 		// Adjust these file paths as needed for your specific GPU and system
// 		vramTotalFilePath := "/sys/class/drm/card0/device/mem_info_vram_total"
// 		vramUsedFilePath := "/sys/class/drm/card0/device/mem_info_vram_used"
//
// 		vramTotalBytes, err := ioutil.ReadFile(vramTotalFilePath)
// 		if err != nil {
// 			fmt.Println("Error reading total VRAM:", err)
// 			return
// 		}
//
// 		vramUsedBytes, err := ioutil.ReadFile(vramUsedFilePath)
// 		if err != nil {
// 			fmt.Println("Error reading used VRAM:", err)
// 			return
// 		}
//
// 		// Convert the read values from bytes to integers
// 		vramTotal, err := strconv.Atoi(strings.TrimSpace(string(vramTotalBytes)))
// 		if err != nil {
// 			fmt.Println("Error converting total VRAM:", err)
// 			return
// 		}
//
// 		vramUsed, err := strconv.Atoi(strings.TrimSpace(string(vramUsedBytes)))
// 		if err != nil {
// 			fmt.Println("Error converting used VRAM:", err)
// 			return
// 		}
//
// 		// Calculate the percentage of VRAM used
// 		vramUsedPercentage := (float64(vramUsed) / float64(vramTotal)) * 100
//
// 		// fmt.Printf("GPU VRAM Usage: %.2f%%\n", vramUsedPercentage)
//
// 		// GPU Utilization %
// 		gpuUtilizationFilePath := "/sys/class/drm/card0/device/gpu_busy_percent" // Adjust as needed
// 		gpuUtilizationStr, err := ioutil.ReadFile(gpuUtilizationFilePath)
// 		if err != nil {
// 			fmt.Println("Error reading GPU utilization:", err)
// 			return
// 		}
//
// 		// Convert the utilization string to an integer
// 		gpuUtilization, err := strconv.Atoi(strings.TrimSpace(string(gpuUtilizationStr)))
// 		if err != nil {
// 			fmt.Println("Error converting GPU utilization:", err)
// 			return
// 		}
//
// 		// Populate the GPUData struct with your data
// 		data := GPUData{
// 			Temperature:        float64(tempDegrees) / 1000.0,        // Convert millidegrees to degrees
// 			PowerConsumption:   float64(powerMilliWatts) / 1000000.0, // Convert microwatts to watts
// 			VramUsedPercentage: vramUsedPercentage,
// 			GPUUtilization:     gpuUtilization,
// 			Sclk:               sclk,
// 			Mclk:               mclk,
// 		}
//
// 		fmt.Printf("Temp: %f\nPower: %f\nSCLK: %d\nMCLK: %d\nGPU_prcnt: %d\nVRam: %f", data.Temperature, data.PowerConsumption, data.Sclk, data.Mclk, gpuUtilization, data.VramUsedPercentage)
//
// 		// Send the data over the WebSocket connection
// 		if err := conn.WriteJSON(data); err != nil {
// 			fmt.Println("Error sending data:", err)
// 			break
// 		}
// 		fmt.Println("\n")
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func getDevices() {
// 	var numPlatforms C.cl_uint
// 	C.clGetPlatformIDs(0, nil, &numPlatforms)

// 	if numPlatforms == 0 {
// 		fmt.Println("No OpenCL platforms found.")
// 		return
// 	}

// 	platforms := make([]C.cl_platform_id, numPlatforms)
// 	C.clGetPlatformIDs(numPlatforms, &platforms[0], nil)

// 	// Assuming you want to use the first platform (Clover in your case)
// 	platform := platforms[0]

// 	var numDevices C.cl_uint
// 	C.clGetDeviceIDs(platform, C.CL_DEVICE_TYPE_GPU, 0, nil, &numDevices)

// 	if numDevices == 0 {
// 		fmt.Println("No OpenCL GPU devices found for this platform.")
// 		return
// 	}

// 	devices := make([]C.cl_device_id, numDevices)
// 	C.clGetDeviceIDs(platform, C.CL_DEVICE_TYPE_GPU, numDevices, &devices[0], nil)

// 	// Assuming you want to use the first GPU device
// 	device := devices[0]

// 	// Create a context
// 	var context C.cl_context
// 	context = C.clCreateContext(nil, 1, &device, nil, nil, nil)

// 	// Create a command queue
// 	var commandQueue C.cl_command_queue
// 	commandQueue = C.clCreateCommandQueue(context, device, 0, nil)

// 	// Compile the kernel
// 	kernelSourceStr := C.CString(kernelSource)
// 	defer C.free(unsafe.Pointer(kernelSourceStr))
// 	var program C.cl_program
// 	program = C.clCreateProgramWithSource(context, 1, &kernelSourceStr, nil, nil)
// 	if errCode := C.clBuildProgram(program, 1, &device, nil, nil, nil); errCode != C.CL_SUCCESS {
// 		// Handle error: You should retrieve the build log for detailed error messages
// 		fmt.Println("Failed to build program.")
// 		return
// 	}

// 	// Create the kernel
// 	kernelName := C.CString("add")
// 	defer C.free(unsafe.Pointer(kernelName))
// 	var kernel C.cl_kernel
// 	kernel = C.clCreateKernel(program, kernelName, nil)

// 	// Assume you have input arrays a and b, and an output array c
// 	// For simplicity, let's say they are all of size 10
// 	a := make([]float32, 10)
// 	b := make([]float32, 10)
// 	c := make([]float32, 10) // This will store the result

// 	// Create buffers for the kernel arguments
// 	var aMem, bMem, cMem C.cl_mem
// 	aMem = C.clCreateBuffer(context, C.CL_MEM_READ_ONLY|C.CL_MEM_COPY_HOST_PTR, C.size_t(len(a)*4), unsafe.Pointer(&a[0]), nil)
// 	bMem = C.clCreateBuffer(context, C.CL_MEM_READ_ONLY|C.CL_MEM_COPY_HOST_PTR, C.size_t(len(b)*4), unsafe.Pointer(&b[0]), nil)
// 	cMem = C.clCreateBuffer(context, C.CL_MEM_WRITE_ONLY, C.size_t(len(c)*4), nil, nil)

// 	// Set kernel arguments
// 	C.clSetKernelArg(kernel, 0, C.size_t(unsafe.Sizeof(aMem)), unsafe.Pointer(&aMem))
// 	C.clSetKernelArg(kernel, 1, C.size_t(unsafe.Sizeof(bMem)), unsafe.Pointer(&bMem))
// 	C.clSetKernelArg(kernel, 2, C.size_t(unsafe.Sizeof(cMem)), unsafe.Pointer(&cMem))

// 	// Enqueue the kernel for execution
// 	var globalWorkSize [1]C.size_t = [1]C.size_t{10} // Match the size of your arrays
// 	C.clEnqueueNDRangeKernel(commandQueue, kernel, 1, nil, &globalWorkSize[0], nil, 0, nil, nil)

// 	// Read the result back into c
// 	C.clEnqueueReadBuffer(commandQueue, cMem, C.CL_TRUE, 0, C.size_t(len(c)*4), unsafe.Pointer(&c[0]), 0, nil, nil)

// 	// Don't forget to release OpenCL resources when done
// 	C.clReleaseMemObject(aMem)
// 	C.clReleaseMemObject(bMem)
// 	C.clReleaseMemObject(cMem)
// 	C.clReleaseKernel(kernel)
// 	C.clReleaseProgram(program)
// 	C.clReleaseCommandQueue(commandQueue)
// 	C.clReleaseContext(context)
// }

// func main() {
// 	getDevices()
// }
