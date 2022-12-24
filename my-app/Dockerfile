FROM node:18-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH

# COPY .npmrc ./
COPY ./my-app/package*.json /.
RUN npm install

COPY ./my-app ./

CMD ["npm", "start"]