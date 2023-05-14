FROM nginx:latest

# remove exists ngix conf and add my own nginx conf
RUN rm /etc/nginx/conf.d/default.conf
COPY default.conf /etc/nginx/conf.d/default.conf