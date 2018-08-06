# frozen_string_literal: true

require 'yaml'

class CalculatePrice
  CONFIG = YAML.load_file('config.yml')['price'].symbolize_keys

  def initialize(distance, duration)
    # Берем только целую часть
    @distance = distance / 1000
    @duration = duration / 60
  end

  def call()
    price = CONFIG[:arrive] + CONFIG[:per_km] * @distance + CONFIG[:per_minute] * @duration
    price < CONFIG[:min] ? CONFIG[:min] : price
  end
end