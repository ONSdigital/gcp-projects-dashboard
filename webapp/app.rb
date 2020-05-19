# frozen_string_literal: true

require 'logger'
require 'sinatra'
require 'google/cloud/firestore'

helpers do
  def d(text)
    Time.parse(text).utc.strftime('%A %d %b %Y %H:%M UTC')
  end

  def h(text)
    Rack::Utils.escape_html(text)
  end
end

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
  user_header = request.env['HTTP_X_GOOG_AUTHENTICATED_USER_EMAIL']
  @user = user_header.partition('accounts.google.com:').last unless user_header.nil?
end

get '/?' do
  firestore_project = ENV['FIRESTORE_PROJECT']
  Google::Cloud::Firestore.configure { |config| config.project_id = firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects = firestore_client.col('gcp-projects-dashboard').list_documents

  erb :index, locals: { title: 'GCP Projects Dashboard', projects: projects }
end
