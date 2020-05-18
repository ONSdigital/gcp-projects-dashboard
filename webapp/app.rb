# frozen_string_literal: true

require 'sinatra'

before do
  headers 'Content-Type' => 'text/html; charset=utf-8'
  user_header = request.env['HTTP_X_GOOG_AUTHENTICATED_USER_EMAIL']
  @user = user_header.partition('accounts.google.com:').last unless user_header.nil?
end

get '/?' do
  erb :index, locals: { title: 'GCP Projects Dashboard' }
end
