#include <bits/stdc++.h>
#include <cstdint>
#include <cstdlib>
#include <cstring>
#include <opencv2/opencv.hpp>
#include <string>
#include <vector>
using namespace std;

void dispImg(vector<cv::Mat> images);
cv::Mat handleimg(int argc, char *argv[]);

int main(int argc, char *argv[]) {
  auto img = handleimg(argc, argv);
  auto gimg = cv::Mat(), bimg = cv::Mat();
  cv::cvtColor(img, gimg, cv::COLOR_BGR2GRAY);
  cv::threshold(gimg, bimg, 0, 255.0, cv::THRESH_BINARY | cv::THRESH_OTSU);
  // write
  cv::imwrite("./outputs/base.png", img);
  cv::imwrite("./outputs/gray_scale.png", gimg);
  cv::imwrite("./outputs/binarized.png", bimg);
  // dispImg({img, gimg, bimg});
  exit(EXIT_SUCCESS);
}

cv::Mat handleimg(int argc, char *argv[]) {
  if (argc == 1) {
    cout << "No image\n";
    exit(EXIT_FAILURE);
  }
  auto img = cv::imread(argv[1]);
  if (img.empty()) {
    cout << "image not found\n";
    exit(EXIT_FAILURE);
  }
  return img;
}

void dispImg(std::vector<cv::Mat> images) {
  size_t idx = 0;
  for (const auto &img : images) {
    cv::imshow(string("station").append(to_string(idx)), img);
    idx++;
  }
  while (1) {
    auto key = cv::waitKey(0);
    if (key == 27 || key == 113) {
      break;
    }
  }
}
