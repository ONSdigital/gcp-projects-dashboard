#!/bin/sh
set -e

#
# Script for updating this project's RubyGems to the latest versions by running `bundle install`
# and committing and pushing the regenerated Gemfile.lock files to GitHub.
#
# Usage: update-rubygems.sh
#
# Author: John Topley (john.topley@ons.gov.uk)
#
echo "Regenerating webapp-parent-image/Gemfile.lock..."
cd ./webapp-parent-image
rm Gemfile.lock
bundle install

git add Gemfile.lock
git commit -m "Updated dependencies"
git push

echo "Waiting 5 minutes for Cloud Build pipeline to build webapp-parent-image..."
sleep 5m

cd ./webapp
rm Gemfile.lock
bundle install

echo "Regenerating webapp/Gemfile.lock..."
git add Gemfile.lock
git commit -m "Updated dependencies"
git push

echo "Finished!"
