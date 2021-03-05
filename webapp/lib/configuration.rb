# frozen_string_literal: true

# Simple class to centralise access to configuration.
class Configuration
  attr_reader :firestore_project,
              :gcp_console_base_url,
              :gcp_console_cloud_armour_base_url,
              :gcp_organisation

  def initialize(env)
    @firestore_project                 = env['FIRESTORE_PROJECT']
    @gcp_console_base_url              = env['GCP_CONSOLE_BASE_URL']
    @gcp_console_cloud_armour_base_url = env['GCP_CONSOLE_CLOUD_ARMOUR_BASE_URL']
    @gcp_organisation                  = env['GCP_ORGANISATION']

    raise 'Missing FIRESTORE_PROJECT environment variable' unless @firestore_project
    raise 'Missing GCP_CONSOLE_BASE_URL environment variable' unless @gcp_console_base_url
    raise 'Missing GCP_CONSOLE_CLOUD_ARMOUR_BASE_URL environment variable' unless @gcp_console_cloud_armour_base_url
    raise 'Missing GCP_ORGANISATION environment variable' unless @gcp_organisation
  end
end
