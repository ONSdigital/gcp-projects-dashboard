FROM ruby:3.1.1

WORKDIR /usr/src/app
COPY . ./

RUN gem install bundler --no-document && \ 
    addgroup --gid 1000 sinatra && \
    adduser --system --uid 1000 --gid 1000 sinatra && \
    chown -R sinatra:sinatra ./ && \
    bundle install