# frozen_string_literal: true

require 'sinatra'
require 'google/cloud/firestore'

helpers do
  def d(text)
    Time.parse(text).utc.strftime('%A %d %b %Y %H:%M:%S UTC')
  end

  def h(text)
    Rack::Utils.escape_html(text)
  end
end

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
end

get '/?' do
  firestore_project = ENV['FIRESTORE_PROJECT']
  Google::Cloud::Firestore.configure { |config| config.project_id = firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects = firestore_client.col('gcp-projects-dashboard').list_documents
  puts projects.size

  erb :index, locals: { title: 'GCP Projects Dashboard', projects: projects }
end
