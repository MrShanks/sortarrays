# Log Collection using fluentd Kibana and elasticsearch

## Installation

create namespace log-monitoring
```
kubectl create namespace log-monitoring
```

add elastic helm repo to your local repos
```
helm repo add elastic https://helm.elastic.co
helm repo update
```

run elasticsearch helm installtion
```
helm install elasticsearch elastic/elasticsearch -n log-monitoring --set replicas=1
```

run kibana helm installation
```
helm install kibana elastic/kibana -n log-monitoring
```

install fluentd as a daemonset
```
kubectl apply -f ./fluentd
```

## Interact with the dashboard

Expand the burger menu and click on **Management -> Stack Management**

On the Stack Management page, select **Data -> Index Management**

Once logstash is indexed  go to **Kibana -> Index Patterns** and then **Create index pattern** button.

Define a new index Pattern called `logstash*` and then click on **Next step** button to continue

Configure the primary time field to use with the new index pattern by selecting the @timestamp option from the **Time field** drop-down. Click the **Create index pattern** button to complete creation of the index pattern.

To explore the indexed data, expand the drop-down menu and click **Analytics -> Discover**

Search for log in the filter column and add it as a column pressing by the plus button.

Everything is set you can now experimenting with queries to retrieve logs you need
