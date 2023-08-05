import tempfile

import cv2
from fastapi import HTTPException

path = "/Users/macos/Downloads/file_example_MOV_1280_1_4MB.mov"

# cap = cv2.VideoCapture(path)
# count = 0
# while cap.isOpened():
#     ret,frame = cap.read()
#     cv2.imshow('window-name', frame)
#     cv2.imwrite("frame%d.jpg" % count, frame)
#     count = count + 1
#     if cv2.waitKey(10) & 0xFF == ord('q'):
#         break

# cap.release()
# cv2.destroyAllWindows() 





if __name__ == '__main__':
    t = 0
    delta = 50

    with tempfile.NamedTemporaryFile() as file_tmp, open(path, 'rb') as fp:
        # Write data to temporary file
        file_tmp.write(fp.read())
        file_tmp.seek(0)

        print(file_tmp.name)

        # Load video using opencv
        try:
            cap = cv2.VideoCapture(file_tmp.name)

            n = 0
            while cap.isOpened():
                if n % delta == 0:
                    _, frame = cap.read()
                    if frame is None:
                        break

                    print(type(frame))

                    cv2.imwrite(f"frame_{t:03d}.jpg", frame)

                n = (n + 1) % delta
        except Exception:
            if 'cap' in locals():
                cap.release()

            raise HTTPException(500, detail="Internal error")
        
    