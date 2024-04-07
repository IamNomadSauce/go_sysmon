#define CL_TARGET_OPENCL_VERSION 300
#include <stdlib.h>
#include <stdio.h>
#include <CL/cl.h>

char* listIntelGPUs() {
    cl_uint platformCount;
    cl_platform_id *platforms;
    cl_uint deviceCount;
    cl_device_id *devices;
    char deviceName[128];
    char* result = NULL;
    size_t resultSize = 0;
    size_t resultCapacity = 0;

    // Get the number of platforms
    clGetPlatformIDs(0, NULL, &platformCount);
    platforms = (cl_platform_id*)malloc(sizeof(cl_platform_id) * platformCount);

    // Get the platform IDs
    clGetPlatformIDs(platformCount, platforms, NULL);

    for (cl_uint i = 0; i < platformCount; i++) {
        // Get the number of GPU devices available on the platform
        clGetDeviceIDs(platforms[i], CL_DEVICE_TYPE_GPU, 0, NULL, &deviceCount);
        devices = (cl_device_id*)malloc(sizeof(cl_device_id) * deviceCount);

        // Get the GPU device IDs
        clGetDeviceIDs(platforms[i], CL_DEVICE_TYPE_GPU, deviceCount, devices, NULL);

        for (cl_uint j = 0; j < deviceCount; j++) {
            // Get the name of the GPU device
            clGetDeviceInfo(devices[j], CL_DEVICE_NAME, 128, deviceName, NULL);
            
            // Allocate or resize result buffer
            size_t neededSize = resultSize + snprintf(NULL, 0, "%d %s\n", j, deviceName);
            if (neededSize >= resultCapacity) {
                resultCapacity = neededSize + 1; // +1 for null-terminator
                result = realloc(result, resultCapacity);
            }
            
            // Append device info to result
            resultSize += sprintf(result + resultSize, "%d %s\n", j, deviceName);
        }

        free(devices);
    }

    free(platforms);
    return result; // Caller is responsible for freeing this memory
}

#ifndef CGO

int main() {
    listIntelGPUs();
    return 0;
}
#endif

