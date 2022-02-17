FROM node:fermium-alpine3.10

COPY package*.json ./

WORKDIR /

RUN npm install

COPY . .

EXPOSE 7070

CMD [ "npm", "start" ]