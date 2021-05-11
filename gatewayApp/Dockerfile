FROM python:3.8-alpine

WORKDIR /app

COPY config/requirements.txt ./config/

RUN pip install -r config/requirements.txt

COPY ./ ./

CMD [ "python", "main.py" ]