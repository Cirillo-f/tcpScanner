FROM nginx:alpine

RUN rm -rf /usr/share/nginx/html/*

COPY index.html /usr/share/nginx/html/
COPY style.css /usr/share/nginx/html/
COPY scripts.js /usr/share/nginx/html/

COPY nginx.conf /etc/nginx/nginx.conf

RUN apk del --purge \
    && rm -rf /var/cache/apk/* \
    && rm -rf /tmp/*

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"] 