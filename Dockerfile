FROM python:3.10-slim

WORKDIR /code

RUN apt-get update && apt-get install ffmpeg libsm6 libxext6  -y

RUN pip install -r requirements.txt

# ENTRYPOINT ["tail", "-f", "/dev/null"]
ENTRYPOINT [ "bash", "run.sh" ]