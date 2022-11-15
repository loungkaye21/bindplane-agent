# RabbitMQ Plugin

Log parser for RabbitMQ

## Configuration Parameters

| Name | Description | Type | Default | Required | Values |
|:-- |:-- |:-- |:-- |:-- |:-- |
| daemon_log_paths | The absolute path to the RabbitMQ Daemon logs | []string | `[/var/log/rabbitmq/rabbit@*.log]` | false |  |
| start_at | At startup, where to start reading logs from the file (`beginning` or `end`) | string | `end` | false | `beginning`, `end` |
| offset_storage_dir | The directory that the offset storage file will be created | string | `$OIQ_OTEL_COLLECTOR_HOME/storage` | false |  |

## Example Config:

Below is an example of a basic config

```yaml
receivers:
  plugin:
    path: ./plugins/rabbitmq_logs.yaml
    parameters:
      daemon_log_paths: [/var/log/rabbitmq/rabbit@*.log]
      start_at: end
      offset_storage_dir: $OIQ_OTEL_COLLECTOR_HOME/storage
```