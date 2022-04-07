FROM europe-west2-docker.pkg.dev/ons-ci/ons-ci/gcp-projects-dashboard-webapp-parent-image:latest

WORKDIR /usr/src/app
COPY . ./

USER sinatra

CMD ["bundle", "exec", "puma", "config.ru", "-C", "puma.rb", "-e", "production"]