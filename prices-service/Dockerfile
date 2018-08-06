FROM ruby:2.5.1-alpine
EXPOSE 9292

RUN apk add --update \
  build-base \
  && rm -rf /var/cache/apk/*

ENV APP_HOME /app
WORKDIR $APP_HOME

ADD Gemfile* $APP_HOME/
RUN bundle install

ADD . $APP_HOME

ENTRYPOINT ["bundle", "exec"]
CMD ["puma", "config.ru", "-t", "8:32", "-w", "4"]
