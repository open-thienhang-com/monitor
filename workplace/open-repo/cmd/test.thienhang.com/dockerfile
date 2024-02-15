FROM python:3.10-alpine

ENV DOCKER_APP True

COPY requirements.txt .

RUN pip install package.whl

COPY . app/

WORKDIR app/

ENTRYPOINT ["python3", "main.py"]