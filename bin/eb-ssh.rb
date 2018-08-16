#!/usr/bin/env ruby

require 'optparse'
require 'aws-sdk-elasticbeanstalk'
require 'aws-sdk-ec2'

params = ARGV.getopts('', 'env-name:', 'region:', 'profile:', 'keyfile:', 'proxy-host:')
env_name = params['env-name']
region = params['region'] || ENV['AWS_REGION']
profile = params['profile'] || ENV['AWS_PROFILE']
keyfile = params['keyfile']
proxy_host = params['proxy-host']

# unless EBENVS.keys.include?(rack_env) && eb_env_name
#   puts 'invalid env.'
#   exit 1
# end

puts "env_name: #{env_name}"
puts "region: #{region}"
puts "profile: #{profile}"
puts "keyfile: #{keyfile}"
puts "proxy_host: #{proxy_host}"

elasticbeanstalk = Aws::ElasticBeanstalk::Client.new(
  region: region,
  profile: profile,
)
