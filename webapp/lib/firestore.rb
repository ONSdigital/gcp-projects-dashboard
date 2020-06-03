# frozen_string_literal: true

# Class to manage access to Firestore.
class Firestore
  FIRESTORE_DATA_COLLECTION  = 'gcp-projects-dashboard'
  FIRESTORE_PREFS_COLLECTION = 'gcp-projects-dashboard-preferences'

  def initialize(project)
    Google::Cloud::Firestore.configure { |config| config.project_id = project }
    @client = Google::Cloud::Firestore.new
  end

  def add_bookmark(user, bookmark)
    preferences = @client.col(FIRESTORE_PREFS_COLLECTION).doc(user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks << bookmark unless bookmarks.include?(bookmark)
    preferences.set({ bookmarks: bookmarks })
  end

  def all_projects
    @client.col(FIRESTORE_DATA_COLLECTION).list_documents.all
  end

  def bookmarked_projects(user)
    bookmarks = bookmarks(user)
    all_projects.collect { |project| project if bookmarks.include?(project.document_id) }
  end

  def bookmarks(user)
    preferences = @client.col(FIRESTORE_PREFS_COLLECTION).doc(user)
    return [] if preferences.get.data.nil?

    preferences.get[:bookmarks]
  end

  def remove_bookmark(user, bookmark)
    preferences = @client.col(FIRESTORE_PREFS_COLLECTION).doc(user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks&.delete(bookmark)
    preferences.set({ bookmarks: bookmarks })
  end
end
