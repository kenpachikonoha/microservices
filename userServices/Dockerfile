FROM node:22.5-alpine3.19

WORKDIR /usr/src/app

COPY . .


RUN npm install
RUN npm install amqplib


EXPOSE 3030

CMD ["node", "run", "start"]

