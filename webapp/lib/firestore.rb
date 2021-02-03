# frozen_string_literal: true

# Class to manage access to Firestore.
class Firestore
  FIRESTORE_CLUSTERS_COLLECTION      = 'gcp-projects-dashboard'
  FIRESTORE_MASTER_ALERTS_COLLECTION = 'gcp-projects-dashboard-gke-master-version-alerts'
  FIRESTORE_NODE_ALERTS_COLLECTION   = 'gcp-projects-dashboard-gke-node-version-alerts'
  FIRESTORE_PREFERENCES_COLLECTION   = 'gcp-projects-dashboard-preferences'

  def initialize(project)
    Google::Cloud::Firestore.configure { |config| config.project_id = project }
    @client = Google::Cloud::Firestore.new
  end

  def add_bookmark(user, bookmark)
    preferences = @client.col(FIRESTORE_PREFERENCES_COLLECTION).doc(user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks << bookmark unless bookmarks.include?(bookmark)
    preferences.set({ bookmarks: bookmarks })
  end

  def all_master_version_alerts
    @client.col(FIRESTORE_MASTER_ALERTS_COLLECTION).list_documents.all
  end

  def all_node_version_alerts
    @client.col(FIRESTORE_NODE_ALERTS_COLLECTION).list_documents.all
  end
  
  def all_projects
    @client.col(FIRESTORE_CLUSTERS_COLLECTION).list_documents.all
  end

  def bookmarked_projects(user)
    bookmarks = bookmarks(user)
    bookmarked_projects = []
    all_projects.each { |project| bookmarked_projects << project if bookmarks.include?(project.document_id) }
    bookmarked_projects
  end

  def bookmarks(user)
    preferences = @client.col(FIRESTORE_PREFERENCES_COLLECTION).doc(user)
    return [] if preferences.get.data.nil?

    preferences.get[:bookmarks]
  end

  def remove_bookmark(user, bookmark)
    preferences = @client.col(FIRESTORE_PREFERENCES_COLLECTION).doc(user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks&.delete(bookmark)
    preferences.set({ bookmarks: bookmarks })
  end
end
