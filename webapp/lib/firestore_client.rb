# frozen_string_literal: true

require 'ons-firestore'

# Class to manage access to Firestore.
class FirestoreClient
  FIRESTORE_CLUSTERS_COLLECTION       = 'gcp-projects-dashboard'
  FIRESTORE_MASTER_ALERTS_COLLECTION  = 'gcp-projects-dashboard-gke-master-version-alerts'
  FIRESTORE_NODE_ALERTS_COLLECTION    = 'gcp-projects-dashboard-gke-node-version-alerts'
  FIRESTORE_PREFERENCES_COLLECTION    = 'gcp-projects-dashboard-preferences'
  FIRESTORE_SECURITY_RULES_COLLECTION = 'gcp-projects-dashboard-cloud-armour-security-rules'

  def initialize(project)
    @firestore = Firestore.new(project)
  end

  def add_bookmark(user, bookmark)
    preferences = @firestore.document_reference(FIRESTORE_PREFERENCES_COLLECTION, user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks << bookmark unless bookmarks.include?(bookmark)
    preferences.set({ bookmarks: })
  end

  def all_master_version_alerts
    @firestore.all_documents(FIRESTORE_MASTER_ALERTS_COLLECTION)
  end

  def all_node_version_alerts
    @firestore.all_documents(FIRESTORE_NODE_ALERTS_COLLECTION)
  end

  def all_projects
    @firestore.all_documents(FIRESTORE_CLUSTERS_COLLECTION)
  end

  def all_security_rules
    @firestore.all_documents(FIRESTORE_SECURITY_RULES_COLLECTION)
  end

  def bookmarked_projects(user)
    bookmarks = bookmarks(user)
    bookmarked_projects = []
    all_projects.each { |project| bookmarked_projects << project if bookmarks.include?(project.document_id) }
    bookmarked_projects
  end

  def bookmarked_security_rules(user)
    bookmarks = bookmarks(user)
    bookmarked_projects = []
    all_security_rules.each { |security_rule| bookmarked_projects << security_rule if bookmarks.include?(security_rule.document_id) }
    bookmarked_projects
  end

  def bookmarks(user)
    preferences = @firestore.document_reference(FIRESTORE_PREFERENCES_COLLECTION, user)
    return [] if preferences.get.data.nil?

    preferences.get[:bookmarks]
  end

  def remove_bookmark(user, bookmark)
    preferences = @firestore.document_reference(FIRESTORE_PREFERENCES_COLLECTION, user)
    bookmarks = []
    bookmarks = preferences.get[:bookmarks] unless preferences.get.data.nil?
    bookmarks&.delete(bookmark)
    preferences.set({ bookmarks: })
  end
end
