# frozen_string_literal: true

require 'sinatra'
require 'sinatra/partial'
require 'google/cloud/firestore'

require_relative 'lib/configuration'
require_relative 'lib/firestore'

set :partial_template_engine, :erb

config = Configuration.new(ENV)
set :firestore_project,                 config.firestore_project
set :gcp_console_base_url,              config.gcp_console_base_url
set :gcp_console_cloud_armour_base_url, config.gcp_console_cloud_armour_base_url
set :gcp_organisation,                  config.gcp_organisation

helpers do
  def d(text)
    Time.parse(text).utc.strftime('%d/%m/%Y %H:%M')
  end

  def h(text)
    Rack::Utils.escape_html(text)
  end

  def n(float)
    float.to_i.to_s.reverse.scan(/\d{3}|.+/).join(',').reverse
  end
end

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
  user_header = request.env['HTTP_X_GOOG_AUTHENTICATED_USER_EMAIL']
  @user = user_header.partition('accounts.google.com:').last unless user_header.nil?
end

get '/?' do
  firestore = Firestore.new(settings.firestore_project)
  erb :index, locals: { title: "#{settings.gcp_organisation} - GCP Projects Dashboard",
                        gcp_console_base_url: settings.gcp_console_base_url,
                        master_version_alerts: firestore.all_master_version_alerts,
                        node_version_alerts: firestore.all_node_version_alerts,
                        projects: firestore.all_projects,
                        bookmarks: firestore.bookmarks(@user) }
end

get '/bookmarks?' do
  firestore = Firestore.new(settings.firestore_project)
  erb :bookmarks, locals: { title: "#{settings.gcp_organisation} Bookmarks - GCP Projects Dashboard",
                            gcp_console_base_url: settings.gcp_console_base_url,
                            master_version_alerts: firestore.all_master_version_alerts,
                            node_version_alerts: firestore.all_node_version_alerts,
                            projects: firestore.bookmarked_projects(@user) }
end

get '/bookmarks-cloudarmour?' do
  firestore = Firestore.new(settings.firestore_project)
  erb :bookmarkscloudarmour, locals: { title: "#{settings.gcp_organisation} Bookmarks - GCP Projects Dashboard",
                                       gcp_console_cloud_armour_base_url: settings.gcp_console_cloud_armour_base_url,
                                       security_rules: firestore.bookmarked_security_rules(@user) }
end

get '/cloudarmour?' do
  firestore = Firestore.new(settings.firestore_project)
  erb :cloudarmour, locals: { title: "#{settings.gcp_organisation} - GCP Projects Dashboard",
                              gcp_console_cloud_armour_base_url: settings.gcp_console_cloud_armour_base_url,
                              security_rules: firestore.all_security_rules,
                              bookmarks: firestore.bookmarks(@user) }
end

get '/health?' do
  halt 200
end

# The routes below are invoked from AJAX actions.
post '/addbookmark?' do
  firestore = Firestore.new(settings.firestore_project)
  firestore.add_bookmark(@user, params[:bookmark])
end

post '/removebookmark?' do
  firestore = Firestore.new(settings.firestore_project)
  firestore.remove_bookmark(@user, params[:bookmark])
end
