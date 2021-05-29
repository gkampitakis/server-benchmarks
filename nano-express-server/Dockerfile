FROM node:lts

WORKDIR /app

COPY package.json . 

RUN npm i 

COPY . .

EXPOSE 5000

USER node

CMD ["node","index.js"]