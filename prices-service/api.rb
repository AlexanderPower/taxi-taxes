# frozen_string_literal: true
require 'grape'
require 'http'
require_relative './calculate_price'

class API < Grape::API
  version 'v1', using: :path
  format :json
  prefix :api

  ROUTES_SERVICE_URL = 'http://routes-service:8080'

  helpers do
    def prepare_route_request(points)
      points.map {|point| "#{point[:lat]},#{point[:lon]}"}.join('|')
    end

  end

  desc 'Получить цену поездки'
  params do
    requires :points, type: Array, allow_blank: false, desc: 'Точки пути - массив [lat, lon]' do
      requires :lat, type: Float
      requires :lon, type: Float
    end
  end
  get 'price' do
    request = prepare_route_request declared(params)[:points]
    begin
      response = HTTP.get "#{ROUTES_SERVICE_URL}/info?points=#{request}"
      body = JSON.parse(response.body)
      if response.status == 200
        CalculatePrice.new(body['Distance'], body['Duration']).call
      else
        error! body['error'], 400
      end
    rescue StandardError => ex
      error! ex, 400
    end
  end
end
