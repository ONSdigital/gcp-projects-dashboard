FROM ruby:4.0.0

WORKDIR /usr/src/app
COPY . ./

RUN gem install bundler --no-document && \ 
    bundle install