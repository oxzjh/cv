#pragma once

#include "core.h"

#ifdef __cplusplus
extern "C"
{
#endif

    Mat imread(const char *filename, int flags);
    Mat imdecode(const unsigned char *data, int size, int flags);
    bool imwrite(const char *filename, Mat img);

#ifdef __cplusplus
}
#endif