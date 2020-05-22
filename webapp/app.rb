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
  raise 'Missing FIRESTORE_PROJECT environment variable' unless firestore_project

  gcp_console_base_url = ENV['GCP_CONSOLE_BASE_URL']
  raise 'Missing GCP_CONSOLE_BASE_URL environment variable' unless gcp_console_base_url

  gcp_organisation = ENV['GCP_ORGANISATION']
  raise 'Missing GCP_ORGANISATION environment variable' unless gcp_organisation

  Google::Cloud::Firestore.configure { |config| config.project_id = firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects = firestore_client.col('gcp-projects-dashboard').list_documents.all

  erb :index, locals: { title: "#{gcp_organisation} - GCP Projects Dashboard",
                        gcp_console_base_url: gcp_console_base_url,
                        projects: projects }
end

get '/health?' do
  halt 200
end
