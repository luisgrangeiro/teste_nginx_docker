FROM nginx:latest
RUN rm /etc/nginx/conf.d/*
COPY index.html /usr/share/nginx/html/index.html
COPY default.conf /etc/nginx/conf.d/