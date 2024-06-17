import base64
from io import BytesIO
from PIL import Image
from nudenet import NudeDetector
import json
import sys

detector = NudeDetector()

def classify_image_base64(image_base64):
    try:
        image_data = base64.b64decode(image_base64)
        image = Image.open(BytesIO(image_data))
        image.save("/tmp/temp_image.jpg")  # Save the image temporarily
        result = detector.detect("/tmp/temp_image.jpg")
        return json.dumps(result)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python3 classify_nsfw.py <base64_image>", file=sys.stderr)
        sys.exit(1)
    
    image_base64 = sys.argv[1]
    result = classify_image_base64(image_base64)
    print(result)
