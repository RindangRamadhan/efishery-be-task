FROM node:fermium-alpine3.10

WORKDIR /app

COPY package*.json ./

RUN yes | npm install

COPY . .

EXPOSE 7070

CMD [ "npm", "start" ]