## docker/kubernetes


helm install --name-template gitlab -n gitlab stable/gitlab-ce  \
 -f values.yaml   -f charts/postgresql/values.yaml  -f charts/redis/values.yaml 