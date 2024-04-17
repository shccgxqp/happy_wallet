FROM nginx

WORKDIR /usr/share/react

RUN curl -sSfL https://deb.nodesource.com/setup_18.x | bash -

RUN apt-get install -y nodejs

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

RUN rm -r /usr/share/nginx/html/*

RUN cp -a build/. /usr/share/nginx/html