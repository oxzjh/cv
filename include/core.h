#pragma once

#include <stdbool.h>

#ifdef __cplusplus
#include <opencv2/core.hpp>
extern "C"
{
#endif

    typedef struct Point
    {
        int x;
        int y;
    } Point;

    typedef struct Point2f
    {
        float x;
        float y;
    } Point2f;

    typedef struct Size
    {
        int width;
        int height;
    } Size;

    typedef struct Size2f
    {
        float width;
        float height;
    } Size2f;

    typedef struct Rect
    {
        int x;
        int y;
        int width;
        int height;
    } Rect;

    typedef struct Rect2f
    {
        float x;
        float y;
        float width;
        float height;
    } Rect2f;

    typedef struct Scalar
    {
        float val1;
        float val2;
        float val3;
        float val4;
    } Scalar;

#ifdef __cplusplus
    typedef cv::Mat *Mat;
#else
typedef void *Mat;
#endif

    int get_num_threads();
    void set_num_threads(int n);

    Mat mat_alloc();
    Mat mat_alloc_with_data(int rows, int cols, int type, void *data);
    void mat_free(Mat m);
    bool mat_empty(Mat m);
    Mat mat_clone(Mat m);
    void mat_convert_to(Mat m, Mat dst, int type);
    void mat_convert_to_ab(Mat m, Mat dst, int type, float alpha, float bata);
    Mat mat_region(Mat m, Rect r);
    Mat mat_reshape(Mat m, int cn, int rows);
    int mat_rows(Mat m);
    int mat_cols(Mat m);
    int mat_channels(Mat m);
    int mat_type(Mat m);
    int mat_total(Mat m);
    int mat_elem_size(Mat m);
    void *mat_data(Mat m);
    Mat mat_mul(Mat m, Mat n, float scale);
    void mat_divide(Mat m, Scalar s);
    void mat_add_uchar(Mat m, unsigned char v);
    void mat_subtract_uchar(Mat m, unsigned char v);
    void mat_multiply_uchar(Mat m, unsigned char v);
    void mat_divide_uchar(Mat m, unsigned char v);
    void mat_add_float(Mat m, float v);
    void mat_subtract_float(Mat m, float v);
    void mat_multiply_float(Mat m, float v);
    void mat_divide_float(Mat m, float v);
    void mat_add_mat(Mat m, Mat n);
    void mat_subtract_mat(Mat m, Mat n);
    Mat mat_multiply_matrix(Mat m, Mat n);
    Mat mat_t(Mat m);
    void mat_split(Mat m, Mat *ms);
    void mat_merge(Mat *ms, int count, Mat dst);
    Scalar mat_mean(Mat m, Mat mask);

#ifdef __cplusplus
}
#endif