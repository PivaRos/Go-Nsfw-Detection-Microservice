# nudenet.py

import base64
from io import BytesIO
from PIL import Image
from nudenet import NudeClassifier
import json
import sys

classifier = NudeClassifier()

def classify_image_base64(image_base64):
    image_data = base64.b64decode(image_base64)
    image = Image.open(BytesIO(image_data))
    image.save("/tmp/temp_image.jpg")  # Save the image temporarily
    result = classifier.classify("/tmp/temp_image.jpg")
    return json.dumps(result)

if __name__ == "__main__":
    image_base64 = sys.argv[1]
    result = classify_image_base64(image_base64)
    print(result)
