FROM python:3.12-alpine3.17

WORKDIR /app

COPY . . 

RUN pip install -r requirements.txt
RUN python3 -m pip install pika --upgrade

EXPOSE 5095

CMD ["waitress-serve", "--port=5095", "--call", "app:create_app"]