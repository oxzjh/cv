#pragma once

#include "core.h"

#ifdef __cplusplus
#include <opencv2/dnn.hpp>
extern "C"
{
#endif

#ifdef __cplusplus
    typedef cv::dnn::Net *Net;
    typedef std::vector<cv::String> *Names;
#else
typedef void *Net;
typedef void *Names;
#endif

    Net net_alloc(const char *model, const char *config);
    Net net_alloc_with_buffer(const char *framework, const unsigned char *model, int modelSize, const unsigned char *config, int configSize);
    void net_free(Net net);
    bool net_empty(Net net);
    void net_set_input(Net net, Mat blob);
    void net_forward(Net net, Mat *outputBlobs, Names outputNames, int count);
    void net_set_preferable_backend(Net net, int backend);
    void net_set_preferable_target(Net net, int target);
    Names net_get_unconnected_out_layers_names(Net net);
    int net_names_size(Names names);
    const char *net_name_at(Names names, int index);
    void net_names_free(Names names);
    Mat net_blob_from_image(Mat img, float scale, Size size, Scalar mean, bool swapRB, bool crop);

#ifdef __cplusplus
}
#endif