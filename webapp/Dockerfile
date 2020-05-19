FROM ruby:2.7.1
# RUN apt-get update && apt-get install -y locales

# ENV LANG en_GB.UTF-8
# ENV LANGUAGE en_GB.UTF-8
# ENV LC_ALL en_GB.UTF-8

WORKDIR /usr/src/app
COPY . ./

RUN gem install bundler --no-document && \
    addgroup --gid 1000 sinatra && \
    adduser --system --uid 1000 --gid 1000 sinatra && \
    chown -R sinatra:sinatra ./ && \
    bundle install

USER sinatra

CMD ["bundle", "exec", "puma", "config.ru", "-C", "puma.rb", "-e", "production"]