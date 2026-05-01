#pragma once

#include "core.h"

#ifdef __cplusplus
extern "C"
{
#endif

    typedef struct Contours
    {
        int num;
        int *counts;
        Point **points;
    } Contours;

    typedef struct AreaRect
    {
        Point2f points[4];
        Size2f size;
        float angle;
    } AreaRect;

    typedef struct AreaRects
    {
        int num;
        AreaRect *rects;
    } AreaRects;

    void cvt_color(Mat src, Mat dst, int code);
    void resize(Mat src, Mat dst, Size size, float fx, float fy, int interp);
    void rotate(Mat src, Mat dst, int code);
    void copy_make_border(Mat src, Mat dst, int top, int bottom, int left, int right, int borderType, Scalar color);
    void warp_affine(Mat src, Mat dst, Mat m, Size size);
    void warp_perspective(Mat src, Mat dst, Mat m, Size size);
    void calc_hist(Mat img, Mat hist);
    void blur(Mat src, Mat dst, Size size);
    void gaussian_blur(Mat src, Mat dst, Size size, float sigmaX, float sigmaY, int borderType);
    void canny(Mat src, Mat edges, float threshold1, float threshold2);
    void circle(Mat img, Point center, int radius, Scalar color, int thickness);
    void ellipse(Mat img, Point center, Size axes, float angle, float startAngle, float endAngle, Scalar color, int thickness);
    void line(Mat img, Point p1, Point p2, Scalar color, int thickness);
    void rectangle(Mat img, Rect rect, Scalar color, int thickness);
    void fill_poly(Mat img, Point *points, int count, Scalar color);
    void polylines(Mat img, Point *points, int count, bool isClosed, Scalar color, int thickness);
    void put_text(Mat img, const char *text, Point org, int fontFace, float fontScale, Scalar color, int thickness);
    void watershed(Mat img, Mat markers);
    AreaRect min_area_rect(Point *points, int count);
    float box_score(Mat img, Point2f points[4]);
    Contours *find_contours(Mat img, int mode, int method);
    void free_contours(Contours *contours);
    AreaRects *find_area_rects(Mat img, int mode, int method);
    void free_area_rects(AreaRects *rects);
    float compare_hist(Mat hist1, Mat hist2);
    Mat get_affine_transform(Point2f *src, Point2f *dst, int count);
    Mat get_perspective_transform(Point2f *src, Point2f *dst, int count);
    Mat get_rotation_matrix_2d(Point2f center, float angle, float scale);
    Mat get_similarity_transform_matrix(float *landmarks);

#ifdef __cplusplus
}
#endif