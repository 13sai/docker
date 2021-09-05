
```sh
docker run --detach \
  -p 8443:443 -p 8080:80 -p 8022:22 \
  --name gitlab \
  --restart always \
  --volume `pwd`/config:/etc/gitlab \
  --volume `pwd`/logs:/var/log/gitlab \
  --volume `pwd`/data:/var/opt/gitlab \
  gitlab/gitlab-ce:latest

# 查看密码
docker exec -it gitlab grep 'Password:' /etc/gitlab/initial_root_password
# pSdE9w2kV3x5fKP29/0oo1GcRAD2xQZrR/ZdCR68W2I=
# config/gitlab.rb
external_url 'http://192.168.8.130'
gitlab_rails['gitlab_ssh_host'] = '192.168.8.130'
```