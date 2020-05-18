# frozen_string_literal: true

require 'logger'
require 'sinatra'
require 'google/cloud/firestore'

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
  user_header = request.env['HTTP_X_GOOG_AUTHENTICATED_USER_EMAIL']
  @user = user_header.partition('accounts.google.com:').last unless user_header.nil?
end

get '/?' do
  logger = Logger.new($stdout)
  firestore_project = ENV['FIRESTORE_PROJECT']
  Google::Cloud::Firestore.configure { |config| config.project_id = firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects_col = firestore_client.col('gcp-projects-dashboard')
  projects_col.get do |project|
    logger.info "#{project.document_id} was last updated at #{project[:updated]}."
  end

  erb :index, locals: { title: 'GCP Projects Dashboard' }
end
