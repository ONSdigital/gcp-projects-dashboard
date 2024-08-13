FROM europe-west2-docker.pkg.dev/ons-ci/ons-ci/gcp-projects-dashboard-webapp-parent-image:latest

WORKDIR /usr/src/app
COPY . ./

RUN groupadd --system --gid 1000 webapp && \
    useradd webapp --uid 1000 --gid 1000 --create-home --shell /bin/bash && \
    chown -R webapp:webapp .

USER 1000:1000

CMD ["bundle", "exec", "puma", "config.ru", "-C", "puma.rb", "-e", "production"]