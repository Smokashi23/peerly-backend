
require 'mina/bundler'
require 'mina/default'
require 'mina/deploy'
require 'mina/git'


set :repository, 'https://github.com/joshsoftware/peerly-backend.git'
set :user, 'ubuntu'
set :forward_agent, true

set :shared_files, [ 
  '.env'
]

task :staging do 
  set :deploy_to, '/www/peerly-backend'
  set :domain, ENV['PEERLY_STAGING_IP']
  set :branch, 'Deployment'
end

task :production do 
  set :deploy_to, '/www/peerly'
  set :domain, ENV['PEERLY_STAGING_IP']
  set :branch, 'Deployment'
end

task :setup do
   command %{mkdir -p "#{fetch(:deploy_to)}/releases"}
  #  command %{createdb -U postgres peerly}
end

task :loadData do
  command %{make seed}
  command %{make loadUser}
end

task :deploy do
  deploy do
    invoke :'git:clone'
    command "git checkout #{fetch(:branch)}"
    invoke :'deploy:link_shared_paths'

    command %{export PATH=$PATH:/usr/local/go/bin}
    command %[echo "-----> go get"]
    command %{go mod tidy}
    command %[echo "-----> go mod vendor"]
    command %{go mod vendor}
    command %[echo "-----> Creating build"]
    command %{go build cmd/main.go}
    command %[echo "-----> Build Done"]
    # command %{make migrate}
    
    command %{sudo systemctl restart golang.service}  
    invoke :'deploy:cleanup' 
  end
end
