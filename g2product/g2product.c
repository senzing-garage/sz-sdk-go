#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include "libg2product.h"
#include "g2product.h"

void* G2Product_resizeStringBuffer(void* ptr, size_t size) {
    //deallocate old buffer
    if (ptr != 0)
        free(ptr);
    //allocate new buffer
    void* buffer = malloc(size);
    return buffer;
}

struct G2Product_validateLicenseFile_result G2Product_validateLicenseFile_helper(const char* licenseFilePath) {
    size_t charBufferSize = 1;
    char* charBuffer = (char*)malloc(charBufferSize);
    resize_buffer_type resizeFuncPointer = &G2Product_resizeStringBuffer;
    int returnCode = G2Product_validateLicenseStringBase64(licenseFilePath, &charBuffer, &charBufferSize, resizeFuncPointer);
    struct G2Product_validateLicenseFile_result result;
    result.response = charBuffer;
    result.returnCode = returnCode;
    return result;
}

struct G2Product_validateLicenseStringBase64_result G2Product_validateLicenseStringBase64_helper(const char* licenseString) {
    size_t charBufferSize = 1;
    char* charBuffer = (char*)malloc(charBufferSize);
    resize_buffer_type resizeFuncPointer = &G2Product_resizeStringBuffer;
    int returnCode = G2Product_validateLicenseStringBase64(licenseString, &charBuffer, &charBufferSize, resizeFuncPointer);
    struct G2Product_validateLicenseStringBase64_result result;
    result.response = charBuffer;
    result.returnCode = returnCode;
    return result;
}