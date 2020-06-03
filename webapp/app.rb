# frozen_string_literal: true

require 'sinatra'
require 'google/cloud/firestore'

FIRESTORE_DATA_COLLECTION  = 'gcp-projects-dashboard'
FIRESTORE_PREFS_COLLECTION = 'gcp-projects-dashboard-preferences'

helpers do
  def d(text)
    Time.parse(text).utc.strftime('%d/%m/%Y %H:%M')
  end

  def h(text)
    Rack::Utils.escape_html(text)
  end
end

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
  @firestore_project = ENV['FIRESTORE_PROJECT']
  raise 'Missing FIRESTORE_PROJECT environment variable' unless @firestore_project

  user_header = request.env['HTTP_X_GOOG_AUTHENTICATED_USER_EMAIL']
  @user = user_header.partition('accounts.google.com:').last unless user_header.nil?
end

get '/?' do
  gcp_console_base_url = ENV['GCP_CONSOLE_BASE_URL']
  raise 'Missing GCP_CONSOLE_BASE_URL environment variable' unless gcp_console_base_url

  gcp_organisation = ENV['GCP_ORGANISATION']
  raise 'Missing GCP_ORGANISATION environment variable' unless gcp_organisation

  Google::Cloud::Firestore.configure { |config| config.project_id = @firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects = firestore_client.col(FIRESTORE_DATA_COLLECTION).list_documents.all

  user_prefs = firestore_client.col(FIRESTORE_PREFS_COLLECTION).doc(@user)
  bookmarks = []
  bookmarks = user_prefs.get[:bookmarks] unless user_prefs.get.data.nil?

  erb :index, locals: { title: "#{gcp_organisation} - GCP Projects Dashboard",
                        gcp_console_base_url: gcp_console_base_url,
                        projects: projects,
                        bookmarks: bookmarks }
end

get '/bookmarks?' do
  gcp_console_base_url = ENV['GCP_CONSOLE_BASE_URL']
  raise 'Missing GCP_CONSOLE_BASE_URL environment variable' unless gcp_console_base_url

  gcp_organisation = ENV['GCP_ORGANISATION']
  raise 'Missing GCP_ORGANISATION environment variable' unless gcp_organisation

  Google::Cloud::Firestore.configure { |config| config.project_id = @firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  projects = firestore_client.col(FIRESTORE_DATA_COLLECTION).list_documents.all

  user_prefs = firestore_client.col(FIRESTORE_PREFS_COLLECTION).doc(@user)
  bookmarks = []
  bookmarks = user_prefs.get[:bookmarks] unless user_prefs.get.data.nil?

  bookmarked_projects = []
  projects.each { |project| bookmarked_projects << project if bookmarks.include?(project.document_id) }

  erb :bookmarks, locals: { title: "#{gcp_organisation} Bookmarks - GCP Projects Dashboard",
                            gcp_console_base_url: gcp_console_base_url,
                            projects: bookmarked_projects }
end

get '/health?' do
  halt 200
end

post '/addbookmark?' do
  Google::Cloud::Firestore.configure { |config| config.project_id = @firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  user_prefs = firestore_client.col(FIRESTORE_PREFS_COLLECTION).doc(@user)
  bookmarks = []
  bookmarks = user_prefs.get[:bookmarks] unless user_prefs.get.data.nil?
  bookmark = params[:bookmark]
  bookmarks << bookmark unless bookmarks.include?(bookmark)
  user_prefs.set({ bookmarks: bookmarks })
end

post '/removebookmark?' do
  Google::Cloud::Firestore.configure { |config| config.project_id = @firestore_project }
  firestore_client = Google::Cloud::Firestore.new
  user_prefs = firestore_client.col(FIRESTORE_PREFS_COLLECTION).doc(@user)
  bookmarks = []
  bookmarks = user_prefs.get[:bookmarks] unless user_prefs.get.data.nil?
  bookmark = params[:bookmark]
  bookmarks&.delete(bookmark)
  user_prefs.set({ bookmarks: bookmarks })
end
