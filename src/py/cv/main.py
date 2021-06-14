import numpy as np
import cv2

# 彩色图像 BGR格式
blue = np.zeros((300, 300, 3), dtype=np.uint8)
blue[:,:,0] = 255
cv2.imshow(blue)