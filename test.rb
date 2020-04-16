#!/usr/bin/env ruby
# frozen_string_literal: true

require 'json'
require 'test/unit/assertions'

include Test::Unit::Assertions

host = 'localhost'
port = '8080'
address = "#{host}:#{port}"
token = 'token'

status = {
  version: 'version',
  commit: 'commit'
}

sensors = {
  timestamp: Time.now.utc.to_i,
  temperature: 20.5,
  humidity: 40.5,
  pressure: 1000.5,
  gas: 10
}

def curl(url, endpoint, headers: [], data: nil, no_json: false)
  cmd = "curl -s '#{url}#{endpoint}'"
  return system(cmd, out: File::NULL, err: File::NULL) if no_json

  cmd += " -d '#{JSON.generate(data)}'" if data
  headers.each { |header| cmd += " -H '#{header}'" }
  JSON.parse(`#{cmd}`, symbolize_names: true)
end

def compose(args)
  system("docker-compose #{args}", exception: true)
end

compose "build --build-arg VERSION=#{status[:version]} --build-arg COMMIT=#{status[:commit]}"
compose 'up -d'
Signal.trap 'EXIT' do
  compose 'stop'
end

print 'Waiting for server to come online '
20.times do
  break if curl(address, '/', no_json: true)

  print '.'
  sleep 2
end
puts

assert_equal status, curl(address, '/')
assert_equal sensors, curl(address, '/sensors', headers: ['Content-Type: application/json', "Authorization: Bearer #{token}"], data: sensors)
assert_equal sensors, curl(address, '/sensors')
